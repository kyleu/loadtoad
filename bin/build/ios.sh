#!/bin/bash

## Builds the iOS framework and application

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

TGT=$1
[ "$TGT" ] || TGT="0.0.0"

echo "building gomobile for iOS..."
GOARCH=arm64 time gomobile bind -o build/dist/mobile_ios_arm64/loadtoadServer.xcframework -target=ios github.com/kyleu/loadtoad/app/cmd
echo "gomobile for iOS completed successfully, building distribution..."
cd "build/dist/mobile_ios_arm64/loadtoadServer.xcframework"
zip --symlinks -r "../../loadtoad_${TGT}_ios_framework.zip" .

echo "Building iOS app..."
cd "$dir/../../tools/ios"

rm -rf ../../build/dist/mobile_ios_app_arm64
mkdir -p ../../build/dist/mobile_ios_app_arm64

xcodegen generate --spec xcodegen.yml --project ../../build/dist/mobile_ios_app_arm64

mv Info.plist ../../build/dist/mobile_ios_app_arm64
cd ../../build/dist/mobile_ios_app_arm64

xcodebuild -project "Load Toad.xcodeproj" -allowProvisioningUpdates
zip -r "$dir/../../build/dist/loadtoad_${TGT}_ios_app.zip" .
