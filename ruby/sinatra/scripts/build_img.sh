#!/bin/bash
# builds docker image for specified environment argument
# ensure this is run from lang/framework folder!

# environments; dev, test, prod
ENV=$1

# if env is null; assume dev
if [ -z $ENV ]
    then
        echo "No environment specified as argument. Using dev." 
        ENV=dev
fi

IMG_REF="ruby/sinatra:$ENV"

# check if image for environment already exists
if docker image inspect $IMG_REF >/dev/null 2>&1
    then
        # check for running containers with image to stop
        CONTAINERS=$(docker ps --filter ancestor=$IMG_REF --filter status=running --format {{.ID}})

        # stop all discovered containers
        for CONTAINER in $CONTAINERS; do
            echo "Container with id '$CONTAINER' running with old image. Stopping container."
            docker stop $CONTAINER >/dev/null 2>&1
        done

        # check for any containers with image to remove
        CONTAINERS=$(docker ps --filter ancestor=$IMG_REF --format {{.ID}})

        # remove all containers with image
        for CONTAINER in $CONTAINERS; do
            docker rm $CONTAINER >/dev/null 2>&1
        done

        echo "Image $IMG_REF already exists. Removing..."
        docker rmi $IMG_REF
fi

# begin docker build
case $ENV in

    dev)
        echo "Building docker image: ruby/sinatra:dev"
        docker build -t ruby/sinatra:dev .
        ;;

    prod)
        echo "Building docker image: ruby/sinatra:prod"
        docker build -t ruby/sinatra:prod --build-arg env=prod .
        ;;

    *)
        echo "Unknown environment given; '${ENV}'."
        exit 1
esac