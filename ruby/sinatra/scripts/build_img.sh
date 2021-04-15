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

# check if image for environment already exists
if [ "$ENV" == "prod" ]
    then
        if docker image inspect ruby/sinatra:prod >/dev/null 2>&1
            then
                echo "Image for ruby/sinatra prod already exists. Removing..."
                docker rmi ruby/sinatra:prod
        fi
    else
        if docker image inspect ruby/sinatra:dev >/dev/null 2>&1
            then
                echo "Image for ruby/sinatra dev already exists. Removing..."
                docker rmi ruby/sinatra:dev
        fi
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