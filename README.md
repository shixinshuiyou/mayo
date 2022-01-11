项目大体架构:
    go-micro v2 版本进行的微服务构造
    gin 作为web开发
    etcd 服务发现
    jaeger 作为链路追踪
    prometheus 日志收集以及机器管理
    K8S 容器分发
    rsync 代码同步（需要发布服务器支持）


## 首次执行脚本(sh run.sh dev)生成镜像容器

包目录结构：
-app    : 微服务总包
-bin    : 编译生成可以执行文件目录
-config : 代码配置中心
-docker : docker环境配置
-proto  : rpc-proto文件
-srv    : rpc-service 
-tool   : 自开发工具包，后续牵出