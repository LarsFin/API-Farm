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
        bundle install
        ;;

    prod)
        echo "Prod environment selected"
        rm .rubocop.yml
        rm .rspec
        rm -rf spec
        bundle install --without dev test
        ;;

    *)
        echo "Invalid environment selected!"
        exit 1
esac