function Controller (videoGamesService) {
    this._videoGamesService = videoGamesService;
}

Controller.prototype.getAll = function (res) {
    const query = this._videoGamesService.getAll();

    res.ok(query.result);
};

Controller.prototype.add = function (req, res) {
    const query = this._videoGamesService.add(req.body);

    if (query.code === 400) {
        res.badRequest(query.result);
        return;
    }

    res.created(query.result);
};

module.exports = Controller;
