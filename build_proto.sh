#!/bin/bash
# -I 指定protobuf文件路径 --go_out 生成文件路径位置
protoc -I ./protobuf \
       --go_out ../../ \
       --go-grpc_out ../../ \
       --grpc-gateway_out ../../ \
       --openapiv2_out ./views/swagger/ \
       ./protobuf/langya_platform_app.proto #proto文件

protoc -I ./protobuf \
       --go_out ../../  \
       --go-grpc_out ../../ \
       ./protobuf/langya_platform_common.proto