#!/bin/bash

LOCAL_DIR=pwd
REMOTE_DIR="/var/applications/mailing_list"

ssh root@172.16.10.10 "mkdir -p $REMOTE_DIR"
ssh root@172.16.10.10 "rm -rf $REMOTE_DIR/*"


echo $LOCAL_DIR
# scp root@172.16.10.10

# application_api
