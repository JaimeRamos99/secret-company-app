#!/bin/sh
sudo docker kill $(sudo docker ps -q)
sudo docker rm $(sudo docker ps -aq)
