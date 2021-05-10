function Controller (videoGamesService) {
    this._videoGamesService = videoGamesService;
}

Controller.prototype.getAll = function (res) {
    const query = this._videoGamesService.getAll();

    res.ok(query.result);
};

module.exports = Controller;
