#!/usr/bin/env bash

source ./arg.sh

# 客户端 ********************************
# 把最新版本软连到 stable
ln -snf "$WORKSPACE/client/release/$VERSION" "$WORKSPACE/client/stable"
# 更新 nginx 配置
rm -rf "$NGINX_CONFIG_PATH/$PROJECT_NAME.*"
\cp -rf "$WORKSPACE"/deploy/nginx/* "$NGINX_CONFIG_PATH"
nginx -s reload

# 服务端 ********************************
dist="$WORKSPACE/server/stable"
# 软连 stable
ln -snf "$WORKSPACE/server/release/$VERSION" "$dist"
server_name="$PROJECT_NAME"-server
# 停止之前的服务
pkill -f "$server_name"
cd "$dist" || exit
# 启动新的服务
nohup "./$server_name" -e prod >"$WORKSPACE/log/runtime.log" 2>&1 &
