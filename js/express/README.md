# JS > Express

## Dependencies

NodeJS uses its own package manager 'npm' to control the dependencies of a project. Packages can be published to an npm respository (https://www.npmjs.com/) for others to install for their own projects. Installed packages can be located within the `node_modules` directory which should be ignored from source control.

The packages for a project are specifed within the `package.json` file at the root of the project. Installing a package on the command line will also append the package to your `package.json` dependencies. Dependencies in the `package.json` can also specified for the environment which requires them.

To install all dependencies for development & testing, run the command below;

```shell
npm install
```

When wanting to install dependencies only for running the the api (e.g; production), use the command below;

```shell
npm install --only=prod
```

## Running

After installing the necessary dependencies, you can choose to run the application via npm or node;

*via npm*
```
npm start
```

*via node*
```
node app.js
```

You can ensure the application is running successfully by sending a GET request to the /ping endpoint of the running api.

```console
C:\> curl http://localhost:8080/ping
pong
```

## Configuration

A configuration file can be found with the naming convention `config.<environment>.json`. This file simply configures which port and host binding to set for the api.

```json
{
    "hostname": "0.0.0.0",
    "port": 8080
}
```

## Testing

For the `js/express` api, we've decided to use Jest, a testing framework developed by Facebook. Once the jest node package has been installed, you can use the command below to run the unit tests;

```shell
npm test
```

## Linting

The `eslint` package is used to check through the code base for syntax and code smells. Use the command below to run eslint;

```shell
npm run lint
```

Note, this will only give a verbose output when issues have been found. The configuration for eslint's ruling has been set to default and tweaked to standards we feel comfortable with (additional/overwriting rules are defined within `.eslintrc` files).

## Docker 🐳

This API is supported with Docker. You can check out how to run it by following the instructions in the root README [here](https://github.com/LarsFin/API-Farm#docker-).

## Resource Documentation

- eslint: https://eslint.org/docs/user-guide/getting-started
- express: https://expressjs.com/en/starter/installing.html
- jest: https://jestjs.io/docs/getting-started
- nodejs: https://nodejs.org/en/docs/
