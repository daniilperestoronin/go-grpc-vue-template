package server

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/daniilperestoronin/go-grpc-vue-template/api"
	"github.com/daniilperestoronin/go-grpc-vue-template/services"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
)

// GrpcServer is Simple grps server structure
type GrpcServer struct {
	Address    string
	WebAddress string
}

// Serve start GrpcServer with defined options
func (server *GrpcServer) Serve() error {
	log.Printf("Starting grpc server...")
	// create a listener on TCP port
	wlis, err := net.Listen("tcp", server.WebAddress)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	s := services.PingService{}
	grpcServer := grpc.NewServer()

	api.RegisterPingServer(grpcServer, &s)

	grpcWebServer := grpcweb.WrapServer(grpcServer)

	httpServer := &http.Server{
		Handler: h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.ProtoMajor == 2 {
				grpcWebServer.ServeHTTP(w, r)
			} else {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-User-Agent, X-Grpc-Web")
				w.Header().Set("grpc-status", "")
				w.Header().Set("grpc-message", "")
				if grpcWebServer.IsGrpcWebRequest(r) {
					grpcWebServer.ServeHTTP(w, r)
				}
			}
		}), &http2.Server{}),
	}
	httpServer.Serve(wlis)

	return nil
}
