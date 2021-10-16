#!/bin/bash
protoc src/maximum_proto/maximum.proto --go_out=. --go-grpc_out=.