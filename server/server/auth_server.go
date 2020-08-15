package server

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/storage"
	"github.com/ory/fosite/token/jwt"
	goauth "golang.org/x/oauth2"
)

// OAuth2Server serve
type OAuth2Server struct {
	Configuration *compose.Config
	Storage       *storage.MemoryStore
	Secret        []byte
	PrivateKey    *rsa.PrivateKey
	oauth         fosite.OAuth2Provider
}

func (oauthServe *OAuth2Server) buildServer() {
	if oauthServe.Configuration == nil {
		oauthServe.Configuration = &compose.Config{
			AccessTokenLifespan: time.Minute * 30,
		}
	}
	if oauthServe.Storage == nil {
		oauthServe.Storage = storage.NewExampleStore()
	}
	if oauthServe.Secret == nil {
		oauthServe.Secret = []byte("some-cool-secret-that-is-32bytes")
	}
	if oauthServe.PrivateKey == nil {
		oauthServe.PrivateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
	}

	oauthServe.oauth = compose.ComposeAllEnabled(
		oauthServe.Configuration,
		oauthServe.Storage,
		oauthServe.Secret,
		oauthServe.PrivateKey)
}

// Serve ff
func (oauthServe *OAuth2Server) Serve() error {
	oauthServe.buildServer()

	http.HandleFunc("/oauth2/auth", oauthServe.authEndpoint)
	http.HandleFunc("/oauth2/token", oauthServe.tokenEndpoint)
	http.HandleFunc("/oauth2/revoke", oauthServe.revokeEndpoint)
	http.HandleFunc("/oauth2/introspect", oauthServe.introspectionEndpoint)
	http.HandleFunc("/login", oauthServe.ownerHandler(clientConf))
	log.Print("Starting OAuth2 server")
	return http.ListenAndServe(":3846", nil)
}

func (oauthServe *OAuth2Server) authEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ar, err := oauthServe.oauth.NewAuthorizeRequest(ctx, req)
	if err != nil {
		log.Printf("Error occurred in NewAuthorizeRequest: %+v", err)
		oauthServe.oauth.WriteAuthorizeError(rw, ar, err)
		return
	}
	var requestedScopes string
	for _, this := range ar.GetRequestedScopes() {
		requestedScopes += fmt.Sprintf(`<li><input type="checkbox" name="scopes" value="%s">%s</li>`, this, this)
	}
	req.ParseForm()
	if req.PostForm.Get("username") != "peter" {
		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
		rw.Write([]byte(`<h1>Login page</h1>`))
		rw.Write([]byte(fmt.Sprintf(`
			<p>Howdy! This is the log in page. For this example, it is enough to supply the username.</p>
			<form method="post">
				<p>
					By logging in, you consent to grant these scopes:
					<ul>%s</ul>
				</p>
				<input type="text" name="username" /> <small>try peter</small><br>
				<input type="submit">
			</form>
		`, requestedScopes)))
		return
	}
	for _, scope := range req.PostForm["scopes"] {
		ar.GrantScope(scope)
	}
	mySessionData := newSession("peter")
	response, err := oauthServe.oauth.NewAuthorizeResponse(ctx, ar, mySessionData)
	if err != nil {
		log.Printf("Error occurred in NewAuthorizeResponse: %+v", err)
		oauthServe.oauth.WriteAuthorizeError(rw, ar, err)
		return
	}
	oauthServe.oauth.WriteAuthorizeResponse(rw, ar, response)
}

func (oauthServe *OAuth2Server) tokenEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	mySessionData := newSession("")
	accessRequest, err := oauthServe.oauth.NewAccessRequest(ctx, req, mySessionData)
	if err != nil {
		log.Printf("Error occurred in NewAccessRequest: %+v", err)
		oauthServe.oauth.WriteAccessError(rw, accessRequest, err)
		return
	}
	if accessRequest.GetGrantTypes().Exact("client_credentials") {
		for _, scope := range accessRequest.GetRequestedScopes() {
			if fosite.HierarchicScopeStrategy(accessRequest.GetClient().GetScopes(), scope) {
				accessRequest.GrantScope(scope)
			}
		}
	}
	response, err := oauthServe.oauth.NewAccessResponse(ctx, accessRequest)
	if err != nil {
		log.Printf("Error occurred in NewAccessResponse: %+v", err)
		oauthServe.oauth.WriteAccessError(rw, accessRequest, err)
		return
	}
	oauthServe.oauth.WriteAccessResponse(rw, accessRequest, response)
}

func (oauthServe *OAuth2Server) introspectionEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	mySessionData := newSession("")
	ir, err := oauthServe.oauth.NewIntrospectionRequest(ctx, req, mySessionData)
	if err != nil {
		log.Printf("Error occurred in NewIntrospectionRequest: %+v", err)
		oauthServe.oauth.WriteIntrospectionError(rw, err)
		return
	}
	oauthServe.oauth.WriteIntrospectionResponse(rw, ir)
}

func (oauthServe *OAuth2Server) revokeEndpoint(rw http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	err := oauthServe.oauth.NewRevocationRequest(ctx, req)
	oauthServe.oauth.WriteRevocationResponse(rw, err)
}

//OwnerHandler hh
func (oauthServe *OAuth2Server) ownerHandler(c goauth.Config) func(rw http.ResponseWriter, req *http.Request) {
	return func(rw http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.Form.Get("username") == "" || req.Form.Get("password") == "" {
			loginPage := template.Must(template.ParseFiles("./server/static/login.html"))
			loginPage.Execute(rw, nil)
			return
		}
		token, err := c.PasswordCredentialsToken(context.Background(), req.Form.Get("username"), req.Form.Get("password"))
		if err != nil {
			rw.Write([]byte(fmt.Sprintf(`<p>I tried to get a token but received an error: %s</p>`, err.Error())))
			rw.Write([]byte(`<p><a href="/">Go back</a></p>`))
			return
		}
		rw.Write([]byte(fmt.Sprintf(`<p>Awesome, you just received an access token!<br><br>%s<br><br><strong>more info:</strong><br><br>%s</p>`, token.AccessToken, token)))
		rw.Write([]byte(`<p><a href="/">Go back</a></p>`))
	}
}

// A valid oauth2 client (check the store) that additionally requests an OpenID Connect id token
func newSession(user string) *openid.DefaultSession {
	return &openid.DefaultSession{
		Claims: &jwt.IDTokenClaims{
			Issuer:      "https://fosite.my-application.com",
			Subject:     user,
			Audience:    []string{"https://my-client.my-application.com"},
			ExpiresAt:   time.Now().Add(time.Hour * 6),
			IssuedAt:    time.Now(),
			RequestedAt: time.Now(),
			AuthTime:    time.Now(),
		},
		Headers: &jwt.Headers{
			Extra: make(map[string]interface{}),
		},
	}
}

var clientConf = goauth.Config{
	ClientID:     "my-client",
	ClientSecret: "foobar",
	RedirectURL:  "http://localhost:3846/callback",
	Scopes:       []string{"photos", "openid", "offline"},
	Endpoint: goauth.Endpoint{
		TokenURL: "http://localhost:3846/oauth2/token",
		AuthURL:  "http://localhost:3846/oauth2/auth",
	},
}
