#!/usr/bin/env bash

_init_env=$1

conf_filepath="./deploy/config"

if [[ -z $_init_env ]]; then
	_init_env="dev"
fi

echo "初始化 -- 环境:${_init_env}"
case $_init_env in
	"dev") ;;

	"test") ;;

	"prod")
		export GIN_MODE=release
		;;
	*)
		echo "未知环境: $_init_env"
		exit 1
		;;
esac
