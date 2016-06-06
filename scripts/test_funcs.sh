#!/bin/bash

dbs="memcache mysql postgres"
tests="insert-by-pk select-by-pk update-by-pk delete-by-pk"
ops=50000
clients="1"

function simple_test_run {
    if [ -z "$1" ]
    then
        echo "Need to provide a database to test"
    else
    db=$1
    ops=$2
    host=$3
    port=$4
    user=$5
    conn=$6
    label=$7
    for t in $tests ; do
        # echo -n "DB, ${db}, Test, $t: "
        bin/kvbench -db ${db} -host ${host} -port ${port} -user=${user} \
            -num-connections ${conn} -num-operations ${ops} -test ${t} \
            -label ${label}
    done
    fi
}
