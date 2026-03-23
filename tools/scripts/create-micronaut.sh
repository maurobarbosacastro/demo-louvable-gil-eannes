#!/bin/bash
RED='\033[0;31m'
CYAN='\033[0;34m'
YELLOW='\033[0;35m'
BLUE='\033[0;37m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo "  __  __ _                                  _   "
echo " |  \/  (_)                                | |  "
echo " | \  / |_  ___ _ __ ___  _ __   __ _ _   _| |_ "
echo " | |\/| | |/ __| '__/ _ \| '_ \ / _\` | | | | __|"
echo " | |  | | | (__| | | (_) | | | | (_| | |_| | |_ "
echo " |_|  |_|_|\___|_|  \___/|_| |_|\__,_|\__,_|\__|"

echo -e "\n\n${CYAN}>${NC} NX Gencerating @nxrocks/nx-micronaut:application\n\n"

# shellcheck disable=SC2059
printf "${GREEN}√${NC}"
# shellcheck disable=SC2162
read -p '  Enter service name in kebab case (e.g., change-me): ' serviceName
projectName=${serviceName^}

# shellcheck disable=SC2059
printf "${GREEN}√${NC}"
# shellcheck disable=SC2162
read -p '  Append to docker-compose ? (y/n): ' printCompose

# shellcheck disable=SC2059
printf "${GREEN}√${NC}"
# shellcheck disable=SC2162
read -p '  Include readme? (y/n): ' isReadme

projectName=$(echo "$serviceName" | sed 's/-/ /g;s/\<./\U&/g' | sed 's/ //g;s/\<./\U&/g')
projectName=${projectName,}
echo "Setting project name to: $projectName..."

projectDir=$(echo "ms-$serviceName" | sed 's/-/-/g;s/\<./\L&/g')
echo "Setting project directory to ./apps/server/$projectDir..."

nx g @nxrocks/nx-micronaut:new "$projectDir" --type default --basePackage "pt.atlanse.$projectName" --buildSystem GRADLE --javaVersion JDK_17 --language GROOVY --micronautVersion current --directory server --testFramework SPOCK --skipFormat true --features ""

# DOMAIN CLASSES
echo "========================= // ========================="
echo "Enter the names of your domain classes (separated by spaces):"
read class_names
configDir="apps/server/ms-$serviceName/src/main/groovy/pt/atlanse/$projectName/configs"
controllerDir="apps/server/ms-$serviceName/src/main/groovy/pt/atlanse/$projectName/controllers"
serviceDir="apps/server/ms-$serviceName/src/main/groovy/pt/atlanse/$projectName/services"
repositoryDir="apps/server/ms-$serviceName/src/main/groovy/pt/atlanse/$projectName/repositories"
domainDir="apps/server/ms-$serviceName/src/main/groovy/pt/atlanse/$projectName/domains"
utilDir="apps/server/ms-$serviceName/src/main/groovy/pt/atlanse/$projectName/utils"
mkdir $domainDir $configDir $controllerDir $serviceDir $repositoryDir

# shellcheck disable=SC2059
printf "${GREEN}CREATE${NC}"
echo "$domainDir $configDir $controllerDir $serviceDir $repositoryDir"

### Create all domain classes inside Domains directory
#for name in $class_names; do
#    echo "Creating skeleton for $name.groovy"
#
#    # Create the new entity file from the template
#    cp ./tools/scripts/dependencies/groovy/TemplateEntity.groovy $domainDir/$name.groovy
#
#    # Replace the placeholder text with the actual entity name
#    sed -i "s|ENTITY_NAME|$name|g" $domainDir/$name.groovy
#    sed -i "s|PROJECT_NAME|$projectName|g" $domainDir/$name.groovy
#
#    echo "${GREEN}CREATE${NC} $name.groovy"
#
#done
#echo "Done creating domain classes."

# COPY DOCKER FILE TO THIS PROJECT
projectPath=./apps/server/"$projectDir"
cp ./tools/scripts/dependencies/Dockerfile-Micronaut "$projectPath"/Dockerfile
# shellcheck disable=SC2059
printf "${GREEN}CREATE${NC}"
echo " apps/server/ms-$serviceName/Dockerfile"

rm -r ./apps/server/"$projectDir"/gradle
# shellcheck disable=SC2059
printf "${RED}DELETE${NC}"
echo " apps/server/ms-$serviceName/gradle"

rm -r ./apps/server/"$projectDir"/gradle.properties
# shellcheck disable=SC2059
printf "${RED}DELETE${NC}"
echo " apps/server/ms-$serviceName/gradle.properties"

rm -r ./apps/server/"$projectDir"/gradlew
# shellcheck disable=SC2059
printf "${RED}DELETE${NC}"
echo " apps/server/ms-$serviceName/gradlew"

rm -r ./apps/server/"$projectDir"/gradlew.bat
# shellcheck disable=SC2059
printf "${RED}DELETE${NC}"
echo " apps/server/ms-$serviceName/gradlew.bat"

rm -r ./apps/server/"$projectDir"/settings.gradle
# shellcheck disable=SC2059
printf "${RED}DELETE${NC}"
echo " apps/server/ms-$serviceName/settings.gradle"

rm -r ./apps/server/"$projectDir"/micronaut-cli.yml
# shellcheck disable=SC2059
printf "${RED}DELETE${NC}"
echo " apps/server/ms-$serviceName/micronaut-cli.yml"

rm -r ./apps/server/"$projectDir"/.gitignore
# shellcheck disable=SC2059
printf "${RED}DELETE${NC}"
echo " apps/server/ms-$serviceName/.gitignore"

if [[ $isReadme == 'n' || $isReadme == 'n' ]]; then
    rm -r ./apps/server/"$projectDir"/README.md
    # shellcheck disable=SC2059
    printf "${RED}DELETE${NC}"
    echo " apps/server/ms-$serviceName/README.md"
fi

projectPath=./apps/server/"$projectDir"
oldName='"name": "server-'"$projectDir"'"'
projectJson=$(sed 's,'"$oldName"',"name": "'"$projectDir"'",' "$projectPath"/project.json)
rm "$projectPath"/project.json
echo "$projectJson" >>"$projectPath"/project.json

oldRoot='"root": "apps/server/'"$projectDir"'"'
targetRoot='"root": "."'
projectJson=$(sed 's,'"$oldRoot"','"$targetRoot"',' "$projectPath"/project.json)
rm "$projectPath"/project.json
echo "$projectJson" >>"$projectPath"/project.json

# shellcheck disable=SC2059
printf "${BLUE}UPDATE${NC}"
echo " apps/server/ms-$serviceName/project.json"

echo "include 'apps:server:$projectDir'" >>./settings.gradle
# shellcheck disable=SC2059
printf "${BLUE}UPDATE${NC}"
echo " settings.gradle"

if [[ $printCompose == 'y' || $printCompose == 'Y' ]]; then
    # shellcheck disable=SC2059
    printf "${BLUE}UPDATE${NC}"
    echo " ./docker-compose.yml"
    trigger="{INNER_TRIGGER_FOR_STARTER_SCRIPTS}"
    # shellcheck disable=SC2154
    testMe=$(echo $projectDir)
    composePartial=$(sed 's|'"$trigger"'|'"$testMe"'|' ./tools/scripts/dependencies/docker-compose-micronaut.yml | sed 's|'"$trigger"'|'"$testMe"'|')
    dockerComposeLocal=./docker-compose.yml
    partial=$(echo -e $composePartial)
    trigger2='# {TRIGGER_FOR_STARTER_SCRIPTS}'
    sed -i 's|'"$trigger2"'|'"$partial"'|' ./docker-compose.yml
elif [[ $printCompose == 'n' || $printCompose == 'N' ]]; then
    printf "${YELLOW}CANCELED${NC}"
    echo " ./docker-compose.yml File untouched..."
else
    echo -e '\nInvalid Input\n'
    echo -e '\nNo changes applied\n'
fi
