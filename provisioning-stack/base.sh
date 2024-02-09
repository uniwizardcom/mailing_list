#!/bin/bash

apt-get update -y
apt-get upgrade -y
apt-get update -y
apt-get install curl gnupg apt-transport-https -y
reboot
