#!/bin/bash
# builds expectations api image, overwriting an already existing one

BUILD_V=$1

if [ -z $BUILD_V ]
    then
        echo "No build version provided. Using 'latest'."
        BUILD_V=latest
fi

IMG_REF="expectations_api:$BUILD_V"

if docker image inspect $IMG_REF >/dev/null 2>&1
    then
        echo "Image '$IMG_REF' already exists. Overwriting..."
        docker rmi $IMG_REF
fi

docker build -t $IMG_REF .