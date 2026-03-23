@echo off
REM set /p "toBuild=Enter ID: "
set toBuild=ms-products
set msVersion=0.0.38
set prodFolder=wavy-prod-30-03-2023
REM set wavyKey=\\wsl.localhost\Ubuntu\home\ubuntu-user\wavy\Wavykey.txt
set wavyKey=dependencies\Wavykey.txt
set targetFile=\\wsl.localhost\Ubuntu\home\ubuntu-user\%prodFolder%\%toBuild%.tar

REM echo gradlew apps:server:%toBuild%:assemble
REM call gradlew apps:server:%toBuild%:assemble

REM echo opening project folder "apps/server/%toBuild%"
REM cd apps/server/%toBuild%

REM call docker images | findstr %toBuild%

REM set /p "msVersion=Enter version: "
REM call docker build . -t %toBuild%:%msVersion% && docker save %toBuild%:%msVersion% -o %targetFile%
REM cd ../../..
REM echo "Building and saving image"

echo Uploading image %toBuild%:%msVersion% to PRODUCTION server WAVY.CARE

docker save %toBuild%:%msVersion% -o %targetFile%

set prodTargetDir=/home/ubuntu/prod-deploy
REM set session=ubuntu@195.15.223.61
call scp -i %wavyKey% %targetFile% ubuntu@195.15.223.61:%prodTargetDir%

echo Finished sending the %toBuild% latest image to PROD server
