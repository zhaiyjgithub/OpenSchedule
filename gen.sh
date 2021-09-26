#!/usr/bin/env bash

 ~/go/bin/gen --connstr "root:123456@tcp(127.0.0.1:3306)/drfinder?&parseTime=True" --database drfinder -t doctors --model generateModel --gorm