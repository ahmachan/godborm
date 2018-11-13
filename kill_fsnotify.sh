#!/bin/sh

CMD_ROOT=$(pwd)

ps|grep ${CMD_ROOT}/gopkg_fsnotify|awk '{ print $1}'|sudo xargs kill -9

