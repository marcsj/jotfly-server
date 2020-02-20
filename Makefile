build-proto:
	mkdir -p notes
	mkdir -p users
	protoc -I/usr/local/include \
		-I./proto \
		-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:notes \
		--grpc-gateway_out=logtostderr=true:notes \
		--go_out=plugins=grpc:./notes \
		proto/notes.proto
	protoc -I/usr/local/include \
    	-I./proto \
    	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    	--swagger_out=logtostderr=true:users \
    	--grpc-gateway_out=logtostderr=true:users \
    	--go_out=plugins=grpc:./users \
    	proto/users.proto