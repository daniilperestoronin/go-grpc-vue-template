install-protoc:
	brew install protobuf
	protoc --version
	brew install protoc-gen-go
	curl -O https://github.com/grpc/grpc-web/releases/download/1.2.0/protoc-gen-grpc-web-1.2.0-darwin-x86_64
	sudo mv ./protoc-gen-grpc-web-1.2.0-darwin-x86_64 /usr/local/bin/protoc-gen-grpc-web
	npm i grpc-web

generate-front-client:
	protoc -I=. ./api/api.proto --js_out=import_style=commonjs:front/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:front/src/

generate-server-interface:
	protoc -I . ./api/api.proto --go_out=plugins=:./server

front-lint:
	npm run lint --prefix front

front-test: front-lint
	npm run test --prefix front

front-build: front-test
	npm run build --prefix front

server-test:
	(cd server/ && go test)

server-build: server-test
	(cd server/ && go build)

app-test: front-test server-test

app-build: generate-front-client front-build generate-server-interface server-build
	mkdir build
	mv ./server/go ./build
	mv ./front/dist ./build
