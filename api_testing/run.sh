# !/bin/bash
# Run from 'api_testing' directory
# Takes on steps to run api tests. Script's steps are outlined below;
# > Check passed argument; should be service name to run tests against.
# > Check api_testing/newman image exists. If it does; delete it. Always build.
# > Ensure api & expectations api is running too.
# > Run postman/newman image with necessary arguments
# > Export api test results to file
# > Copy file to local machine
# > Stop and remove api testing container

# Name of container running api
API_CONTAINER=$1

# Check api name argument has been provided
if [ -z $API_CONTAINER ]
    then
        echo "An api name must be provided as an argument."
        exit 1
fi

# Delete api_testing image if it already exists
if docker image inspect api_testing/newman >/dev/null 2>&1
    then
        docker rmi api_testing/newman
fi

# Build api_testing image
docker build -t api_testing/newman .

# Query docker for containers running on api_farm_dev network
DOCKER_QUERY_RESULT=$(docker ps -f status=running --format "{{.Names}}") # TODO: add network query!

# Ensure api container is running on network
if [[ $DOCKER_QUERY_RESULT != *$API_CONTAINER* ]]
    then
        echo "No running container '$API_CONTAINER' could be found with network api_farm_dev."
        exit 1
fi

# Ensure expectations api is running on network
if [[ $DOCKER_QUERY_RESULT != *"expectations_api"* ]]
    then
        echo "No running container for expectations api could be found with network api_farm_dev"
        exit 1
fi

# Define results output file name
RESULTS_FILE=${API_CONTAINER}_api_test_results

# Run api tests
docker run --network=api_testing --name= -t api_testing/newman run API_farm.postman_collection.json \
    --folder Tests -e ApiTesting.api_farm.json --env-var host=$API_CONTAINER --reporters=cli,json --reporter-json-export ${RESULTS_FILE}.json > ${RESULTS_FILE}.txt
# TODO: Change network!

# 

# !!! All services for api testing should be on api_farm_dev network