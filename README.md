## grpctest
- 此项目用于学习grpc以及grpc-gateway相关知识

### 注意事项一
- 在项目里面增加google/api目录
- 去官网copy两个文件到api目录下
  - [annotations.proto](https://github.com/googleapis/googleapis/blob/master/google/api/annotations.proto)
  - [http.proto](https://github.com/googleapis/googleapis/blob/master/google/api/http.proto)

### 注意事项二
- 使用如下命令生成.gw.proto
```bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative --grpc-gateway_out . --grpc-gateway_opt paths=source_relative ./helloworld.proto
```

### 注意事项三
- 可能报错 descriptor.proto not found
- 解决办法
- 新建google/protobuf
  - 新建文件[descriptor.proto](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/descriptor.proto)

### 在重新生成.go文件，基本上就解决了此问题
