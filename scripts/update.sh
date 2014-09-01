#!/bin/bash

ignored_folders=()

while read line
do
    ignored_folders+=("$line")
done < .gitignore

ignored_folders+=(.)
ignored_folders+=(build/)
ignored_folders+=(test/)
ignored_folders+=(scripts/)
ignored_folders+=(web/)

for i in `find . -type d -maxdepth 1`
do
    fold=${i:2}/
    if ! [[ "${ignored_folders[@]}" =~ "${fold} " || "${ignored_folders[${#ignored_folders[@]}-1]}" == "${fold}" ]]; then
        echo "Not ignored: $i"
        rm -rf src/$i && cp -rf $i src
    fi
done