#!/bin/bash

## Meant to be run as part of the release process, builds desktop apps

set -eo pipefail
dir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$dir/../.."

TGT=$1
[ "$TGT" ] || TGT="v0.0.0"

if command -v retry &> /dev/null
then
  retry -t 4 -- docker build -f tools/desktop/Dockerfile.desktop -t solitaire .
else
  docker build -f tools/desktop/Dockerfile.desktop -t solitaire .
fi


rm -rf tmp/release
mkdir -p tmp/release

cd "tmp/release"

id=$(docker create solitaire)
docker cp $id:/dist - > ./desktop.tar
docker rm -v $id
tar -xvf "desktop.tar"
rm "desktop.tar"

mv dist/darwin_amd64/solitaire "solitaire.darwin"
mv dist/darwin_arm64/solitaire "solitaire.darwin.arm64"
mv dist/linux_amd64/solitaire "solitaire"
mv dist/windows_amd64/solitaire "solitaire.exe"
rm -rf "dist"

# darwin amd64
cp -R "../../tools/desktop/template" .

mkdir -p "./Solitaire.app/Contents/Resources"
mkdir -p "./Solitaire.app/Contents/MacOS"

cp -R "./template/darwin/Info.plist" "./Solitaire.app/Contents/Info.plist"
cp -R "./template/darwin/icons.icns" "./Solitaire.app/Contents/Resources/icons.icns"

cp "solitaire.darwin" "./Solitaire.app/Contents/MacOS/solitaire"

echo "signing amd64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Solitaire.app/Contents/MacOS/solitaire"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Solitaire.app"

cp "./template/darwin/appdmg.config.json" "./appdmg.config.json"

echo "building macOS amd64 DMG..."
appdmg "appdmg.config.json" "./solitaire_${TGT}_darwin_amd64_desktop.dmg"
zip -r "solitaire_${TGT}_darwin_amd64_desktop.zip" "./Solitaire.app"

# darwin arm64
cp "solitaire.darwin.arm64" "./Solitaire.app/Contents/MacOS/solitaire"

echo "signing arm64 desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Solitaire.app/Contents/MacOS/solitaire"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Solitaire.app"

echo "building macOS arm64 DMG..."
appdmg "appdmg.config.json" "./solitaire_${TGT}_darwin_arm64_desktop.dmg"
zip -r "solitaire_${TGT}_darwin_arm64_desktop.zip" "./Solitaire.app"

# macOS universal
rm "./Solitaire.app/Contents/MacOS/solitaire"
lipo -create -output "./Solitaire.app/Contents/MacOS/solitaire" solitaire.darwin solitaire.darwin.arm64

echo "signing universal desktop binary..."
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Solitaire.app/Contents/MacOS/solitaire"
codesign -f --options=runtime --verbose=4 --deep --force --strict -s 'Developer ID Application: Kyle Unverferth (C6S478FYLD)' "./Solitaire.app"

echo "building macOS universal DMG..."
appdmg "appdmg.config.json" "./solitaire_${TGT}_darwin_all_desktop.dmg"
zip -r "solitaire_${TGT}_darwin_all_desktop.zip" "./Solitaire.app"

# linux
echo "building Linux zip..."
zip "solitaire_${TGT}_linux_amd64_desktop.zip" "./solitaire"

#windows
echo "building Windows zip..."
curl -L -o webview.dll https://github.com/webview/webview/raw/master/dll/x64/webview.dll
curl -L -o WebView2Loader.dll https://github.com/webview/webview/raw/master/dll/x64/WebView2Loader.dll
zip "solitaire_${TGT}_windows_amd64_desktop.zip" "./solitaire.exe" "./webview.dll" "./WebView2Loader.dll"

mkdir -p "../../build/dist"
mv "./solitaire_${TGT}_darwin_amd64_desktop.dmg" "../../build/dist"
mv "./solitaire_${TGT}_darwin_amd64_desktop.zip" "../../build/dist"
mv "./solitaire_${TGT}_darwin_arm64_desktop.dmg" "../../build/dist"
mv "./solitaire_${TGT}_darwin_arm64_desktop.zip" "../../build/dist"
mv "./solitaire_${TGT}_darwin_all_desktop.dmg" "../../build/dist"
mv "./solitaire_${TGT}_darwin_all_desktop.zip" "../../build/dist"
mv "./solitaire_${TGT}_linux_amd64_desktop.zip" "../../build/dist"
mv "./solitaire_${TGT}_windows_amd64_desktop.zip" "../../build/dist"
