# ch06 gRPC安全相关

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

- secure-channel: 单向安全连接
- mutual-tls-channel: 双向安全连接
- basic-authentication: Basic 认证
- token-based-authentication: OAuth2