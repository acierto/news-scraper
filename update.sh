#!/bin/sh

find . -type d -d 1 -not \( -path "./.*" -o -path "./pkg" -o -path "./src" \) -exec rm -rf src/{} \;
find . -type d -d 1 -not \( -path "./.*" -o -path "./pkg" -o -path "./src" \) -exec cp -rf {} src \;
find . -type d -d 1 -not \( -path "./.*" -o -path "./pkg" -o -path "./src" \) -exec echo "Updated folder: src/{}" \;