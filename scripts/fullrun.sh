#!/bin/bash

total_ops=160000
runs=3
. scripts/test_funcs.sh

# Memcache
docker start memcachetst
sleep 5

full_test memcache ${total_ops} 11211 asdf memcache_aws_t2med memcachetst $runs > logs/memcache_aws02.log
docker stop memcachetst

# MySql with sql and memcache interface
docker start innodb-memcache
sleep 5

full_test memcache ${total_ops} 11211 asdf innomemcache_aws_t2med innodb-memcache $runs > logs/memcache_inno_aws02.log

#sleep 5
# Run against the same mysql
full_test mysql ${total_ops} 3306 root mysql_aws_t2med innodb-memcache $runs > logs/mysql_aws02.log

docker stop innodb-memcache

# Pgsql for comparison sake

docker start postgrestest
sleep 5

full_test postgres ${total_ops} 5432 postgres postgres_aws_t2med postgrestest $runs > logs/pgsql_aws02.log

docker stop postgrestest
