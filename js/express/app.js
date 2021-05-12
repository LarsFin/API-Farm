const express = require('express');
const app = express();
const port = 8080;
const hostname = '0.0.0.0';

require('./lib/extensions/response');

app.use(express.json({
    verify: (_, res, buf, __) => {
        try {
            JSON.parse(buf);
        } catch {
            res.badRequest('Invalid JSON in body.');
        }
    }
}));

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

app.get('/video_games/:id', (req, res) => {
    controller.get(req, res);
});

app.post('/video_games', (req, res) => {
    controller.add(req, res);
});

app.put('/video_games/:id', (req, res) => {
    controller.put(req, res);
});

app.listen(port, hostname, () => {
    console.log(`Example app listening at http://localhost:${port}`);
});
