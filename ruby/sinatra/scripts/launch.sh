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
        ruby app.rb
        ;;

    test)
        echo "Running test suite..."
        rspec
        ;;

    lint)
        echo "Linting codebase..."
        rubocop
        ;;

    *)
        echo "Unknown command given; ${CMD}"
        exit 1
esac