#!/bin/bash

LOCAL_DIR="/mnt/sda2/Uniwizard/mailing_list/application-api"
REMOTE_DIR="/var/applications/mailing_list"

ssh root@172.16.10.10 "mkdir -p $REMOTE_DIR"
ssh root@172.16.10.10 "rm -rf $REMOTE_DIR/*"
ssh root@172.16.10.10 "mkdir -p $REMOTE_DIR/logs"

scp -r "$LOCAL_DIR/configs" root@172.16.10.10:"$REMOTE_DIR/"
scp "$LOCAL_DIR/application-api" root@172.16.10.10:"$REMOTE_DIR/"

ssh root@172.16.10.10 "supervisorctl reread"
ssh root@172.16.10.10 "supervisorctl update"
ssh root@172.16.10.10 "supervisorctl restart application-api-worker:*"

