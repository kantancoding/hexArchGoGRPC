gen:
	#	--go_out=plugins=grpc:path/of/pb/package
	# --proto_path=path/of/imports.proto
	protoc \
		--go_out=plugins=grpc:internal/adapters/framework/in/grpc \
		--proto_path=internal/adapters/framework/in/grpc/proto \
		internal/adapters/framework/in/grpc/proto/*.proto

clean:
	rm -rf internal/adapters/framework/in/grpc/pb
	
run:
	echo "run"
