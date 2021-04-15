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