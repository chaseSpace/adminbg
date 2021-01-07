#!/usr/bin/env bash

#
# 打包二进制和配置文件到压缩包，放到哪里都能运行！
# 示例：cd project_root/; sh ./script/build.sh
# 会得到一个根目录下的zip压缩包，把这个压缩包上传至其他环境，把run.sh也复制过去，执行run.sh即可

source ./script/init.sh
if [[ ! $? -eq 0 ]]; then
	exit 1
fi

c_go_bin_filename="adminbg_bin"

go build -ldflags "-s -w"  -o $c_go_bin_filename ./cmd/appbg/main.go

# 需压缩的文件名
c_conf_filename="$conf_filepath/conf.$_init_env.yml"

tar cvzf adminbg.zip $c_conf_filename $c_go_bin_filename

rm -f $c_go_bin_filename

# gofmt -l -s -w .
