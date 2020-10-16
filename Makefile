.PHONY: all
	all: build-protobuf build-html;

.PHONY: build-protobuf
build-protobuf:
	@protoc -I ./grpc \
		--go_out=plugins=grpc:./grpc\
		--go_opt=module=github.com/takasho-tutrial/dummy-api\
		dummy-api.proto

.PHONY: build-html
build-html:
	@protoc --doc_out=html,index.html:./grpc\
		./grpc/dummy-api.proto

.PHONY: format-protobuf
format-protobuf:
	@clang-format -i ./grpc/dummy-api.proto

.PHONY: clean
clean:
	@rm -f ./grpc/dummy-api.pd.go ./grpc/index.html