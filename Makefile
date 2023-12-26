generate_grpc_code:
	protoc --go_out=cinnamon --go_opt=paths=source_relative --go-grpc_out=cinnamon --go-grpc_opt=paths=source_relative cinnamon.proto
