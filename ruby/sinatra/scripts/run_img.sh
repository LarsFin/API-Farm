#!/bin/bash

# commands; run_prod, run_dev, test, lint
CMD=$1

# if cmd is null; assume run_dev
if [ -z $CMD ]
    then
        echo "No command specified as argument. Using 'run_dev'."
        CMD=run_dev
fi

# check if image exists for specified command
if [ "$CMD" == "run_prod" ]
    then
        if ! docker image inspect ruby/sinatra:prod >/dev/null 2>&1
            then
                echo "The required image 'ruby/sinatra:prod' does not exist."
                exit 1
        fi
    else
        if ! docker image inspect ruby/sinatra:dev >/dev/null 2>&1
            then
                echo "The required image 'ruby/sinatra:dev' does not exist."
                exit 1
        fi
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

        echo "Running container with run command for image 'ruby/sinatra:dev'"
        export RUBY_SINATRA_ENV=DEV
        docker run --rm -p 8080:8080 --network=api_farm_dev --name ruby_sinatra ruby/sinatra:dev run
        ;;
        
    run_prod)
        echo "Checking if api_farm_prod network is available."
        if ! docker network inspect api_farm_prod >/dev/null 2>&1
            then
                echo "The network 'api_farm_prod' does not exist. Creating now..."
                docker network create api_farm_prod
        fi

        echo "Running container with run command for image 'ruby/sinatra:prod'"
        export RUBY_SINATRA_ENV=PROD
        docker run --rm -p 8080:8080 --network=api_farm_dev --name ruby_sinatra ruby/sinatra:prod run
        ;;

    test)
        echo "Running container with test command for image 'ruby/sinatra:dev'"
        docker run --rm ruby/sinatra:dev test
        ;;

    lint)
        echo "Running container with lint command for image 'ruby/sinatra:dev"
        docker run --rm ruby/sinatra:dev lint
        ;;
    
    *)
        echo "Unknown command given; '${CMD}'."
        exit 1
esac