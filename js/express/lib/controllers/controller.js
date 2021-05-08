function Controller (videoGamesService) {
    this.videoGamesService = videoGamesService;
}

Controller.prototype.getAll = function (res) {
    const videoGames = this.videoGamesService.getAll();

    res.status(200);
    res.header('Content-Type', 'application/json');
    res.json(videoGames);
};

module.exports = Controller;
