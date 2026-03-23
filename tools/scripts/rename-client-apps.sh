#!/bin/bash
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[0;35m'
CYAN='\033[0;36m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

function join_by {
    local d=${1-} f=${2-}
    if shift 2; then
        printf %s "$f" "${@/#/$d}"
    fi
}

#https://patorjk.com/software/taag/#p=display&f=ANSI%20Shadow&t=Rename%20Client%20Apps
echo " "
echo " ██████╗ ███████╗███╗   ██╗ █████╗ ███╗   ███╗███████╗     ██████╗██╗     ██╗███████╗███╗   ██╗████████╗     █████╗ ██████╗ ██████╗ ███████╗"
echo " ██╔══██╗██╔════╝████╗  ██║██╔══██╗████╗ ████║██╔════╝    ██╔════╝██║     ██║██╔════╝████╗  ██║╚══██╔══╝    ██╔══██╗██╔══██╗██╔══██╗██╔════╝"
echo " ██████╔╝█████╗  ██╔██╗ ██║███████║██╔████╔██║█████╗      ██║     ██║     ██║█████╗  ██╔██╗ ██║   ██║       ███████║██████╔╝██████╔╝███████╗"
echo " ██╔══██╗██╔══╝  ██║╚██╗██║██╔══██║██║╚██╔╝██║██╔══╝      ██║     ██║     ██║██╔══╝  ██║╚██╗██║   ██║       ██╔══██║██╔═══╝ ██╔═══╝ ╚════██║"
echo " ██║  ██║███████╗██║ ╚████║██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╗███████╗██║███████╗██║ ╚████║   ██║       ██║  ██║██║     ██║     ███████║"
echo " ╚═╝  ╚═╝╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝╚══════╝╚═╝╚══════╝╚═╝  ╚═══╝   ╚═╝       ╚═╝  ╚═╝╚═╝     ╚═╝     ╚══════╝"
echo ""

paramName=$1

if [ $# -ne 0 ]; then
    printf "${GREEN}√${NC} "
    echo "Name given: $paramName"
else
    printf "${GREEN}√${NC} "
    read -p "No name given, enter one : " paramName
fi

IN=$paramName
lowerCase=$(echo $IN | tr '[:upper:]' '[:lower:]')
arrIN=(${lowerCase// / })
arrINCase=(${IN// / })

snakeCase=$(join_by _ "${arrIN[@]}")
kebabCase=$(join_by - "${arrIN[@]}")

titleCaseVar=()
for i in "${arrINCase[@]}"; do
    foo="$(tr '[:lower:]' '[:upper:]' <<<${i:0:1})${i:1}"
    titleCaseVar+=("$foo")
done

titleCase=$(join_by ' ' "${titleCaseVar[@]}")
noSpace=$(join_by '' "${arrIN[@]}")


echo -e "\n"
printf "${BLUE}UPDATE - MOBILE${NC}"


defaultSnakeCase="change_me_mobile"
echo -e "\n\n${GREEN}Android folder name change to match package name${NC}"
cp -R "./apps/client/mobile/android/app/src/main/kotlin/pt/atlanse/$defaultSnakeCase" "./apps/client/mobile/android/app/src/main/kotlin/pt/atlanse/$snakeCase"
rm -rf "./apps/client/mobile/android/app/src/main/kotlin/pt/atlanse/$defaultSnakeCase"

echo -e "\n\n${GREEN}Snake case subsitutions: $defaultSnakeCase -> $snakeCase${NC}"
printf "${GREEN}√${NC} "
echo "MainActivity.kt"
path="./apps/client/mobile/android/app/src/main/kotlin/pt/atlanse/$snakeCase"
data=$(sed 's,'"$defaultSnakeCase"','"$snakeCase"',' "$path/MainActivity.kt")
rm "$path"/MainActivity.kt
echo "$data" >> "$path"/MainActivity.kt

printf "${GREEN}√${NC} "
echo "AndroidManifest.xml"
path="./apps/client/mobile/android/app/src/main"
data=$(sed 's,'"$defaultSnakeCase"','"$snakeCase"',' "$path/AndroidManifest.xml")
rm "$path"/AndroidManifest.xml
echo "$data" >> "$path"/AndroidManifest.xml

printf "${GREEN}√${NC} "
echo "Info.plist"
path="./apps/client/mobile/ios/Runner"
data=$(sed 's,'"$defaultSnakeCase"','"$snakeCase"',' "$path/Info.plist")
rm "$path"/Info.plist
echo "$data" >> "$path"/Info.plist

printf "${GREEN}√${NC} "
echo "pubspec.yaml"
path="./apps/client/mobile"
data=$(sed 's,'"$defaultSnakeCase"','"$snakeCase"',' "$path/pubspec.yaml")
rm "$path"/pubspec.yaml
echo "$data" >> "$path"/pubspec.yaml

printf "${GREEN}√${NC} "
echo "widget_test.dart"
path="./apps/client/mobile/test"
data=$(sed 's,'"$defaultSnakeCase"','"$snakeCase"',' "$path/widget_test.dart")
rm "$path"/widget_test.dart
echo "$data" >> "$path"/widget_test.dart

printf "${GREEN}√${NC} "
echo "manifest.json"
path="./apps/client/mobile/web"
data=$(sed 's,'"$defaultSnakeCase"','"$snakeCase"',' "$path/manifest.json")
rm "$path"/manifest.json
echo "$data" >> "$path"/manifest.json


defaultSpaceCase="change me mobile"
echo -e "\n\n${GREEN}Pascal case subsitutions: $defaultSpaceCase -> $titleCase${NC}"

printf "${GREEN}√${NC} "
echo "pubspec.yaml"
path="./apps/client/mobile"
data=$(sed 's,'"$defaultSpaceCase"','"$titleCase"',' "$path/pubspec.yaml")
rm "$path"/pubspec.yaml
echo "$data" >> "$path"/pubspec.yaml

printf "${GREEN}√${NC} "
echo "Info.plist"
path="./apps/client/mobile/ios/Runner"
data=$(sed 's,'"$defaultSpaceCase"','"$titleCase"',' "$path/Info.plist")
rm "$path"/Info.plist
echo "$data" >> "$path"/Info.plist

printf "${GREEN}√${NC} "
echo "manifest.json"
path="./apps/client/mobile/web"
data=$(sed 's,'"$defaultSpaceCase"','"$titleCase"',' "$path/manifest.json")
rm "$path"/manifest.json
echo "$data" >> "$path"/manifest.json

printf "${GREEN}√${NC} "
echo "strings.dart"
path="./apps/client/mobile/lib/utils/constants"
data=$(sed 's,'"$defaultSpaceCase"','"$titleCase"',' "$path/strings.dart")
rm "$path"/strings.dart
echo "$data" >> "$path"/strings.dart


defaultNoSpace="changememobile"
echo -e "\n\n${GREEN}No space case subsitutions: $defaultNoSpace -> $noSpace${NC}"

printf "${GREEN}√${NC} "
echo "AndroidManifest.xml"
path="./apps/client/mobile/android/app/src/main"
data=$(sed 's,'"$defaultNoSpace"','"$noSpace"',' "$path/AndroidManifest.xml")
rm "$path"/AndroidManifest.xml
echo "$data" >> "$path"/AndroidManifest.xml

printf "${GREEN}√${NC} "
echo "AndroidManifest.xml debug"
path="./apps/client/mobile/android/app/src/debug"
data=$(sed 's,'"$defaultNoSpace"','"$noSpace"',' "$path/AndroidManifest.xml")
rm "$path"/AndroidManifest.xml
echo "$data" >> "$path"/AndroidManifest.xml

printf "${GREEN}√${NC} "
echo "AndroidManifest.xml profile..."
path="./apps/client/mobile/android/app/src/profile"
data=$(sed 's,'"$defaultNoSpace"','"$noSpace"',' "$path/AndroidManifest.xml")
rm "$path"/AndroidManifest.xml
echo "$data" >> "$path"/AndroidManifest.xml

printf "${GREEN}√${NC} "
echo "strings.dart"
path="./apps/client/mobile/lib/utils/constants"
data=$(sed 's,'"$defaultNoSpace"','"$noSpace"',' "$path/strings.dart")
rm "$path"/strings.dart
echo "$data" >> "$path"/strings.dart

printf "${GREEN}√${NC} "
echo "project.pbxproj"
path="./apps/client/mobile/ios/Runner.xcodeproj"
data=$(sed 's,'"$defaultNoSpace"','"$noSpace"',' "$path/project.pbxproj")
rm "$path"/project.pbxproj
echo "$data" >> "$path"/project.pbxproj

printf "${GREEN}√${NC} "
echo "build.gradle"
path="./apps/client/mobile/android/app"
data=$(sed 's,'"$defaultNoSpace"','"$noSpace"',' "$path/build.gradle")
rm "$path"/build.gradle
echo "$data" >> "$path"/build.gradle


echo -e "\n\n"
printf "${BLUE}UPDATE - BACKOFFICE${NC}"
defaultBO="change me bo"
echo -e "\n\n${GREEN}Pascal case subsitutions: $defaultBO -> $titleCase${NC}"

printf "${GREEN}√${NC} "
echo "index.html"
path="./apps/client/backoffice/src"
data=$(sed 's,'"$defaultBO"','"$titleCase"',' "$path/index.html")
rm "$path"/index.html
echo "$data" >> "$path"/index.html


echo -e "\n"
echo Terminated...
