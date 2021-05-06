#!/bin/bash
# runs docker image for specified lang/framework. Must be passed as 
# script argument

# language and framework; used for pathing to Dockerfile
LANG_FRAME=${1%/}

# check lang/framework was passed
if [ -z $LANG_FRAME ]
    then
        echo "Requires language and framework as first argument e.g; ruby/sinatra"
        exit 1
fi

# build service name
SERVICE_NAME=$(echo $LANG_FRAME | tr / _)

# commands; run_prod, run_dev, test, lint
CMD=$2

# if cmd is null; assume run_dev
if [ -z $CMD ]
    then
        echo "No command specified as argument. Using 'run_dev'."
        CMD=run_dev
fi

# determine image reference
case $CMD in

    run_dev)
        IMG_REF="$LANG_FRAME:dev"
        ;;
    
    run_prod)
        IMG_REF="$LANG_FRAME:prod"
        ;;

    test)
        IMG_REF="$LANG_FRAME:dev"
        ;;

    lint)
        IMG_REF="$LANG_FRAME:dev"
        ;;

    *)
        echo "Unknown command given; '${CMD}'."
        exit 1

esac

# check if image exists
if ! docker image inspect $IMG_REF >/dev/null 2>&1
    then
        echo "The required image '$IMG_REF' does not exist."
        exit 1
fi

# run image inside container with specified command
case $CMD in

    run_dev)
        echo "Checking if api_farm_dev network is available."
        if ! docker network inspect api_farm_dev >/dev/null 2>&1
            then
                echo "The network 'api_farm_dev' does not exist. Creating now..."
                docker network create api_farm_dev
        fi

        echo "Running container with run command for image '$IMG_REF'"
        docker run --rm -dt -p 8080:8080 --network=api_farm_dev --name $SERVICE_NAME -e API_ENV=DEV $IMG_REF run
        ;;
        
    run_prod)
        echo "Checking if api_farm_prod network is available."
        if ! docker network inspect api_farm_prod >/dev/null 2>&1
            then
                echo "The network 'api_farm_prod' does not exist. Creating now..."
                docker network create api_farm_prod
        fi

        echo "Running container with run command for image '$IMG_REF'"
        docker run --rm -dt -p 8080:8080 --network=api_farm_prod --name $SERVICE_NAME -e API_ENV=PROD $IMG_REF run
        ;;

    test)
        echo "Running container with test command for image '$IMG_REF'"
        docker run --rm $IMG_REF test
        ;;

    lint)
        echo "Running container with lint command for image '$IMG_REF'"
        docker run --rm $IMG_REF lint
        ;;
esac