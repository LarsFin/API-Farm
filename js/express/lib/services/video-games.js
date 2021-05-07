const Query = require('../utils/query');

function VideoGamesService (storage) {
    this._storage = storage;
}

VideoGamesService.prototype.getAll = function () {
    const videoGames = this._storage.getAllVideoGames();

    return Query.success(videoGames);
};

module.exports = VideoGamesService;
