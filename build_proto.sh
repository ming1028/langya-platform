#!/bin/bash

protoc -I ./protobuf \
       --go_out ../ \
       --go-grpc_out ../ \
       ./protobuf/langya_platform_app.proto

protoc -I ./protobuf \
       --go_out ../ \
       --go-grpc_out ../ \
       ./protobuf/langya_platform_common.proto