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
    python app.py
    ;;

  test)
    echo "Running test suite..."
    pytest
    ;;

  lint)
    echo "Linting codebase..."
    pylint flask
    ;;
  
  *) 
    echo "Unknown command given; ${CMD}"
    exit 1
esac