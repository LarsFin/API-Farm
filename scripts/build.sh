#!/bin/bash
# builds docker image for specified environment argument
# must be passed the lang framework as argument in the form 'lang/framework'
# e.g; ruby/sinatra

# language and framework; used for pathing to Dockerfile
LANG_FRAME=${1%/}

# check lang/framework was passed
if [ -z $LANG_FRAME ]
    then
        echo "Requires language and framework as first argument e.g; ruby/sinatra"
        exit 1
fi

# check if directory for language and framework exists
if ! dir $LANG_FRAME >/dev/null 2>&1
    then
        echo "A directory doesn't exist for '$LANG_FRAME'"
        exit 1
fi

# environments; dev, test, prod
ENV=$2

# if env is null; assume dev
if [ -z $ENV ]
    then
        echo "No environment specified as argument. Using dev." 
        ENV=dev
fi

IMG_REF="$LANG_FRAME:$ENV"

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

# check if language/framework has compilation stage
if [ -f "$LANG_FRAME/.compile" ]
    then
        # set target stage for docker build
        if [ $ENV -eq "prod" ]
            then
                $TARGET_FLAG="--target prod-env"
            else
                $TARGET_FLAG="--target dev-env"
        fi
    else
        $TARGET_FLAG=""
fi

# begin docker build
case $ENV in

    dev)
        echo "Building docker image: $LANG_FRAME:dev"
        docker build $TARGET_FLAG -t $IMG_REF $LANG_FRAME
        ;;

    prod)
        echo "Building docker image: $LANG_FRAME:prod"
        docker build $TARGET_FLAG -t $IMG_REF --build-arg env=prod $LANG_FRAME
        ;;

    *)
        echo "Unknown environment given; '${ENV}'."
        exit 1
esac