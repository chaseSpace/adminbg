#!/usr/bin/env bash

# readonly const
readonly textcolor_black=30
readonly textcolor_red=31
readonly textcolor_green=32
readonly textcolor_yellow=33
readonly textcolor_blue=34
readonly textcolor_fuchsia=35 # 紫红色
readonly textcolor_cyan=36 # 青色
readonly textcolor_white=37

readonly bgcolor_black=40
readonly bgcolor_red=41
readonly bgcolor_green=42
readonly bgcolor_yellow=43
readonly bgcolor_blue=44
readonly bgcolor_fuchsia=45  # 紫红色
readonly bgcolor_cyan=46 # 青色
readonly bgcolor_white=47
readonly bgcolor_default=49

fn_display_method() {
	local   color_id=$1
	local   bgcolor_id=$2
	local   msg=$3
	echo -e "\033[${color_id};${bgcolor_id}m $msg \033[0m"
}

fn_get_checked_color() {
	# 验证传入的color参数
	local color_arg=$(eval echo '$'$1)
	if [[ -z "$color_arg" ]]; then
		return
	fi
	# 通过echo返回 (运行时不会打印)
	echo $color_arg
}

fn_get_checked_bgcolor() {
	# 验证传入的bgcolor参数
	local _bgcolor_arg=$(eval echo '$'$1)
	if [[ -z "$_bgcolor_arg" ]]; then

		return
	fi
	# 通过echo返回 (运行时不会打印)
	echo $_bgcolor_arg
}

fn_echo_color_msg() {
	# default
	local color_id=$textcolor_red
	local bgcolor_id=$bgcolor_default
	local color_arg='textcolor_red'

	local message

	#	echo "color_id ${1} $color_id"

	# 仅提供1个参数
	if [[ -z "$2" && $1 ]]; then
		message=$1

	# 仅提供3个参数
	elif [[ -z "$4" && $1 && $2 && $3 ]]; then
		message=$3
		color_arg=$1
		local _bgcolor_arg=$2

		color_id=$(fn_get_checked_color $color_arg)
		if [ -z "$color_id" ]; then
			echo "Invalid color_id arg"
			exit 1
		fi

		bgcolor_id=$(fn_get_checked_bgcolor $_bgcolor_arg)
		if [ -z "$bgcolor_id" ]; then
			echo "Invalid bg color_id arg"
			exit 1
		fi

	# 仅提供2个参数
	elif [[ -z "$3" && $1 && $2 ]]; then
		message=$2
		color_arg=$1

		color_id=$(  fn_get_checked_color $color_arg)
		if [ -z "$color_id" ]; then
			echo "Invalid color_id arg"
			exit 1
		fi

	else
		echo "fn_echo_color_msg: not allowed args"
		exit 1
	fi

	#	echo $color_arg $color_id ${bgcolor_id} ${message}
	fn_display_method      ${color_id} ${bgcolor_id} "$message" # 双引号传参，避免空格导致的参数分割问题
}

if [ $# -ne 0 ]; then
	fn_name="$1"
	${fn_name} "${@:2}"
fi

# comment
<<EOF
# example-1, default print red text, white background
imported_fn_echo_color_msg xxx

# example-2, set textcolor, see also "readonly const"
imported_fn_echo_color_msg 'textcolor_green' xxx

# example-3, set textcolor, bgcolor
msg="with space"  # if msg contains space, you must declare msg as a var, then pass it use "$var".
imported_fn_echo_color_msg 'textcolor_red' 'bgcolor_yellow'  "$msg"
EOF
