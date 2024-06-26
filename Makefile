gen_auth:
	@protoc -I proto proto/auth/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
gen_manager:
	@protoc -I proto proto/manager/*.proto --go_out=./gen/go/ --go_opt=paths=source_relative --go-grpc_out=./gen/go/ --go-grpc_opt=paths=source_relative
gen_auth_js:
	@protoc -I proto proto/auth/*.proto --js_out=library=commonjs:./gen/js/auth --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./gen/js/


