# Ruby > Sinatra

Sinatra is a light weight application development framework for ruby.

## Dependencies

Ruby uses gems to control the dependencies required for an application. Installed gems can be found within your installed ruby binary directory. 
To ensure dependencies can easily be installed on other machines we use a Gemfile to detail which dependencies we require. 
Bundler is used to install the gems using the Gemfile.

```shell
bundle install
```

## Running

Install the `sinatra` gem. Then simply run `app.rb`!

```shell
ruby app.rb
```

You can check that the application is running successfully by sending a `GET` request to the `/ping` endpoint of the running api.

```console
C:\> curl http://localhost:8080/ping
pong
```

### Configuration

A configuration file can be found with the naming convention `config.<environment>.json`. This file simply configures which port and bind to set for the api.

```json
{
    "bind": "0.0.0.0",
    "port": 8080
}
```

In later stages, configuration will allow for various storage methods (e.g; database, file, etc.).

## Testing

Install the `rspec`, `simplecov` and `simplecov-console` gem. Then run the test suite using the command below.

```shell
rspec
```

## Linting

Ensure you have installed the `rubocop` gem. Then simply run the following line in your command line.

```shell
rubocop
```

Configuration for what rubocop scrutinises is within the `.rubocop.yml` file.

*Note, you can also lint the source code using Docker as detailed below.*

## Docker 🐳

It is possible to run the api server, source code testing and linting of code base using docker. In a bash terminal you can use the scripts found in `./scripts` to build new images and run existing images!

### Building an Image

To build an image, you can run the `./scripts/build_img.sh` script in a bash terminal (make sure to run it from the `ruby/sinatra` directory). It takes an argument of the envrionment you wish to prepare `dev` or `prod` (production lacks any of the dev & test gem dependencies). If you wish to run linting or unit tests; use `dev`, otherwise use `prod` as the build times and image size is smaller. Not providing an argument will default to `dev`.

```
./scripts/build_img.sh <env>
```

If the image already exists, it will remove it and build a fresh image.

### Running an Image

To run an existing image, you can run the `./scripts/run_img.sh` script in a bash terminal (make sure to run it from the `ruby/sinatra` directory). It takes an argument for the kind of process you wish to run. The available options are; `run_prod`, `run_dev`, `test` or `lint`. It will default to `run_dev` if no argument is provided. The latter three commands require an existing image `ruby/sinatra:dev` where as the first command needs the image `ruby/sinatra:prod`. If these images do not exist; it will exit with a failure code.

```
./scripts/run_img.sh <cmd>
```

When run, the ports on 8080 will be exposed for access. Additionally, all processes run via the script instruct Docker to delete the running container after the process has exited.

## Resource Documentation

- **rspec:** https://rspec.info/documentation/
- **rubocop:** https://docs.rubocop.org/rubocop/1.11/index.html
- **ruby:** https://www.ruby-lang.org/en/documentation/
- **sinatra:** http://sinatrarb.com/documentation.html