#!/bin/bash

curl http://localhost:8080 &
sleep 0.5

# 2 - means CTRL+C, SIGINT
killall -2 server1

curl http://localhost:8080

killall -2 server2

