#!/bin/bash

dbs="memcache mysql postgres"
tests="prepare select-by-pk update-by-pk"
ops=50000
for d in $dbs ; do
    for t in $tests ; do
        echo -n "DB, $d, Test, $t: "
        bin/kvbench -db $d -num-connections 1 -num-operations $ops -test $t
    done
done
