prepare_dev:
	npm i grpc-web
	curl -O https://github.com/grpc/grpc-web/releases/download/1.2.0/protoc-gen-grpc-web-1.2.0-darwin-x86_64
	sudo mv ./protoc-gen-grpc-web-1.2.0-darwin-x86_64 /usr/local/bin/protoc-gen-grpc-web

generate_fe_client:
	protoc -I=. ./api/api.proto --js_out=import_style=commonjs:front/src/ --grpc-web_out=import_style=commonjs,mode=grpcwebtext:front/src/

generate_be_interface:
	protoc -I . ./api/api.proto --go_out=plugins=:./server