
.PHONY: all
all: generate

.PHONY: build
build:
	go build -o bin/cli generated/cli/*.go

.PHONY: generate
generate: grpcgen gapicgen gencli genrest genopenapiyaml

.PHONY: grpcgen
grpcgen:
	mkdir -p generated/grpcgen
	protoc -I /home/wen/soft/googleapis \
	--go_out=generated/grpcgen \
	--go_opt=paths=source_relative \
    --go-grpc_out=generated/grpcgen \
	--go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=generated/grpcgen \
	--grpc-gateway_opt=paths=source_relative \
	--grpc-gateway_opt=generate_unbound_methods=true \
    --proto_path=protos \
	protos/server.proto

.PHONY: gapicgen
gapicgen:
	rm -rf generated/gapicgen
	mkdir -p generated/gapicgen
	protoc -I /home/wen/soft/googleapis \
		--go_gapic_out=generated/gapicgen \
		--go_gapic_opt='go-gapic-package=core.wcloud.io/generated/gapicgen/server;server' \
		--go_gapic_opt=transport=grpc+rest \
		--proto_path=protos protos/server.proto
	mv generated/gapicgen/core.wcloud.io/generated/gapicgen/server/ generated/gapicgen
	rm -r generated/gapicgen/core.wcloud.io
	#hack/doreplace.sh google.golang.org/genproto/googleapis/example/server/v1 core.wcloud.io/generated/grpcgen
	#hack/doreplace.sh core.wcloud.io/generated/grpcgen core.wcloud.io/generated/grpcgen
	hack/doreplace.sh core.wcloud.io/generated/grpcgen core.wcloud.io/generated/grpcgen
	#sed -i 's,core.wcloud.io/generated/grpcgen,core.wcloud.io/generated/grpcgen,g' generated/gapicgen/server/server_client.go
	#sed -i 's,core.wcloud.io/generated/grpcgen,core.wcloud.io/generated/grpcgen,g' generated/gapicgen/server/auxiliary.go

.PHONY: gencli
gencli:
	mkdir -p generated/cli
	protoc -I /home/wen/soft/googleapis \
		--experimental_allow_proto3_optional \
		--proto_path=protos \
		--go_cli_out=generated/cli \
		--go_cli_opt=gapic=core.wcloud.io/generated/gapicgen/server \
		--go_cli_opt=fmt=false \
		--go_cli_opt=root=wcloud \
		protos/server.proto
	sed -i 's,gapic.ServerClient,gapic.Client,g'  generated/cli/server_service.go
	sed -i 's,gapic.NewServerClient,gapic.NewClient,g'  generated/cli/server_service.go

# ln -sf $HOME/github/googleapis/gapic-showcase/schema .
.PHONY: genrest
genrest:
	# cp -r ~/github/googleapis/gapic-showcase/schema .
	mkdir -p generated/genrest
	protoc -I /home/wen/soft/googleapis \
		--experimental_allow_proto3_optional \
		--proto_path=protos \
		--go_rest_server_out=generated/genrest \
		protos/server.proto
	sed -i 's,github.com/googleapis/gapic-showcase/server/services,core.wcloud.io/server/services,g' generated/genrest/genrest.go
	sed -i '/Location/d' generated/genrest/genrest.go
	sed -i '/Permission/d' generated/genrest/genrest.go
	sed -i '/Policy/d' generated/genrest/genrest.go
	sed -i 's,core.wcloud.io/generated/grpcgen,core.wcloud.io/generated/grpcgen,g' generated/genrest/serverservice.go
	mv -f generated/genrest/iampolicy.go generated/genrest/iampolicy.go.bak || :
	mv -f generated/genrest/locations.go generated/genrest/locations.go.bak || :

.PHONY: genopenapiyaml
genopenapiyaml:
	protoc -I ~/github/googleapis/googleapis -I=protos --openapi_out=protos server.proto 
