function Controller (videoGamesService) {
    this._videoGamesService = videoGamesService;
}

Controller.prototype.getAll = function (res) {
    const query = this._videoGamesService.getAll();

    res.status(200);
    res.header('Content-Type', 'application/json');
    res.json(query.result);
};

module.exports = Controller;
