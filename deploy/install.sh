#!/bin/bash

VERSION=1.0.0

echo VERSION
docker load < dlv_${VERSION}.tar
echo "加载完毕。。。"

docker run -d -p 8080:8080 --security-opt=seccomp:unconfined dlv:${VERSION}
echo "成功跑起。。。"