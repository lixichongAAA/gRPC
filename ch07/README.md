# Chapter 7: Running gRPC in Production

## ``ProductInfo`` Service and Client 

- Online retail scenario has a `` ProductInfo`` microservice which is responsible for managing the products and their information. The consumer of that service can add, retrieve products via that service. 

## Service Definition 

```proto
package ecommerce;

service ProductInfo {
    rpc addProduct(Product) returns (ProductID);
    rpc getProduct(ProductID) returns (Product);
}

message Product {
    string id = 1;
    string name = 2;
    string description = 3;
    float price = 4;
}

message ProductID {
    string value = 1;
}
```

## Implementation
- gRPC Continous Integration
- Deploy in Docker
- Deploy in Kubernetes: 待实现，Kubernetes待学习
- OpenCensus Metrics
- OpenCensus Tracing
- OpenTracing
- Prometheus