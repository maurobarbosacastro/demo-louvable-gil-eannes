#!/bin/bash
RED='\033[0;31m'
BLUE='\033[0;34m'
YELLOW='\033[0;35m'
CYAN='\033[0;36m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

echo "   _____       _ "
echo "  / ____|     | |  "
echo " | |  __  ___ | | __ _ _ __   __ _ "
echo " | | |_ |/ _ \| |/ _\` \| '_ \ / _\` |"
echo " | |__| | (_) | | (_| | | | | (_| |"
echo "  \_____|\___/|_|\__,_|_| |_|\__, |"
echo "                              __/ |"
echo "                             |___/ "


echo -e "\n\n${CYAN}>${NC} NX Generating @nx-go/nx-go:app\n\n"

printf "${GREEN}√${NC}"
read -p '  Enter service name in kebab case (e.g., change-me): ' serviceName
projectName=${serviceName}

projectName=$(echo "$serviceName" | sed 's/-/ /g;s/\<./\U&/g' | sed 's/ //g;s/\<./\U&/g')
projectName=${projectName}
echo "Setting project name to: $projectName..."

projectDir=$(echo "ms-$serviceName" | sed 's/-/-/g;s/\<./\L&/g')
echo "Setting project directory to ./apps/server/$projectDir..."

nx g @nx-go/nx-go:app "$projectDir" --directory "apps/server/$projectDir"

projectPath=./apps/server/"$projectDir"

rm ./apps/server/$projectDir/main.go
cp -r ./tools/scripts/dependencies/golang/. "$projectPath"/
mv ./apps/server/$projectDir/main_test.go ./apps/server/$projectDir/test/main_test.go
find "$projectPath"/ -type f -name "*.go" -exec sed -i '' -e "s|ms-changeme|$projectDir|g" {} +
# shellcheck disable=SC2059
printf "${GREEN}CREATE${NC}"
echo " apps/server/ms-$serviceName/Dockerfile"
sed -i '' "s|CHANGEME|$projectName|g" ./apps/server/$projectDir/Dockerfile
cd ./apps/server/$projectDir && go mod init "$projectDir"
# shellcheck disable=SC2059
printf "${GREEN}CREATE${NC}"
echo " apps/server/ms-$serviceName/go.mod"


###### DEPENDENCIES
# shellcheck disable=SC2059
printf "\n${GREEN}INSTALLING DEPENDENCIES .${NC}"
go get -u gorm.io/gorm  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get -u gorm.io/driver/postgres  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get github.com/google/uuid  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get github.com/labstack/echo/v4  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get github.com/golang-jwt/jwt  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get github.com/labstack/echo/v4/middleware  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get github.com/samber/lo@v1  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get github.com/Nerzal/gocloak/v13  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get github.com/joho/godotenv  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get github.com/rs/zerolog  > /dev/null 2>&1



###### SWAGGO
# shellcheck disable=SC2059
printf "\n\n${GREEN}OpenAPI initialization with Swaggo .${NC}"
go install github.com/swaggo/swag/cmd/swag  > /dev/null 2>&1
printf "${GREEN}.${NC}"
go get -u github.com/swaggo/echo-swagger  > /dev/null 2>&1
printf "\n"
swag init

printf "\n${GREEN}Update project.json to add new command openapi to generate OpenAPI docs .${NC}"
brew install jq  > /dev/null
printf "${GREEN}.${NC}"
jsonFile="project.json"
jq '.targets += {"openapi": {"command": "cd {projectRoot} && swag init && cd -"}}' "$jsonFile" > tmp.json && mv tmp.json "$jsonFile"
printf "${GREEN}.${NC}"


printf "\n\n${GREEN}Migration handling with Pressly/Goose${NC}\n"
go install github.com/pressly/goose/v3/cmd/goose@latest
source .env
printf "${GREEN}Update project.json to add new command migrate${NC}\n"
jq --arg projectName "$projectName" '.targets += {"migrate": { "executor": "nx:run-commands", "options": { "command": [ "cd {projectRoot} && GOOSE_DRIVER=postgres GOOSE_DBSTRING=\"user={args.user} password={args.password} dbname={args.db} sslmode={args.sslmode} host={args.host} port={args.port}\" goose -dir migrations up && cd -" ], "parallel": false }, "configurations": { "default": { "user": "postgres", "host": "localhost", "port": 5432, "password": "postgres", "db": $projectName, "sslmode": "disable" } }, "defaultConfiguration": "default" }}' "$jsonFile" > tmp.json && mv tmp.json "$jsonFile"
jq --arg projectName "$projectName" '.targets += {"create-migration": { "executor": "nx:run-commands", "options": { "command": [ "cd {projectRoot} && goose -dir migrations create {args.name} sql && cd -" ], "parallel": false } }}' "$jsonFile" > tmp.json && mv tmp.json "$jsonFile"

brew rm jq  > /dev/null
cd -  > /dev/null 2>&1


######.ENV UPDATE
# shellcheck disable=SC2059
printf "\n${GREEN}UPDATE .ENV${NC}\n"
nameUppercase=$(printf '%s\n' "$serviceName" | awk '{ print toupper($0) }')

#ADD DB VAR
LINE_NUMBER=$(grep -n "#DB" .env | cut -d: -f1)
DB_VARIABLE=$nameUppercase"_DB="$serviceName
sed -i '' "$((LINE_NUMBER))a\\
$DB_VARIABLE\\
" .env
echo "Added DB environment variable $DB_VARIABLE"

#ADD DB PORT
LINE_NUMBER=$(grep -n "#PORTS" .env | cut -d: -f1)
PORT_VARIABLE=$nameUppercase"_PORT=9999"
sed -i '' "$((LINE_NUMBER))a\\
$PORT_VARIABLE\\
" .env
echo "Added PORT environment variable $PORT_VARIABLE"

# Convert the provided service name to uppercase
SERVICE_NAME=$(echo "$projectName" | tr '[:lower:]' '[:upper:]')

GITLAB_CI_FILE="./.gitlab-ci.yml"
# Define the configuration to append
CONFIG="
#=== ms-$projectName ===
build-and-deploy:ms-$projectName:qa:
  <<: *atl-build-server-go
  rules:
    - if: \$CI_COMMIT_BRANCH == \$BRANCH_QA
      changes:
        - \$DIR_SERVER/\$SERVICE_MS_$SERVICE_NAME/**/*
      when: on_success
    - when: never
  variables:
    SERVICE: \$SERVICE_MS_$SERVICE_NAME
    ENV: \"QA\"
  tags:
    - \$TAG_ATL_BUILD_MS_$SERVICE_NAME

build-and-deploy:ms-$projectName:pre:
  <<: *atl-build-server-go
  rules:
    - if: \$CI_COMMIT_BRANCH == \$BRANCH_PRE
      changes:
        - \$DIR_SERVER/\$SERVICE_MS_$SERVICE_NAME/**/*
      when: on_success
    - when: never
  variables:
    SERVICE: \$SERVICE_MS_$SERVICE_NAME
    ENV: \"PRE\"
  tags:
    - \$TAG_ATL_BUILD_MS_$SERVICE_NAME

restart:pod:ms-$projectName:qa:
  <<: *atl-rollout-restart-deployment-qa
  rules:
    - if: \$CI_COMMIT_BRANCH == \$BRANCH_QA
      changes:
        - \$DIR_SERVER/\$SERVICE_MS_$SERVICE_NAME/**/*
      when: on_success
    - when: never
  variables:
    SERVICE: \$SERVICE_MS_$SERVICE_NAME
    ENV: \"QA\"

restart:pod:ms-$projectName:pre:
  <<: *atl-rollout-restart-deployment-pre
  rules:
    - if: \$CI_COMMIT_BRANCH == \$BRANCH_PRE
      changes:
        - \$DIR_SERVER/\$SERVICE_MS_$SERVICE_NAME/**/*
      when: on_success
    - when: never
  variables:
    SERVICE: \$SERVICE_MS_$SERVICE_NAME
    ENV: \"PRE\"
"

# Append configuration to the .gitlab-ci.yml file
echo "$CONFIG" >> "$GITLAB_CI_FILE"

# Print success message
printf "Configuration successfully added to $GITLAB_CI_FILE"

printf "\n\n"
cat <<EOF
✅ All setup - Don't forget to create your database ($projectName) in postgres

All information is in the README of $projectDir but here is a recap:

To create a new migration
    ▶ nx run $projectDir:create-migration --name=add-new-field
To run migrations:
    ▶ nx run $projectDir:migrate

To generate OpenAPI specs
    ▶ nx run $projectDir:openapi

To run the micro-service
    ▶ nx run $projectDir:serve

Before starting the ms, don't forget to check ports and DB names and edit the environment variable for GORM at '$projectDir/internal/config/db.go'
EOF

printf "\n\n"
echo "Happy Coding 🤓"
