const express = require('express');

// 2** Responses

express.response.ok = function (body) {
    this.status(200);
    this.header('Content-Type', 'application/json');
    this.json(body);
};
