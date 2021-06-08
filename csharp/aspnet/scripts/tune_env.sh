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
        echo "Nothing to tune."
        ;;

    prod)
        echo "Prod environment selected"
        dotnet publish ApiFarm -c Release -r linux-x64
        ;;

    *)
        echo "Invalid environment selected!"
        exit 1
esac