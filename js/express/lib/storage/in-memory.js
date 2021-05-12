const Clone = require('../utils/clone');

function InMemory () {
    this._idCounter = 0;
    this._videoGames = [];
}

InMemory.prototype.getVideoGame = function (id) {
    for (const videoGame of this._videoGames)
        if (videoGame.id === id)
            return Clone.object(videoGame);
};

InMemory.prototype.getAllVideoGames = function () {
    return Clone.object(this._videoGames);
};

InMemory.prototype.addVideoGame = function (videoGame) {
    this._idCounter++;

    videoGame = Clone.object(videoGame);

    videoGame.id = this._idCounter;

    this._videoGames.push(videoGame);

    return Clone.object(videoGame);
};

InMemory.prototype.updateVideoGame = function (id, updatedVideoGame) {
    for (const index in this._videoGames)
        if (this._videoGames[index].id === id)
            this._videoGames[index] = Clone.object(updatedVideoGame);

    return updatedVideoGame;
};

module.exports = InMemory;
