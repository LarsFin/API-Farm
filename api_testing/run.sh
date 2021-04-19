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

echo "🚀 Starting API Testing Process!";

# Name of container running api
API_CONTAINER=$1

# Check api name argument has been provided
if [ -z $API_CONTAINER ]
    then
        echo "An api name must be provided as an argument."
        exit 1
fi

# Setup logging
DATE_TIME=$(date '+%d%m%Y%H%M')
LOGGING_FILE=./logs/${DATE_TIME}_api_test_run.log

# Create logs directory if it does not exist
if ! dir logs >/dev/null 2>&1
    then
        echo "No ./logs directory. Creating one..."
        mkdir logs
        echo "./logs created."
fi

# Create log file
touch $LOGGING_FILE

# Delete api_testing image if it already exists
if docker image inspect api_testing/newman >>$LOGGING_FILE
    then
        echo "Removing old api testing image..."
        docker rmi api_testing/newman >>$LOGGING_FILE 2>&1
        echo "Old api testing image removed."
fi

# Build api_testing image
echo "Building api testing image..."
docker build -t api_testing/newman . >>$LOGGING_FILE 2>&1
echo "API testing image successfully created."

# Query docker for containers running on api_farm_dev network
echo "Checking docker for required containers on api_farm_dev network; '${API_CONTAINER}' and 'api_expectations'..."
DOCKER_QUERY_RESULT=$(docker ps -f status=running -f network=api_farm_dev --format "{{.Names}}")

# Ensure api container is running on network
if [[ $DOCKER_QUERY_RESULT != *$API_CONTAINER* ]]
    then
        echo "No running container '$API_CONTAINER' could be found with network api_farm_dev."
        exit 1
fi

# Ensure expectations api is running on network
if [[ $DOCKER_QUERY_RESULT != *"expectations_api"* ]]
    then
        echo "No running container for expectations api could be found with network api_farm_dev."
        exit 1
fi

echo "Services located."

# Define results output file name
RESULTS_FILE=${API_CONTAINER}_api_test_results
API_TESTS_CONTAINER=api_testing_$DATE_TIME

# Create results directory if it does not exist
if ! dir results >>$LOGGING_FILE 2>&1
    then
        echo "No ./results directory. Creating one..."
        mkdir results
        echo "./results created."
fi

# Run api tests
echo "Running api testing image..."
docker run --network=api_farm_dev --name=$API_TESTS_CONTAINER -t api_testing/newman run API_farm.postman_collection.json \
    --folder Tests -e ApiTesting.api_farm.json --env-var host=$API_CONTAINER --reporters=cli,json --reporter-json-export ${RESULTS_FILE}.json >>$LOGGING_FILE 2>&1

# Query docker for api tests container
DOCKER_QUERY_RESULT=$(docker ps -f status=exited --format "{{.Names}}")

# If api tests container does not exist. Fail under assumption there was an issue running the container.
if [[ $DOCKER_QUERY_RESULT != *$API_TESTS_CONTAINER* ]]
    then
        echo "There was an issue running the api tests container;"
        cat ${RESULTS_FILE}.txt
        exit 1
fi

# Copy output file from api tests
docker cp $API_TESTS_CONTAINER:/etc/newman/${RESULTS_FILE}.json ./results/${RESULTS_FILE}.json

echo "API testing image successfully run."

# Delete api tests container
echo "Removing api testing container..."
docker rm $API_TESTS_CONTAINER >>$LOGGING_FILE 2>&1
echo "API testing container removed."

echo "API Testing Process Complete ✔️"