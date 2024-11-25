#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/gomall.hello"
exec "$CURDIR/bin/gomall.hello"