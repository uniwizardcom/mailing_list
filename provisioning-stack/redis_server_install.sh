#!/bin/bash

apt install redis-server
systemctl restart redis.service

# for required password
# vi /etc/redis/redis.conf

# bind 0.0.0.0 ::
# requirepass St1r.o2n,gP3a?5ssword
