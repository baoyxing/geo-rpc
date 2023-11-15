#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/geo-rpc"
exec "$CURDIR/bin/geo-rpc"