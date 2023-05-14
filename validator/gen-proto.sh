#!/bin/bash

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false,paths=source_relative \
    --validate_out=paths=source_relative,lang=go:../generated   \
    proto/validator.proto
