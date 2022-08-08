#!/bin/bash
curpath=`dirname $0`
cd $curpath/

go get  && go build -o relay-remind
