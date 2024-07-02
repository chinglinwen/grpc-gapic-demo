
.PHONY: all
all: generate

.PHONY: build
build:
	go build -o bin/cli generated/cli/*.go

.PHONY: generate
generate: grpcgen gapicgen gencli genrest

.PHONY: grpcgen
grpcgen:
	mkdir -p generated/grpcgen
	protoc -I /home/wen/soft/googleapis \
	--go_out=generated/grpcgen \
	--go_opt=paths=source_relative \
    --go-grpc_out=generated/grpcgen --go-grpc_opt=paths=source_relative \
    --proto_path=protos \
	protos/server.proto

.PHONY: gapicgen
gapicgen:
	rm -rf generated/gapicgen
	mkdir -p generated/gapicgen
	protoc -I /home/wen/soft/googleapis \
		--go_gapic_out=generated/gapicgen \
		--go_gapic_opt='go-gapic-package=apidemo.io/generated/gapicgen/library;library' \
		--go_gapic_opt=transport=grpc+rest \
		--proto_path=protos protos/server.proto
	mv generated/gapicgen/apidemo.io/generated/gapicgen/library/ generated/gapicgen
	rm -r generated/gapicgen/apidemo.io
	#hack/doreplace.sh google.golang.org/genproto/googleapis/example/library/v1 apidemo.io/generated/grpcgen
	#hack/doreplace.sh apidemo.io/generated/grpcgen apidemo.io/generated/grpcgen
	hack/doreplace.sh apidemo.io/generated/grpcgen apidemo.io/generated/grpcgen
	#sed -i 's,apidemo.io/generated/grpcgen,apidemo.io/generated/grpcgen,g' generated/gapicgen/library/library_client.go
	#sed -i 's,apidemo.io/generated/grpcgen,apidemo.io/generated/grpcgen,g' generated/gapicgen/library/auxiliary.go

.PHONY: gencli
gencli:
	mkdir -p generated/cli
	protoc -I /home/wen/soft/googleapis \
		--experimental_allow_proto3_optional \
		--proto_path=protos \
		--go_cli_out=generated/cli \
		--go_cli_opt=gapic=apidemo.io/generated/gapicgen/library \
		--go_cli_opt=fmt=false \
		--go_cli_opt=root=apidemo \
		protos/server.proto
	sed -i 's,gapic.LibraryClient,gapic.Client,g'  generated/cli/library_service.go
	sed -i 's,gapic.NewLibraryClient,gapic.NewClient,g'  generated/cli/library_service.go

.PHONY: genrest
genrest:
	# cp -r ~/github/googleapis/gapic-showcase/schema .
	mkdir -p generated/genrest
	protoc -I /home/wen/soft/googleapis \
		--experimental_allow_proto3_optional \
		--proto_path=protos \
		--go_rest_server_out=generated/genrest \
		protos/server.proto
	sed -i 's,github.com/googleapis/gapic-showcase/server/services,apidemo.io/server/services,g' generated/genrest/genrest.go
	sed -i '/Location/d' generated/genrest/genrest.go
	sed -i '/Permission/d' generated/genrest/genrest.go
	sed -i '/Policy/d' generated/genrest/genrest.go
	sed -i 's,apidemo.io/generated/grpcgen,apidemo.io/generated/grpcgen,g' generated/genrest/libraryservice.go
	mv -f generated/genrest/iampolicy.go generated/genrest/iampolicy.go.bak || :
	mv -f generated/genrest/locations.go generated/genrest/locations.go.bak || :
	
#		--go_gapic_opt=grpc-service-config=schema/showcase_grpc_service_config.json" \
#		--go_gapic_opt=api-service-config=schema/showcase_v1beta1.yaml" \
#		--plugin=protoc-gen-go_gapic \

#		"--go_gapic_opt=grpc-service-config=schema/google/showcase/v1beta1/showcase_grpc_service_config.json",
#		"--go_gapic_opt=api-service-config=schema/google/showcase/v1beta1/showcase_v1beta1.yaml",
#		"--go_out=plugins=grpc:" + outDir,
#		--go_cli_opt=root=apidemo \

#		--go_gapic_out=generated/gapicgen \
#		--go_gapic_opt=go-gapic-package=apidemo.io/generated/gapicgen/library;library \
#		--go_gapic_opt=metadata \
#		--go_gapic_opt=transport=grpc+rest \

.PHONY: genrest2
genrest2:
	mkdir -p generated/gapicgen
	protoc -I /home/wen/soft/googleapis \
		--experimental_allow_proto3_optional \
		--proto_path=protos \
		--go_cli_out=generated/cli \
		--go_cli_opt=gapic=apidemo.io/generated/gapicgen/library \
		--go_cli_opt=fmt=false \
		--go_gapic_out=generated/gapicgen \
		--go_gapic_opt='go-gapic-package=apidemo.io/generated/gapicgen/library;library' \
		--go_gapic_opt=metadata \
		--go_gapic_opt=transport=grpc+rest \
		--go_rest_server_out=generated/genrest \
		--go_gapic_opt=grpc-service-config=schema/showcase_grpc_service_config.json \
		--go_gapic_opt=api-service-config=schema/showcase_v1beta1.yaml \
		--go_cli_opt=root=apidemo \
		protos/server.proto


# 		--go_out=plugins=grpc: \

#		--go_gapic_opt=grpc-service-config=schema/google/showcase/v1beta1/showcase_grpc_service_config.json" \
#		--go_gapic_opt=api-service-config=schema/google/showcase/v1beta1/showcase_v1beta1.yaml" \

.PHONY: genopenapiyaml
genopenapiyaml:
	protoc -I ~/github/googleapis/googleapis -I=protos --openapi_out=protos server.proto 
