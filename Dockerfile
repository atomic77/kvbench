# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Uncomment to build based off local copy
# ADD . /go/src/github.com/atomic77/kvbench

# Seems like go get will grab/install the code and dependencies
RUN go get github.com/atomic77/kvbench

# Run the kvbench cmd - cmdline params can be provided after
# e.g. 
# docker run --rm -i -t --link postgrestest:pgtst atomic77/kvbench
#      -db postgres -user postgres -host pgtst -port 5432 -num-connections 4
#       -num-operations 1000 -test select-by-pk
# Where postgrestest is a container running pgsql

ENTRYPOINT ["kvbench"]

