#!/bin/bash

## Meant to be run as part of the release process, builds desktop apps

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

if command -v retry &> /dev/null
then
  retry -t 4 -- docker build -f tools/desktop/Dockerfile.desktop -t loadtoad .
else
  docker build -f tools/desktop/Dockerfile.desktop -t loadtoad .
fi


rm -rf tmp/release
mkdir -p tmp/release

cd "tmp/release"

id=$(docker create loadtoad)
docker cp $id:/dist - > ./desktop.tar
docker rm -v $id
tar -xvf "desktop.tar"
rm "desktop.tar"

mv dist/darwin_amd64/loadtoad "loadtoad.darwin"
mv dist/darwin_arm64/loadtoad "loadtoad.darwin.arm64"
mv dist/linux_amd64/loadtoad "loadtoad"
mv dist/windows_amd64/loadtoad "loadtoad.exe"
rm -rf "dist"

# darwin amd64
cp -R "../../tools/desktop/template" .

mkdir -p "./Load Toad.app/Contents/Resources"
mkdir -p "./Load Toad.app/Contents/MacOS"

cp -R "./template/darwin/Info.plist" "./Load Toad.app/Contents/Info.plist"
cp -R "./template/darwin/icons.icns" "./Load Toad.app/Contents/Resources/icons.icns"

cp "loadtoad.darwin" "./Load Toad.app/Contents/MacOS/loadtoad"

echo "signing amd64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Load Toad.app/Contents/MacOS/loadtoad"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Load Toad.app"

cp "./template/darwin/appdmg.config.json" "./appdmg.config.json"

echo "building macOS amd64 DMG..."
appdmg "appdmg.config.json" "./loadtoad_${TGT}_darwin_amd64_desktop.dmg"
zip -r "loadtoad_${TGT}_darwin_amd64_desktop.zip" "./Load Toad.app"

# darwin arm64
cp "loadtoad.darwin.arm64" "./Load Toad.app/Contents/MacOS/loadtoad"

echo "signing arm64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Load Toad.app/Contents/MacOS/loadtoad"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Load Toad.app"

echo "building macOS arm64 DMG..."
appdmg "appdmg.config.json" "./loadtoad_${TGT}_darwin_arm64_desktop.dmg"
zip -r "loadtoad_${TGT}_darwin_arm64_desktop.zip" "./Load Toad.app"

# macOS universal
rm "./Load Toad.app/Contents/MacOS/loadtoad"
lipo -create -output "./Load Toad.app/Contents/MacOS/loadtoad" loadtoad.darwin loadtoad.darwin.arm64

echo "signing universal desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Load Toad.app/Contents/MacOS/loadtoad"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Load Toad.app"

echo "building macOS universal DMG..."
appdmg "appdmg.config.json" "./loadtoad_${TGT}_darwin_all_desktop.dmg"
zip -r "loadtoad_${TGT}_darwin_all_desktop.zip" "./Load Toad.app"

# linux
echo "building Linux zip..."
zip "loadtoad_${TGT}_linux_amd64_desktop.zip" "./loadtoad"

#windows
echo "building Windows zip..."
curl -L -o webview.dll https://github.com/webview/webview/raw/master/dll/x64/webview.dll
curl -L -o WebView2Loader.dll https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll
zip "loadtoad_${TGT}_windows_amd64_desktop.zip" "./loadtoad.exe" "./webview.dll" "./WebView2Loader.dll"

mkdir -p "../../build/dist"
mv "./loadtoad_${TGT}_darwin_amd64_desktop.dmg" "../../build/dist"
mv "./loadtoad_${TGT}_darwin_amd64_desktop.zip" "../../build/dist"
mv "./loadtoad_${TGT}_darwin_arm64_desktop.dmg" "../../build/dist"
mv "./loadtoad_${TGT}_darwin_arm64_desktop.zip" "../../build/dist"
mv "./loadtoad_${TGT}_darwin_all_desktop.dmg" "../../build/dist"
mv "./loadtoad_${TGT}_darwin_all_desktop.zip" "../../build/dist"
mv "./loadtoad_${TGT}_linux_amd64_desktop.zip" "../../build/dist"
mv "./loadtoad_${TGT}_windows_amd64_desktop.zip" "../../build/dist"
