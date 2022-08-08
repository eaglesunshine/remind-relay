#!/bin/bash

#propagate terminated signal to all sub process from this script.
#trap 'kill -s SIGTERM 0' SIGTERM SIGINT

#get web server service port, from docker run -e setting.
port="9701"

#start http server 
if [ -n "$port" ]; then
    /data/relay-remind/bin/relay-remind -cmd=HttpServer -arg=$port >> /data/relay/var/HttpServer.log 2>&1
fi

#wait
