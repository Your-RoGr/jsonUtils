#!/bin/bash

cd "$(dirname $0)"
cd ..

go test -cover -race ./jsonUtils/...
