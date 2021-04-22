#!/bin/bash
# runs expectations api image

IMG_V=$1

# If no version of image is provided as argument; assume latest
if [ -z $IMG_V ]
    then
        echo "No version specified, using 'latest'"
        IMG_V=latest
fi

IMG_REF="expectations_api:$IMG_V"

# Check if expecatations api image exists. Exit if not built.
if ! docker image inspect $IMG_REF >/dev/null 2>&1
    then
        echo "The image '$IMG_REF' does not exist. Please build the image first."
        exit 1
    else
        # Check for running containers with image to stop
        CONTAINERS=$(docker ps --filter ancestor=$IMG_REF --filter status=running --format {{.ID}})

        # Stop all discovered containers
        for CONTAINER in $CONTAINERS; do
            echo "Stopping already runnning container with id '$CONTAINER'."
            docker stop $CONTAINER
        done
fi

# Check network dev api testing network exists. Create if it doesn't
if ! docker network inspect api_farm_dev >/dev/null 2>&1
    then
        echo "No network 'api_farm_dev' exists. Creating now."
        docker network create api_farm_dev
fi

docker run --rm -p 3000:3000 --name expectations_api --network api_farm_dev $IMG_REF