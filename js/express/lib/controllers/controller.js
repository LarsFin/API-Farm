function Controller (videoGamesService) {
    this._videoGamesService = videoGamesService;
}

Controller.prototype.get = function (req, res) {
    const rawId = req.params.id;
    const id = parseInt(rawId);

    if (isNaN(id)) {
        res.badRequest(`The provided id '${rawId}' is invalid.`);
        return;
    }

    const query = this._videoGamesService.get(id);

    if (query.code === 404) {
        res.notFound(query.result);
        return;
    }

    res.ok(query.result);
};

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
