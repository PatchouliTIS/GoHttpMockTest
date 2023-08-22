#!/bin/bash

gb -c 100 -n 100000 -k -T 'application/x-protobuf' -p 'GetGroupDetail.txt' -G 4 http://localhost:14000/ZeusService/GetGroupDetail > ./GetGroupDetailLog.log &
