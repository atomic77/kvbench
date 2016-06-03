#!/bin/bash

# Deploy the docker containers


# Running simple container pgsql:
docker run --name postgrestest -e POSTGRES_PASSWORD=pw1pw1pw1 -p 5432:5432 -d postgres


# Running simple mysql container:
docker run --name mysqltest -e MYSQL_ROOT_PASSWORD=pw1pw1pw1 -d -p 3306:3306 mysql:latest

# memcached
docker run --name memcachetst -d -p 11211:11211 memcached 


