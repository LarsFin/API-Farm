function InMemory () {
    this._idCounter = 0;
    this._videoGames = [];
}

InMemory.prototype.getVideoGame = function (id) {
    for (const videoGame of this._videoGames)
        if (videoGame.id === id)
            return videoGame;
};

InMemory.prototype.getAllVideoGames = function () {
    return this._videoGames;
};

InMemory.prototype.addVideoGame = function (videoGame) {
    this._idCounter++;

    videoGame.id = this._idCounter;

    this._videoGames.push(videoGame);

    return videoGame;
};

InMemory.prototype.updateVideoGame = function (id, updatedVideoGame) {
    for (const index in this._videoGames)
        if (this._videoGames[index].id === id)
            this._videoGames[index] = updatedVideoGame;

    return updatedVideoGame;
};

module.exports = InMemory;
