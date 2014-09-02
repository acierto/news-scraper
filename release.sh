#!/bin/sh

function build() {
    chmod +x build.sh
    ./build.sh

    export GOPATH=`pwd`

    go get github.com/go-martini/martini
    go get github.com/PuerkitoBio/goquery
    go get github.com/robfig/cron

    go build

    executed_file_name=${PWD##*/}
    mv $executed_file_name news-scraper
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