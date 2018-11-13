#!/bin/sh

CMD_ROOT=$(pwd)
PROJECT_ROOT=$(pwd)

${CMD_ROOT}/gopkg_fsnotify -f ${PROJECT_ROOT}/main.go -o ${PROJECT_ROOT}/test.bin

