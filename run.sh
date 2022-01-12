#!/usr/bin/env bash

CONFIG_ENV=$1

if [ ! -n "${CONFIG_ENV}" ]
then
    echo '请指定当前环境'
    echo '例如:'
    echo '    sh run.sh dev'
    echo '    sh run.sh test'
    echo '    sh run.sh prod'

else
    echo "now env is:${CONFIG_ENV}"

    if [ "${CONFIG_ENV}" == "dev" ]
    then
        DOCKER_IMAGE="shixinshuiyou/mayo:dev"
        ETCD_ADDR="host.docker.internal:2379"
    elif [ "${CONFIG_ENV}" == "test" ]
    then
        DOCKER_IMAGE="shixinshuiyou/mayo:test"
    elif [ "${CONFIG_ENV}" == "prod" ]
    then
        DOCKER_IMAGE="shixinshuiyou/mayo:prod"
    fi

    echo "now set docker image is:${DOCKER_IMAGE}"

    # 后台运行
        sudo docker run \
        --name shixinshuiyou-mayo-${CONFIG_ENV} \
        -h shixinshuiyou-mayo-${CONFIG_ENV} \
        -e TZ=Asia/Shanghai \
        -e ETCD_ADDR=${ETCD_ADDR} \
        -e RUN_MODE=${CONFIG_ENV} \
        -v `pwd`:/go/src/mayo \
        -v `pwd`/docker/${CONFIG_ENV}/supervisord.d:/etc/supervisord.d \
        --cap-add=SYS_PTRACE \
        --security-opt \
        seccomp=unconfined \
        --security-opt \
        apparmor=unconfined \
        -d \
        -p 9081-9090:9081-9090 \
        ${DOCKER_IMAGE} \
        supervisord -n && docker exec shixinshuiyou-mayo-${CONFIG_ENV} /bin/bash './reload.sh'

fi