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

generate_cart:
	mkdir -p gen/go/cart
	protoc --proto_path=proto \
		--proto_path=./third_party \
		--go_out=gen/go --go_opt=paths=source_relative \
		--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		cart/cart.proto

generate_shop:
	mkdir -p gen/go/shop
	protoc --proto_path=proto \
		--proto_path=./third_party \
		--go_out=gen/go --go_opt=paths=source_relative \
		--go-grpc_out=gen/go --go-grpc_opt=paths=source_relative \
		shop/shop.proto

clean:
	rm -rf gen/go/*
