# MaYO

    个人测试go-micro架构

## 项目大体架构

    go-micro v2 版本进行的微服务构造
    gin 作为web开发
    etcd 服务发现
    jaeger 作为链路追踪
    prometheus 日志收集以及机器管理
    K8S 容器分发
    rsync 代码同步（需要发布服务器支持）

## 启动

    生成镜像（docker build -f docker/base.dockerfile -t shixinshuiyou/mayo:dev .）
    执行脚本(sh run.sh dev)生成镜像容器
    重新编译:容器根目录内 sh reload.sh

## 包目录结构

    -app    : 微服务总包
    -bin    : 编译生成可以执行文件目录
    -config : 代码配置中心
    -docker : docker环境配置
    -proto  : rpc-proto文件
    -srv    : rpc-service
    -tool   : 自开发工具包，后续牵出

## 参考文章

    go-micro-api文章 <https://zhuanlan.zhihu.com/p/252824566>
    jaeger文章 <https://kebingzao.com/2020/12/25/jaeger-use/>
    opentracing介绍 https://xiaoming.net.cn/2021/03/25/Opentracing%E6%A0%87%E5%87%86%E5%92%8CJaeger%E5%AE%9E%E7%8E%B0/
    go-micro v2升级v3 https://micro.dev/v2-to-v3-upgrade-guide
