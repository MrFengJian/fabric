# 说明
- 基于fabric 1.4.6源码定制，增加国密算法SM2、SM3、SM4的支持。使用SM2WithSM3作为证书签名验签算法
- 由于SM2算法硬编码替代了ECDSA算法的部分实现，**不再支持ECDSA算法**

# 测试

1.创建一个干净的目录作为`GOPATH`，防止其他代码污染。
```Shell
mkdir go_path
export GOPATH=$PWD/go_path
```

2.下载代码，并切换到`fjj_sm`分支
```Shell
mkdir -p $GOPATH/src/github.com/hyperledger
git clone -b fjj_sm https://github.com/MrFengjian/fabric.git $GOPATH/src/github.com/hyperledger
```

3.编译

编译二进制文件需要依赖`gcc`、`make`，在源码根目录即可编译：
`make release`

编译docker镜像需要依赖docker，并且必须本地运行，执行如下命令即可：
`make docker&&make docker-tag-sm`
最终生成如下四个镜像:
```shell
hyperledger/fabric-tools:1.4.6-sm
hyperledger/fabric-buildenv:1.4.6-sm
hyperledger/fabric-ccenv:1.4.6-sm
hyperledger/fabric-peer:1.4.6-sm
hyperledger/fabric-orderer:1.4.6-sm
```

4.运行测试
使用[sm_usage](https://github.com/MrFengJian/fabric_examples/tree/master/sm_usage)，启动测试
```shell script
./start.sh
```
使用脚本`stop.sh`，清理测试资源。
# 参考资料

- [Hyperledger Fabric国密改造](https://www.cnblogs.com/laolieren/p/hyperledger_fabric_gm_summary.html)
- [Hyperledger Fabric密码模块系列之BCCSP（五） - 国密算法实现](https://www.cnblogs.com/informatics/p/7648039.html)
- [fabric1.4.x国密改造过程全记录](https://blog.csdn.net/dyj5841619/article/details/90638054)