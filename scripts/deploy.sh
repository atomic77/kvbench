#!/bin/bash

# Deploy the docker containers

# docker-machien we're using for this test
machine="aws02"
pw="pw1pw1pw1"
source test_funcs.sh


# Running simple container pgsql:
docker run --name postgrestest -e POSTGRES_PASSWORD=${pw} -p 5432:5432 -d postgres


# Running simple mysql container:
docker run --name mysqltest -e MYSQL_ROOT_PASSWORD=${pw} -d -p 3306:3306 mysql:latest

# memcached
docker run --name memcachetst -d -p 11211:11211 memcached 

# innodb-memcached
docker run --name innodb-memcache -e MYSQL_ROOT_PASSWORD=${pw} -d -p 11212:11211 -p 3307:3306 mysql:latest

ip=$(docker-machine ip ${machine})
# TODO: Get this working with the docker-entrypoint-initdb stuff
# the base mysql image provides
docker run --rm -i -t --link innodb-memcache:db mysql:latest mysql -u root -p${pw} -h db -e "source /usr/share/mysql/innodb_memcached_config.sql"
docker run --rm -i -t --link innodb-memcache:db mysql:latest mysql -u root -p${pw} -h db -e "INSTALL PLUGIN daemon_memcached soname "libmemcached.so";"
