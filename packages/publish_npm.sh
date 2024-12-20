#!/bin/bash

set -e

# Mac
mkdir -p ./packages/npm-mcpmock-darwin-x64/bin
cp dist/mcpmock_darwin_amd64_v1/mcpmock ./packages/npm-mcpmock-darwin-x64/bin/mcpmock
chmod +x ./packages/npm-mcpmock-darwin-x64/bin/mcpmock
mkdir -p ./packages/npm-mcpmock-darwin-arm64/bin
cp dist/mcpmock_darwin_arm64_v8.0/mcpmock ./packages/npm-mcpmock-darwin-arm64/bin/mcpmock
chmod +x ./packages/npm-mcpmock-darwin-arm64/bin/mcpmock

# Linux
mkdir -p ./packages/npm-mcpmock-linux-x64/bin
cp dist/mcpmock_linux_amd64_v1/mcpmock ./packages/npm-mcpmock-linux-x64/bin/mcpmock
chmod +x ./packages/npm-mcpmock-linux-x64/bin/mcpmock
mkdir -p ./packages/npm-mcpmock-linux-arm64/bin
cp dist/mcpmock_linux_arm64_v8.0/mcpmock ./packages/npm-mcpmock-linux-arm64/bin/mcpmock
chmod +x ./packages/npm-mcpmock-linux-arm64/bin/mcpmock

# Windows
mkdir -p ./packages/npm-mcpmock-win32-x64/bin
cp dist/mcpmock_windows_amd64_v1/mcpmock.exe ./packages/npm-mcpmock-win32-x64/bin/mcpmock.exe
mkdir -p ./packages/npm-mcpmock-win32-arm64/bin
cp dist/mcpmock_windows_arm64_v8.0/mcpmock.exe ./packages/npm-mcpmock-win32-arm64/bin/mcpmock.exe

cd packages/npm-mcpmock-darwin-x64
npm publish --access public

cd ../npm-mcpmock-darwin-arm64
npm publish --access public

cd ../npm-mcpmock-linux-x64
npm publish --access public

cd ../npm-mcpmock-linux-arm64
npm publish --access public

cd ../npm-mcpmock-win32-x64
npm publish --access public

cd ../npm-mcpmock-win32-arm64
npm publish --access public

cd ../npm-mcpmock
npm publish --access public

cd -