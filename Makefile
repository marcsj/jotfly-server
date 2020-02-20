build-proto:
	mkdir -p notes
	protoc -I/usr/local/include \
		-I./proto \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true,allow_merge=true:. \
		--grpc-gateway_out=logtostderr=true:notes \
		--go_out=plugins=grpc:./notes \
		proto/notes.proto