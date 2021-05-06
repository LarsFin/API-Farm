#!/bin/bash
# cleans up running containers for specified lang/framework which is 
# provided as an argument to this script

# get lang/framework
LANG_FRAME=${1%/}

# get service name
SERVICE_NAME=$(echo $LANG_FRAME | tr / _)

# check if service name is populated
if [ -z $SERVICE_NAME ]
    then
        echo "Requires language and framework as first argument e.g; ruby/sinatra"
        exit 1
fi

# check if service container exists
if ! docker container inspect $SERVICE_NAME >/dev/null 2>&1
    then
        echo "No container exists for $LANG_FRAME."
        exit 0
fi

# check if service is running and stop
if [ $(docker container inspect $SERVICE_NAME --format {{.State.Status}}) == running ]
    then
        docker stop $SERVICE_NAME >/dev/null 2>&1
        echo "Stopped container for $LANG_FRAME"
fi

# check if container still exists in stopped state
if docker container inspect $SERVICE_NAME >/dev/null 2>&1
    then
        echo "Removing stopped container for $LANG_FRAME"
        docker rm $SERVICE_NAME
fi