# 系统相关的util方法

_fn_is_windows() {
	# 98%的可能性
	kernel_info=$(uname -a)
	if [[ -e "/c/Windows" ]] && [[ ! $kernel_info =~ "Linux" ]]; then
		echo true
		return
	fi
	echo false
}

# 转换路径为win风格路径，以便win平台的特定工具能够识别，如 /d/User/Files/xxx.sh ==> d:\\User\\Files\\xxx.sh
_fn_convert_to_windows_path() {
	local path=$1
	if [[ -z $path ]]; then
		return
	fi
	replace_to_double_slash=$(echo $path | sed 's/[\\/]/\\\\/g')
	# 若前2个字符是\\，则去掉他们
	if [[ ${replace_to_double_slash::2} == '\\' ]]; then
		# 取第三个字符为盘符
		pan_char=${replace_to_double_slash:2:1}
		suffix=${replace_to_double_slash:3}
		echo "$pan_char:$suffix"
		return
	fi
	echo $replace_to_double_slash
}

if [ $# -ne 0 ]; then
	func_name="$1"
	${func_name} "${@:2}"
fi

__test__fn_is_windows() {
	echo "====== TEST ONLY ======"
	_fn_is_windows
}

__test__fn_convert_to_windows_path() {
	echo "====== TEST ONLY ======"
	#	_fn_convert_to_windows_path 'd:\xx\xxxx'
	_fn_convert_to_windows_path '/d/xx/xxxx'
}

#__test__fn_is_windows
#__test__fn_convert_to_windows_path
