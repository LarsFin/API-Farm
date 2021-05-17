# Expectations API

A simple express server which serves static json files for API testing. This API is a dependency of running api tests, without it; tests which could pass will always fail.

## Running

You can choose to run this api on your local machine using the command below;

```shell
node app.js
```

However, it's best advised to run the api through docker. Like with the language/framework apis, there are helper scripts to boot this up. Simply run the commands below;

```shell
./scripts/build_img.sh
./scripts/run_img.sh
```

It is only necessary to build the image once (unless changes to the expectations api have been made). If the image exists in your local docker repo; you can simply run the `run_img.sh` script to start it up.

## Expected Data

The directory `expected_data` includes all the static json files containing the expected responses for various requests against the lang/framework apis.
These data files have been organised into directories of the request endpoint. The files have been named according to the response type (these are `ok` and `created`).
There are no expected json responses for Bad Requests or Not Found responses due to their body being simple enough to include within the Postman tests themselves.
