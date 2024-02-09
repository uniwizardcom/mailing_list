#!/bin/bash

apt-get update -y
apt-get upgrade -y

apt install postgresql postgresql-contrib

# Enable remote access to server
#
vi /etc/postgresql/xx/main/postgresql.conf
# Then, find the line #listen_addresses = 'localhost' and uncomment it (remove the # character at the beginning of the line).
# and set new value: listen_addresses = '*'
#
vi /etc/postgresql/xx/main/pg_hba.con
# # host    all             all             127.0.0.1/32            scram-sha-256
# host    all             all             0.0.0.0/0               md5
#
sudo ufw allow 5432/tcp


sudo -u postgres psql
# ALTER USER postgres PASSWORD 'postgres';
service postgresql restart

# create structures for application
sudo -u postgres psql
# CREATE DATABASE yourdbname;
# CREATE USER youruser WITH ENCRYPTED PASSWORD 'yourpass';
# GRANT ALL PRIVILEGES ON DATABASE yourdbname TO youruser;

# CREATE USER "mailing" WITH ENCRYPTED PASSWORD 'St1r.o2n,gP3a?5ssword';
# CREATE DATABASE "mailing";
# GRANT ALL PRIVILEGES ON DATABASE "mailing" TO "mailing";
