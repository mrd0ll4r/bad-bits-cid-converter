#!/bin/bash -e

# Build image
docker build -t mrd0ll4r/bad-bits-converter .

# Extract binary
docker create --name extract mrd0ll4r/bad-bits-converter
mkdir -p out
docker cp extract:/bad-bits-cid-converter/bad-bits-cid-converter ./out/bad-bits-cid-converter
docker rm extract

