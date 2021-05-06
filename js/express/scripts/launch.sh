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
    npm run start
    ;;

  test)
    echo "Running test suite..."
    npm run test
    ;;

  lint)
    echo "Linting codebase..."
    npm run lint
    ;;
  
  *) 
    echo "Unknown command given; ${CMD}"
    exit 1
esac