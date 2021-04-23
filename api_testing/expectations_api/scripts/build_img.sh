#!/bin/bash
# builds expectations api image, overwriting an already existing one

BUILD_V=$1

# If no version of expectations api image is provided. Assume creation of latest
if [ -z $BUILD_V ]
    then
        echo "No build version provided. Using 'latest'."
        BUILD_V=latest
fi

IMG_REF="expectations_api:$BUILD_V"

# If expectations api image with version exists already; remove
if docker image inspect $IMG_REF >/dev/null 2>&1
    then
        # Check for running containers with image to stop
        CONTAINERS=$(docker ps --filter ancestor=$IMG_REF --filter status=running --format {{.ID}})

        # Stop all discovered containers
        for CONTAINER in $CONTAINERS; do
            docker stop $CONTAINER >/dev/null 2>&1
        done

        # Check for any containers with image to remove
        CONTAINERS=$(docker ps --filter ancestor=$IMG_REF --format {{.ID}})

        # Remove all containers with image
        for CONTAINER in $CONTAINERS; do
            docker rm $CONTAINER >/dev/null 2>&1
        done

        echo "Image '$IMG_REF' already exists. Overwriting..."
        docker rmi $IMG_REF
fi

docker build -t $IMG_REF .