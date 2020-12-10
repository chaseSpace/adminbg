#!/usr/bin/env bash
call_echo_color='./bash_util/_echo_color.sh'
call_os_util='./bash_util/_os_util.sh'
chmod +x $call_echo_color $call_os_util

# 执行各种构建、安装、分析等操作的脚本
# 函数命名习惯：main.sh内定义且在main func内调用的func命名以 fn_ 开头, 其他func命名以 _fn_开头
# 1. 如何使用？
# -	看脚本结尾处说明
# 2. 想要修改脚本以扩展新功能？
# -	先了解脚本中函数的命名风格，标识“DO NOT NEED EDIT THIS FUNC”的函数是不用修改的
# -	如新增flag支持，只需要修改fn_init_cmd函数，新增逻辑代码建议新建文件，不要耦合在main.sh，请仔细阅读其中注释后再修改

flag=$1
project_dir=$2

# 默认以脚本执行目录为项目根目录
PROJECT_DIR=$(dirname $(pwd))
# 声明所有支持的命令数组
declare CMD_ARRAY=()

fn_init_cmd() {
	# ------------------- 所有的CMD选项(若添加新命令则需添加到这个数组) ----------------------
	readonly CMD_ARRAY=("gen" "gofmt" "govet")
	# xxx_cmd_on_ok后缀的指令 表示这个cmd指令执行成功后要继续执行的指令，类似的还有_on_fail,  _on_any
	# 新增命令，只需在这里定义即可，无需其他操作
	# 注意：这里定义的变量当做全局变量使用，请不要在此函数外定义xxx_cmd这样的全局变量，会干扰

	# gofmt
	readonly    gofmt_cmd="gofmt -l -s -w $PROJECT_DIR"
	# govet
	readonly    govet_cmd="go vet $PROJECT_DIR/..." # go vet可能会修改go.mod文件，执行tidy来恢复
	readonly    govet_cmd_on_ok="go mod tidy"

}

# DO NOT NEED EDIT THIS FUNC
_fn_execute() {
	local    cmd=$1
	local    cmd_on_ok=$2
	local    cmd_on_fail=$3
	local    cmd_on_any=$4

	echo    -e "EXECUTE> $cmd \n"
	echo    "******** output start *********"

	# 执行cmd对应指令
	$cmd
	local result_code=$?

	# 下面执行可能需要执行的附加指令

	$cmd_on_any    # 任何时候都需要执行的cmd

	if    [[ ${result_code} -eq 0 ]]; then
		$cmd_on_ok
		echo       -e "\n --- EXECUTE successful"
	else
		$cmd_on_fail
	fi

	echo    -e "\n******** output end *********"
}

# DO NOT NEED EDIT THIS FUNC
_fn_concat_if_not_empty() {
	local old=$1
	local mid=$2
	local append=$3

	if [[ -z "$append" ]]; then
		echo $old
		return
	fi
	echo "$old $mid $append\n"
}

# DO NOT NEED EDIT THIS FUNC
_fn_print_usage() {
	echo    "You need to provide flag to continue, see also below:"
	output="[Flag] <> [Cmd]\n"

	for flag in    "${CMD_ARRAY[@]}"; do
		cmd=$(      eval echo '$'"${flag}_cmd")
		cmd_on_ok=$(      eval echo '$'"${flag}_cmd_on_ok")
		cmd_on_fail=$(   eval echo '$'"${flag}_cmd_on_fail")
		cmd_on_any=$(   eval echo '$'"${flag}_cmd_on_any")

		# 注意双引号传参
		output=$(_fn_concat_if_not_empty "$output" "${flag} <>" "$cmd")
		output=$(_fn_concat_if_not_empty "$output" "---${flag}_cmd_on_ok <>" "$cmd_on_ok")
		output=$(_fn_concat_if_not_empty "$output" "---${flag}_cmd_on_fail <>" "$cmd_on_fail")
		output=$(_fn_concat_if_not_empty "$output" "---${flag}_cmd_on_any <>" "$cmd_on_any")
	done

	echo    -e $output | column -t -s "<>"
}

# DO NOT NEED EDIT THIS FUNC
fn_parse_flag() {
	if    [[ -z $flag  ]] || [[ $flag == "-h" ]]; then
		_fn_print_usage
		return
	fi

	local    will_do_cmd
	local    will_do_cmd_on_ok
	local    will_do_cmd_on_fail
	local    will_do_cmd_on_any

	will_do_cmd=$(   eval echo '$'"${flag}_cmd")

	if    [[ -z ${will_do_cmd} ]]; then
		echo       "Invalid flag:$flag"
		return
	fi

	# 获取变量的间接引用变量值
	will_do_cmd_on_ok=$(   eval echo '$'"${flag}_cmd_on_ok")
	will_do_cmd_on_fail=$(   eval echo '$'"${flag}_cmd_on_fail")
	will_do_cmd_on_any=$(   eval echo '$'"${flag}_cmd_on_any")
	#	echo $will_do_cmd 111
	#	echo $will_do_cmd_on_ok 222
	#	echo $will_do_cmd_on_fail 333
	#	echo $will_do_cmd_on_any 444

	# 调用执行方法(每个参数都含有空格，需要双引号包括)
	_fn_execute    "$will_do_cmd" "$will_do_cmd_on_ok" "$will_do_cmd_on_fail" "$will_do_cmd_on_any"
}

fn_init() {
	# 替换PROJECT_DIR
	if    [[ -n $project_dir     ]]; then
		PROJECT_DIR=$project_dir
	fi
	readonly    PROJECT_DIR

	$call_echo_color fn_echo_color_msg "> initial vars"
	_echo="
    EMPTY
    "
	$call_echo_color fn_echo_color_msg "$_echo"
}

main() {
	#	msg="***main.sh started***" # 问题：传入的前三个*全部丢失
	msg="---------- main.sh started ----------"
	$call_echo_color fn_echo_color_msg 'textcolor_red' "$msg"

	fn_init
	fn_init_cmd
	fn_parse_flag
}

# start
main

<<comment
cd /go/go-kit-examples/template/script
- Check usage:
./main.sh

- Usage examples:
./main.sh gofmt ../../   注意，为避免代码格式化的范围超出你的预期，可以用绝对路径指定，如/path/to/project_root
./main.sh govet ../../   这里也可以用绝对路径指定，比如/path/to/project_root

# 以上命令末尾也可不加路径
comment

# 脚本格式化
# cd go-util/tool
# ./shfmt.exe -s -w -i 4 -bn -ci -sr -kp ../script/main.sh

# 关于shfmt工具， https://github.com/mvdan/sh
