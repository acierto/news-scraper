#!/bin/sh

chmod +x ./scripts/update.sh
./scripts/update.sh

npm install
node_modules/.bin/bower-installer bower.json

cp -rf src/main/web/libs web/libs
rm -rf src/main