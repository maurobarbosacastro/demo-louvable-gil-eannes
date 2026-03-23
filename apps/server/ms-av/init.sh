#!/bin/bash
echo "[init.sh] Starting mongod service..."
mongod --fork --logpath /var/log/mongod.log
#
#echo "[init.sh] Building mongo collection 'scanned'..."
#echo "[init.sh] Going mongoloid..."
#mongo
#  use admin
#  db.createCollection("scanned")
#  exit

ls -l /home/app/
echo "[init.sh] Starting MN app service..."
java -jar /home/app/application.jar
