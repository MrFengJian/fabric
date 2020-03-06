# 说明
- 基于fabric 1.4.6源码定制，增加RSA证书的支持，并允许通过配置切换

# 测试

1.创建一个干净的目录作为`GOPATH`，防止其他代码污染。
```Shell
mkdir go_path
export GOPATH=$PWD/go_path
```

2.下载代码，并切换到`fjj_rsa`分支
```Shell
mkdir -p $GOPATH/src/github.com/hyperledger
git clone -b fjj_rsa https://github.com/MrFengjian/fabric.git $GOPATH/src/github.com/hyperledger
```

3.编译

编译二进制文件需要依赖`gcc`、`make`，在源码根目录即可编译：
```shell script
make release
```

编译docker镜像需要依赖docker，并且必须本地运行，执行如下命令即可：
`make docker&&make docker-tag-rsa`
最终生成如下四个镜像:
```shell
hyperledger/fabric-tools:1.4.6-rsa
hyperledger/fabric-buildenv:1.4.6-rsa
hyperledger/fabric-ccenv:1.4.6-rsa
hyperledger/fabric-peer:1.4.6-rsa
hyperledger/fabric-orderer:1.4.6-rsa
```

4.运行测试
使用[rsa_usage](https://github.com/MrFengJian/fabric_examples/tree/master/rsa_usage)，启动测试
```shell script
./start.sh
```
使用脚本`stop.sh`，清理测试资源。