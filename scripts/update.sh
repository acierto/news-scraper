#!/bin/bash

ignored_folders=()

while read line
do
    ignored_folders+=("$line")
done < .gitignore

ignored_folders+=(web/)
ignored_folders+=(test/)

for i in `find . -type d -d 1`
do
    fold=${i:2}/
    if ! [[ "${ignored_folders[@]}" =~ "${fold} " || "${ignored_folders[${#ignored_folders[@]}-1]}" == "${fold}" ]]; then
        echo "Not ignored: $i"
        rm -rf src/$i && cp -rf $i src
    fi
done