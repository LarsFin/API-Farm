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
    echo "Setting ASPNETCORE_ENVIRONMENT."
    export ASPNETCORE_ENVIRONMENT=$API_ENV
    
    echo "Running app..."
    dotnet run --project ApiFarm
    ;;

  test)
    echo "Running test suite..."
    dotnet test /p:Exclude=\"[*]ApiFarm.Scaffolding.*,[*]ApiFarm.Models.*,[*]ApiFarm.Utils.*\" /p:CollectCoverage=true /p:Threshold=100
    ;;

  lint)
    echo "Linting codebase..."
    dotnet build /warnaserror
    ;;
  
  *) 
    echo "Unknown command given; ${CMD}"
    exit 1
esac
