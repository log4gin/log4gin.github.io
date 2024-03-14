import grpc
import hello_pb2
import hello_pb2_grpc


def run():
    # 连接 rpc 服务器
    channel = grpc.insecure_channel('localhost:2333')
    # 调用 rpc 服务
    stub = hello_pb2_grpc.GreeterStub(channel)
    resp = stub.SayHello(hello_pb2.HelloReq(Name='python client '))
    print(resp.Resp)


if __name__ == '__main__':
    run()
