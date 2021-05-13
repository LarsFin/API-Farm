const DataLoader = require('./data-loader');
const InvalidEnvironmentError = require('../errors/invalid-environment');

const buildPath = env => `config.${env.toLowerCase()}.json`;

exports.fromEnvironment = env => new Promise((res, rej) => {
    if (['dev', 'prod'].includes(env.toLowerCase()))
        DataLoader.load(buildPath(env)).then(res).catch(rej);
    else
        rej(new InvalidEnvironmentError(env));
});