## KV-Bench

This is a simple sysbench-inspired benchmarking tool for RDBMS and key-value systems written in Go.
It has been written to test systems for key-value type workloads (for the reasons
why I would want to do such a thing, see [the blog post](https://notsobigdata.blogspot.com) this tool was
built for)

### Tests
The tool currently supports four test types:

* insert-by-pk
* update-by-pk
* select-by-pk
* delete-by-pk

Each operation runs individually and the keys are randomly chosen without replacement for the size of the test.

### System support

Currently, the tool supports:

* MySQL
* Postgres
* Vanilla Memcached
* Memcached-innodb (from)

Eventually I would like to expand this to test other NoSQLish systems and JSONB and HStore for Postgres.

### Usage

The tool has been dockerized with the main kvbench binary as the entry point to make it easy
to deploy along with the desired database to test.

Parameters supported:

      -db string
            Database type (default "postgres")
      -host string
            Target host (default "192.168.42.223")
      -label string
            Label to add to test result output (default "test")
      -num-connections int
            Number of connections to DB
      -num-operations int
            Number of queries/updates/etc. per conn
      -password string
            Password if required (default "pw1pw1pw1")
      -port int
            Port of db
      -test string
            Test type: prepare, select-by-pk, etc.
      -user string
            Username if required (default "u1")


E.g. to use with docker, you can run with a command like:

    docker run --rm -i -t --link postgrestest:pgtst atomic77/kvbench \
        -db postgres -user postgres -host pgtst -port 5432 -num-connections 4 \
        -num-operations 1000 -test select-by-pk

The above runs the kvbench tool against the pgsql database running in the `postgrestest` container,
with 4 client connections running random selects with 1000 selects per connection. See the scripts
directory for more examples of how you can run the tool.
