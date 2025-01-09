# pip install protobuf==3.20.3 grpcio==1.43.0 grpcio-tools==1.43.0 --only-binary=:all
# ？？？https://blog.csdn.net/yueguangMaNong/article/details/127502700
grpc_generate_python:
	python3 -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. --pyi_out=. ccanalyzer.proto
	#protoc --proto_path=. -I . --python_out=. --pyi_out=. --grpc_python_out=. ccanalyzer_client.proto