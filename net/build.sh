#!/usr/bin/env bash

GOOS=linux go build -o tcp_server tcp_server.go


echo "begin upload file"
scp tcp_server root@39.107.111.1:/root/project/tcp/tcp_server1
echo "file uploaded successfully"


ssh root@39.107.111.1 sh /root/project/tcp/kill_process.sh
echo "killed process"


ssh root@39.107.111.1 sh /root/project/tcp/start_process.sh
echo "process started successfully"
