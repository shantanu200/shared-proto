.PHONY: generate clean

generate:
	mkdir -p gen/go/product
	protoc --proto_path=proto \
		--proto_path=./third_party \
		--go_out=gen/go --go_opt=paths=source_relative \
		--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		product/product.proto

clean:
	rm -rf gen/go/*
