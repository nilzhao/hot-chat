#!/usr/bin/env bash

# 参考 https://www.jianshu.com/p/ad05cffede0b
# 处理脚本参数
# -w 工作空间
# -v 版本好
while getopts ":w:v:" opt_name; do # 通过循环，使用 getopts，按照指定参数列表进行解析，参数名存入 opt_name
  case "$opt_name" in              # 根据参数名判断处理分支
  'w')                             # -w
    WORKSPACE="$OPTARG"            # 从 $OPTARG 中获取参数值
    ;;
  'v')
    VERSION="$OPTARG"
    ;;
  ?) # 其它未指定名称参数
    echo "Unknown argument(s)."
    exit 2
    ;;
  esac
done

# nginx 配置地址
NGINX_CONFIG_PATH="/etc/nginx/conf.d"
# 项目名称
PROJECT_NAME="hot-chat"
