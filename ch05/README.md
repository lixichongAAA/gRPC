# ch05 对gRPC的一些特性的实现(《gRPC 与云原生应用开发》)

## ``OrderManagement`` Service and Client
- Online retail scenario has a `` OrderManagement`` microservice which is responsible for managing the orders and their information. The consumer of that service can add, retrieve, search and update order via that service. 

## Service Definition

```proto
package ecommerce;

service OrderManagement {
    rpc addOrder(Order) returns (google.protobuf.StringValue);
    rpc getOrder(google.protobuf.StringValue) returns (Order);
    rpc searchOrders(google.protobuf.StringValue) returns (stream Order);
    rpc updateOrders(stream Order) returns (google.protobuf.StringValue);
    rpc processOrders(stream google.protobuf.StringValue) returns (stream CombinedShipment);
}

message Order {
    string id = 1;
    repeated string items = 2;
    string description = 3;
    float price = 4;
    string destination = 5;
}

message CombinedShipment {
    string id = 1;
    string status = 2;
    repeated Order ordersList = 3;
}
```
## Implementation

- Interceptors: 拦截器，实现了客户端和服务器端的一元RPC拦截器和流拦截器
- Deadline  截止时间，使用context.WithDeadline() AddOrder方法
- Cancellation 取消，使用context.WithCancel() ProcessOrders方法
- Compress 压缩，在发送RPC时设置一个压缩器，服务器端导入包即可
- error-handing 错误处理，使用gRPC的codes包和status包等，错误码和错误状态
- multiplexing 多路复用，同一个gRPC服务器端运行对各gRPC服务(主要用于同一服务器端一个服务的多个版本)
- metadata 元数据，在一元RPC和流RPC场景下，实现了客户端和服务器端发送和接收元数据
- loadbalance 负载均衡，实现了厚客户端