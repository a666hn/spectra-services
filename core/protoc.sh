#!/bin/bash
protoc --proto_path=${pwd}:. --micro_out=. --go_out=. proto/*.proto