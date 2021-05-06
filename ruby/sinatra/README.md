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

## Docker 🐳

This API is supported with Docker. You can check out how to run it by following the instructions in the root README [here](https://github.com/LarsFin/API-Farm).

## Resource Documentation

- **rspec:** https://rspec.info/documentation/
- **rubocop:** https://docs.rubocop.org/rubocop/1.11/index.html
- **ruby:** https://www.ruby-lang.org/en/documentation/
- **sinatra:** http://sinatrarb.com/documentation.html