@echo off
set /p "toBuild=Enter ID: "

echo gradlew apps:server:%toBuild%:assemble
call gradlew apps:server:%toBuild%:assemble

echo opening project folder "apps/server/%toBuild%"
cd apps/server/%toBuild%

call docker images | findstr %toBuild%

set /p "msVersion=Enter version: "
call docker build . -t %toBuild%:%msVersion% && docker save %toBuild%:%msVersion% -o \\wsl.localhost\Ubuntu\home\ubuntu-user\wavy-27-03-23\%toBuild%.tar
cd ../../..

call scp \\wsl.localhost\Ubuntu\home\ubuntu-user\wavy-27-03-23\%toBuild%.tar dev@10.0.0.209:/home/dev/

echo Finished!
