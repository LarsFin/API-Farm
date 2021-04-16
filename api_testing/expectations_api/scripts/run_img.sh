#!/bin/bash
# runs expectations api image

IMG_V=$1

if [ -z $IMG_V ]
    then
        echo "No version specified, using 'latest'"
        IMG_V=latest
fi

IMG_REF="expectations_api:$IMG_V"

if ! docker image inspect $IMG_REF >/dev/null 2>&1
    then
        echo "The image '$IMG_REF' does not exist. Please build the image first."
        exit 1
fi

docker run --rm -p 3000:3000 --name expectations_api $IMG_REF