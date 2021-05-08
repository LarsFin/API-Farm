const express = require('express');
const app = express();
const port = 8080;

const Controller = require('./lib/controllers/controller');
const VideoGamesService = require('./lib/services/video-games');
const InMemory = require('./lib/storage/in-memory');

const storage = new InMemory();
const videoGamesService = new VideoGamesService(storage);
const controller = new Controller(videoGamesService);

app.get('/ping', (_, res) => {
    res.send('pong');
});

app.get('/video_games', (_, res) => {
    controller.getAll(res);
});

app.listen(port, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});
