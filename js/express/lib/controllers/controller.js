function Controller (videoGamesService) {
    this._videoGamesService = videoGamesService;
}

const invalidIdMessage = id => `The provided id '${id}' is invalid.`;

Controller.prototype.get = function (req, res) {
    const rawId = req.params.id;
    const id = parseInt(rawId);

    if (isNaN(id)) {
        res.badRequest(invalidIdMessage(rawId));
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

Controller.prototype.put = function (req, res) {
    const rawId = req.params.id;
    const id = parseInt(rawId);

    if (isNaN(id)) {
        res.badRequest(invalidIdMessage(rawId));
        return;
    }

    const query = this._videoGamesService.update(id, req.body);

    switch (query.code) {
    case 400:
        res.badRequest(query.result);
        break;

    case 404:
        res.notFound(query.result);
        break;

    default:
        res.ok(query.result);
    }
};

Controller.prototype.delete = function (req, res) {
    const rawId = req.params.id;
    const id = parseInt(rawId);

    if (isNaN(id)) {
        res.badRequest(invalidIdMessage(rawId));
        return;
    }

    const query = this._videoGamesService.delete(id);

    if (query.code === 404) {
        res.notFound(query.result);
        return;
    }

    res.okText(query.result);
};

module.exports = Controller;
