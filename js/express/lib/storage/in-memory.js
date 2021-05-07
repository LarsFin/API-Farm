function InMemory () {
    this._videoGames = [];
}

InMemory.prototype.getAllVideoGames = function () {
    return this._videoGames;
};

module.exports = InMemory;
