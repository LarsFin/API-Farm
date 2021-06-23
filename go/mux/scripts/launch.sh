#!/bin/bash

# commands; run, test, lint
CMD=$1

# if cmd is null; assume run
if [ -z $CMD ]
    then CMD=run
fi

# run process correlating to command
case $CMD in

    run)
        echo "Running app..."
        ./apifarm
        ;;

    test)
        echo "Running test suite..."
        go test apifarm/src
        ;;

    lint)
        echo "Linting codebase..."
        golangci-lint run
        ;;

    *)
        echo "Unknown command given; ${CMD}"
        exit 1
esac