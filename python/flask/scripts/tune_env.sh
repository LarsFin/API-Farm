#!/bin/bash
# tunes necessary dependencies against the build environment

# environments; dev, test, prod
ENV=$1

# if env is null; assume dev
if [ -z $ENV ]
    then ENV=dev
fi

# delete files unneeded to environments
case $ENV in

    dev)
        echo "Dev environment selected"
        echo "No files to remove."
        pip install -r requirements.txt
        ;;

    prod)
        echo "Prod environment selected"
        rm -rf spec
        pip install -r requirements.txt
        ;;

    *)
        echo "Invalid environment selected!"
        exit 1
esac