#!/bin/bash

# tunes necessary dependencies against the build environment

# environments dev or prod
ENV=$1

# if env is null; assume dev
if [ -z $ENV ]
    then ENV=dev
fi

case $ENV in

    dev)
        echo "Dev environment selected"
        go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1 # linting
        go get -u github.com/mcubik/goverreport                            # testing
        ;;

    prod)
        echo "Prod environment selected"
        echo "Nothing to tune."
        ;;

    *)
        echo "Invalid environment selected!"
        exit 1
esac
