const DataLoader = require('./data-loader');
const InvalidEnvironmentError = require('../errors/invalid-environment');

const buildPath = env => `config.${env.toLowerCase()}.json`;

exports.fromEnvironment = env => {
    if (['dev', 'prod'].includes(env.toLowerCase()))
        return DataLoader.load(buildPath(env));

    throw new InvalidEnvironmentError(env);
};
