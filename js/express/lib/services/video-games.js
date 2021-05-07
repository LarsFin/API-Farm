function VideoGamesService (storage) {
    this._storage = storage;
}

VideoGamesService.prototype.getAll = function () {
    const videoGames = this._storage.getAllVideoGames();

    return {
        code: 0,
        result: videoGames
    };
};

module.exports = VideoGamesService;
