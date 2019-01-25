#!/bin/sh

function do_start() {
    nohup ./journey --log logs/access.log &
}

function show_usage() {
    echo "Usage: sh control.sh <start|stop|restart>"
}

function do_stop() {
    ps -ef | grep journey | awk '{print $2;}' | xargs -i kill -9 {}
}

if [[ $# -lt 1 ]]
then
    show_usage
    exit 1
fi

op=$1

if [[ $op == 'start' ]]
then
    do_start
elif [[ $op == 'stop' ]]
then
    do_stop
elif [[ $op == 'restart' ]]
then
    do_start
    do_stop
else
    show_usage
fi