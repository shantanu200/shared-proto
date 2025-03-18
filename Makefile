.PHONY: generate clean

generate_product:
	mkdir -p gen/go/product
	protoc --proto_path=proto \
		--proto_path=./third_party \
		--go_out=gen/go --go_opt=paths=source_relative \
		--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		product/product.proto

generate_user:
	mkdir -p gen/go/user
	protoc --proto_path=proto \
		--proto_path=./third_party \
		--go_out=gen/go --go_opt=paths=source_relative \
		--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		user/user.proto

clean:
	rm -rf gen/go/*
