#!/bin/bash
mkdir -p /data/db
touch /var/log/mongodb.log
chmod 777 /var/log/mongodb.log
/entrypoint.sh mongod --logpath /var/log/mongodb.log --logappend --replSet rs &
COUNTER=0
grep -q 'waiting for connections on port' /var/log/mongodb.log
while [[ $? -ne 0 && $COUNTER -lt 60 ]] ; do
    sleep 2
    let COUNTER+=2
    echo "Waiting for mongo to initialize... ($COUNTER seconds so far)"
    grep -q 'waiting for connections on port' /var/log/mongodb.log
done
if [ -f /data/db/.init_done ]; then
  exit 0
fi
mongo local --eval "rs.initiate()"
sleep 2
mongo local --eval "rs.conf()"
mongorestore --drop /home/dump
touch /data/db/.init_done
tail -f /dev/null
