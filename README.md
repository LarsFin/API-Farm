# API Farm

A repository of APIs written in different programming languages and frameworks. Each framework's development has been unit and API tested. Also, documented with instructions for how to start up locally (addtional docker support included too).

## API Use Case

The api developed in each environment should support creation, reading, updating and deletion of a resource. The management of data can vary from in memory, file storage or a database.

Each api should be unit tested with a language supported framework where possible. Also, generic API tests should be run against each working instance of the API. Results from these tests should be collected for comparison.

To allow fair metric comparison from api tests, the api model resource will be standardised across each of the languages and frameworks. The api model resource will be a `video game`. This resource has been structured below in JSON;

```json
{
    "name": "Banjo Kazooie",
    "developers": [
        "Rare"
    ],
    "publishers": [
        "Nintendo"
    ],
    "directors": [
        "Gregg Mayles",
        "George Andreas"
    ],
    "producers": [
        "Tim Stamper",
        "Chris Stamper"
    ],
    "designers": [
        "Gregg Mayles"
    ],
    "programmers": [
        "Chris Sutherland"
    ],
    "artists": [
        "Steve Mayles",
        "John Nash",
        "Kevin Bayliss",
        "Tim Stamper"
    ],
    "composers": [
        "Grant Kirkhope"
    ],
    "platforms": [
        "Nintendo 64"
    ],
    "date_released": "29/06/1998"
}
```

## Languages & Frameworks

Below, the various languages and frameworks are listed which are used within this repository;

| Language | API Framework | Test Framework | Storage Support | Docker Support | Link |
| -------- | ------------- | -------------- | --------------- | :------------: | ---- |
| Ruby     | Sinatra       | Rspec          | In Memory       |       ✔️      | [ruby/sinatra](ruby/sinatra)  |
| JavaScript | ExpressJS   | Jest           | In Memory       |       ✔️      | [js/express](js/express) |

## Project Structure

This repository is designed to clearly silo each of the api samples.

```
api_farm
|___lang-1
|   |___framework-1
|   |___framework-2
|___lang-2
    |___framework-1
```

Each framework will have its own standalone api example within a separate directory.

## Docker 🐳

Most apis can be run and tested using Docker. Additionally, lint checks can also be run against the api's code base too. Scripts exist within `./scripts` to aid in managing api containers and images. Alternatively, you can build the lang/framework Dockerfile yourself.

The guides below use the shell scripts.

### Building an Image

An image for a lang/framework can be build under two environmental contexts;
- **dev:** a heavier image including tools necessary for linting and src code testing.
- **prod:** a light weight image, for compiled languages this may include only the executable.

The `./scripts/build.sh` script takes two arguments. The first instructs which lang/framework to build. The second; under which environment to build the image. For example; wanting to build the `ruby/sinatra` api image for a dev environment you could run the command below.

```shell
./scripts/build.sh ruby/sinatra dev
```

Running the command above will build the image `ruby/sinatra:dev`. The tag would be suffixed with `prod` if it were built for a production environment.

Not providing an environment will result in an image being built for the dev environment by default. It should be noted that this script will overwrite existing built images for the target lang/framework with the environment. As part of this, it will stop and remove any containers using older versions of the image.

### Running an Image

Docker images for an api can be run using the `./scripts/run.sh` script. Like the build script, a lang/framework must be specified as the first argument. The second argument is the way in which you wish to run the image, these are listed below;

- **lint:** uses language tools to comb the api code base for code smells and unwanted syntax styles.
- **test:** runs the unit tests for the api src code.
- **run_dev:** runs the api using the `dev` environment image.
- **run_prod:** runs the api using the `prod` environment image.

The `lint` and `test` commands use the `dev` image as the necessary tools should have been stripped from the production image. If the required image for the command cannot be found; the script process return a non zero exit code.

If no command is passed as an argument to the `run.sh` script, it will assumes the `run_dev` command.

Below, is an example of how to run the `run.sh` script.

```shell
./scripts/run.sh ruby/sinatra test
```

Containers started for linting and testing will be temporary and removed after the process has finished. When running the api, the containers will continue to run in a detached state. To stop these containers, read the following section.

### Stopping Containers

The script `./scripts/clean.sh` stops containers for the passed lang/framework. Below, is an example of how to run the script;

```shell
./scripts/clean.sh ruby/sinatra
```

There is no harm in using docker directly to stop containers. This script has a check to ensure a container exists for the lang/framework before stopping it (this functionality is beneficial in the post stage of the pipeline).

```shell
docker stop ruby_sinatra
```

## CI ⚙️

Jenkins is used to ensure pull requests into lang/framework branches and the master branch are healthy. A multi branch pipeline has been configured within Jenkins to *build, lint, unit test, health check and run api tests* against pull requests. Below, is a screen shot of a successful pull request within Jenkins.

![Jenkins Build](./img/jenkins-build.JPG)

Above, is a pull request which was made into master (`ruby/sinatra` -> `master`). For pull requests into a lang/framework branch (i.e; `ruby/sinatra_post` -> `ruby/sinatra`), all stages above are run except for api tests.

A GitHub user has been created '[RangdaBot](https://github.com/RangdaBot)' to communicate the status of builds within Jenkins for a pull request. Rules have been configured on branches to ensure builds were successful in Jenkins before merging.

Below, is an image of the bot user account in action.

![GitHub PR](./img/github-pr.JPG)

The AWS EC2 instance running Jenkins uses Docker to run its stages. This prevents disk space on the EC2 instance getting cramp from language dependencies.

### Stage Details

Below, is an illustration of the pipeline stages with a little more information on what exactly each stage includes.

![API Farm Pipeline](./img/api-farm-pipeline.jpg)

As mentioned in the image above, check out the API Testing [README](api_testing/README.md) for more info on API Testing.
