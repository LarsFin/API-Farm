const express = require('express');

// 2** Responses

express.response.ok = function (body) {
    this.status(200);
    this.header('Content-Type', 'application/json');
    this.json(body);
};

express.response.okText = function (message) {
    this.status(200);
    this.header('Content-Type', 'text/plain');
    this.send(message);
}

express.response.created = function (body) {
    this.status(201);
    this.header('Content-Type', 'application/json');
    this.json(body);
};

// 4** Responses

express.response.badRequest = function (reason) {
    this.status(400);
    this.header('Content-Type', 'text/plain');
    this.send(reason);
};

express.response.notFound = function (reason) {
    this.status(404);
    this.header('Content-Type', 'text/plain');
    this.send(reason);
};

// 5** Responses

express.response.internalServerError = function () {
    this.status(500);
    this.header('Content-Type', 'text/plain');
    this.send('Something went wrong.');
};
