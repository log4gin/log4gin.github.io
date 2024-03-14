from concurrent import futures
import time
import grpc

import hello_pb2
import hello_pb2_grpc


class RequestRpc(hello_pb2_grpc.Greeter):
    # 实现 proto 文件中定义的 rpc 调用
    def SayHello(self, request, context):
        return hello_pb2.HelloResp(Resp='hello {msg}'.format(msg=request.Name))


def serve():
    # 启动 rpc 服务，这里可定义最大接收和发送大小(单位M)，默认只有4M
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10), options=[
        ('grpc.max_send_message_length', 100 * 1024 * 1024),
        ('grpc.max_receive_message_length', 100 * 1024 * 1024)])

    hello_pb2_grpc.add_GreeterServicer_to_server(RequestRpc(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    try:
        while True:
            time.sleep(60*60*24)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
