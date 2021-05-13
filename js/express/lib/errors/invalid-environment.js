class InvalidEnvironmentError extends Error {
    constructor (env) {
        super(`The environment '${env}' is not recognised.`);
        this.name = this.constructor.name;
        Error.captureStackTrace(this, this.constructor);
    }
}

module.exports = InvalidEnvironmentError;
