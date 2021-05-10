function InMemory () {
    this._idCounter = 0;
    this._videoGames = [];
}

InMemory.prototype.getAllVideoGames = function () {
    return this._videoGames;
};

InMemory.prototype.addVideoGame = function (videoGame) {
    this._idCounter++;

    videoGame.id = this._idCounter;

    this._videoGames.push(videoGame);

    return videoGame;
};

module.exports = InMemory;
