[TOC]


## 命令参考

### 环境搭建
- 安装go环境
    - 安装go管理插件g：参考[github](https://github.com/voidint/g)
    - 设置代理：`go env -w GOPROXY=https://goproxy.cn,direct`
- 安装protoc
    - 进入[protobuf release](https://github.com/protocolbuffers/protobuf/releases)
    - 下载对应压缩包并解压进入目录
    - 将启动的protoc二进制文件移动到环境变量path下，如：
      > mv protoc /root/bin/
    - 验证安装结果：
      > protoc --version
    - 安装goctl：
      > go get -u github.com/zeromicro/go-zero/tools/goctl
      > cp /root/go/bin/goctl /root/bin/

### 构建工程
```
cd path/proto
goctl rpc proto -src x.proto -dir ../ --style go_zero
```

### 构建model
```
cd path/model
goctl model mysql ddl -src configs.sql -dir . -style go_zero -c
```

### 生成base.proto
```
protoc --go_out=../pb/ --go_opt=paths=source_relative commands.proto message.proto

```

------------


### docker相关

#### centos7安装docker

###### 卸载已有docker
- `service stop docker`
- `yum remove docker`
- `yum remove docker-common`
- `yum remove docker-client`

######安装docker
```
curl -fsSL https://get.docker.com/ | sh
service docker start
```
###### 修改docker存储路径
- 新建存储路径
> mkdir /opt/data/docker
- 修改docker.service配置
```
# vim /usr/lib/systemd/system/docker.service
#ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock (注释原先的)
ExecStart=/usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock --graph=/data/docker（新增的）
```
- 重启服务
-  `systemctl daemon-reload`
- `systemctl restart docker`
------------


#### 生成Dockerfile并运行

###### Dockerfile
- `goctl docker -go hello.go`

###### build
- `docker build -t hello:v1 -f Dockerfile .`
- `docker build -t call:v1 -f call/rpc/Dockerfile .`

###### run
- ` docker run --rm -it -p 8888:8888 --name hello hello:v1`
- `docker run  --network host --name call -d call:v2`

------------


#### docker常用命令
- 进入bash：`docker exec -it {container_id} sh `
- 查看日志：`docker logs -f {container name}`
- 移除none镜像：`docker image prune`
- 批量移除镜像：`docker images | grep trunking-debug | awk '{print $3}' | xargs docker rmi -f`

#### 华为云构建容器
- `docker tag greet:v2 swr.cn-south-1.myhuaweicloud.com/yyxxgame_houtai/greet:v2`
- `docker push swr.cn-south-1.myhuaweicloud.com/yyxxgame_houtai/greet:v2`

### 参考文档
[go-zero官方文档](https://go-zero.dev/cn/)
[华为云容器镜像上传服务](https://support.huaweicloud.com/usermanual-swr/swr_01_0011.html)