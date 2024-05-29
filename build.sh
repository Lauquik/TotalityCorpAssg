#!/bin/bash

PROTO_DIR="./api/proto"
OUT_DIR="./api/pb"

rm -rf $OUT_DIR
mkdir -p ${OUT_DIR}

# Generate Go files from .proto files
for proto_file in ${PROTO_DIR}/*.proto; do
    protoc \
        --proto_path=${PROTO_DIR} \
        --go_out=${OUT_DIR} \
        --go_opt=paths=source_relative \
        --go-grpc_out=${OUT_DIR} \
        --go-grpc_opt=paths=source_relative \
        ${proto_file}
done

echo "Proto files generated successfully"
