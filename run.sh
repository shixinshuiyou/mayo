#!/usr/bin/env bash

CONFIG_ENV = $1

if [ ! -n "${CONFIG_ENV}" ]
then
    echo '请指定当前环境'
    echo '例如：'
    echo '    sh run.sh dev'
    echo '    sh run.sh test'
    echo '    sh run.sh prod'

else
    echo "now env is:${CONFIG_ENV}"

    if [ "${CONFIG_ENV}" == "dev" ]
    then
        OUT_ADDR =  ":9001"
    elif [ "${CONFIG_ENV}" == "test" ]
    then
        echo "this env not complete"

    fi

    sudo go run app/user/main.go -service_address=${OUT_ADDR}