#!/usr/bin/env bash

# 解压缩并执行程序
# 假设压缩包在当前位置

source ./init.sh
if [[ ! $? -eq 0 ]]; then
	exit 1
fi

zip_filename="adminbg.zip"
c_go_bin_filename="adminbg_bin"
c_conf_filename="conf.$_init_env.yml"

# 3是忽略压缩时的三层目录结构，不能随意设置(更好的方式是压缩时不带目录结构)
tar zvxf $zip_filename --strip-components 3
if [[ ! $? -eq 0 ]]; then
	exit 1
fi

# kill old
pid=$(ps -ef | grep "$c_go_bin_filename" | grep -v 'grep' | awk '{printf $2" "}')
if [[ ! -z $pid ]]; then
	echo ">>kill所有正在运行的进程：kill -n 2 $pid"
	kill -n 2 $pid
fi

# run...
./$c_go_bin_filename -cfgFile $c_conf_filename
