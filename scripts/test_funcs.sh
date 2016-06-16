#!/bin/bash

tests="insert-by-pk select-by-pk update-by-pk delete-by-pk"
conns="1 2 4 8 16"
basecmd="docker run --rm -i -t "
kvimage="atomic77/kvbench"
# linkcmd="--link postgrestest:pgtst"
# atomic77/kvbench -db postgres -user postgres -host pgtst -port 5432 -num-connections 1 -num-operations 10 -test insert-by-pk
# Example full command:
# docker run --rm -i -t --link postgrestest:pgtst atomic77/kvbench
#      -db postgres -user postgres -host pgtst -port 5432 -num-connections 4
#       -num-operations 1000 -test select-by-pk


function simple_test_run {
    if [ -z "$1" ]
    then
        echo "Need to provide a database to test"
    else
    db=$1
    ops=$2
    port=$3
    user=$4
    conn=$5
    label=$6
    container=$7
    for t in $tests ; do
        # echo -n "DB, ${db}, Test, $t: "
        linkcmd="--link ${container}:dbhost"
        cmd="${basecmd} ${linkcmd} ${kvimage}  "
        ${cmd} -db ${db} -host dbhost -port ${port} -user=${user} \
            -num-connections ${conn} -num-operations ${ops} -test ${t} \
            -label ${label}
        sleep 1
    done
    fi
}

function full_test {

    # Do a full suite of tests. Total ops will be divided by the number of
    # connections so that the total number of operations is the same across
    # all tests

    if [ -z "$1" ]
    then
        echo "Need to provide a database to test"
    else

    db=$1
    totalops=$2
    port=$3
    user=$4
    #conn=$5
    baselabel=$5
    container=$6
    runs=$7

    for run in `seq 1 ${runs}`; do

        for cn in ${conns}; do
            ops=$(( $totalops / $cn ))
            label="${baselabel}-${cn}-${run}"
            simple_test_run ${db} ${ops} ${port} ${user} ${cn} ${label} ${container}
        done
    done
    fi

}

