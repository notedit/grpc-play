#!/bin/bash




 protoc -I  ../protos signaling.proto \
    --js_out=import_style=typescript:./protos \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:./protos


# protoc \
#   -I ../protos \
#   --ts_out=service=grpc-web:./protos \
#    signaling.proto