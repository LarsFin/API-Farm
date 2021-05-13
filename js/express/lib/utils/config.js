const DataLoader = require('./data-loader');
const InvalidEnvironmentError = require('../errors/invalid-environment');

const buildPath = env => `config.${env.toLowerCase()}.json`;

exports.fromEnvironment = env => new Promise((resolve, reject) => {
    if (['dev', 'prod'].includes(env.toLowerCase()))
        DataLoader.load(buildPath(env)).then(resolve).catch(reject);
    else
        reject(new InvalidEnvironmentError(env));
});
