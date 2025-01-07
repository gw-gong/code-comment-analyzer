import grpc
from concurrent import futures
from grpc_server import ccanalyzer_pb2
from grpc_server import ccanalyzer_pb2_grpc

class CcAnalyzer(ccanalyzer_pb2_grpc.CcAnalyzerServicer):

    # 实现proto文件中rpc的调用
    def AddUser(self, request: ccanalyzer_pb2.UserRequest, context):
        print(request.name, request.age)
        return ccanalyzer_pb2.UserResponse(msg='add user(name={},age={}) success'.format(request.name, request.age), code=0)

    def GetUser(self, request: ccanalyzer_pb2.GetUserRequest, context):
        print(request.name)
        return ccanalyzer_pb2.GetUserResponse(name=request.name, age="1888")


def serve():
    # 使用线程池来完成grpc的请求
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=5))
    ccanalyzer_pb2_grpc.add_CcAnalyzerServicer_to_server(CcAnalyzer(), server)
    print('listening ' + '127.0.0.1:8888')
    server.add_insecure_port('127.0.0.1:8888')  # 绑定端口
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    print('start')
    serve()