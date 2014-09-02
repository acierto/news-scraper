#!/bin/sh

function build() {
    sudo chmod +x build.sh
    ./build.sh
}

function restartServer() {
    used_pid=`ps -aef | grep news-scraper | grep -v grep | awk '{print $2}'`

    if [[ ! -z $used_pid ]]
    then
        sudo kill -9 $used_pid
    fi

    nohup ./news-scraper &
}

build
restartServer