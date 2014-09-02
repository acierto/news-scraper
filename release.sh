#!/bin/sh

function build() {
    chmod +x build.sh
    ./build.sh
}

function restartServer() {
    used_pid=`ps -aef | grep news-scraper | grep -v grep | awk '{print $2}'`

    if [[ ! -z $used_pid ]]
    then
        kill -9 $used_pid
    fi

    nohup ./news-scraper &
}

build
restartServer