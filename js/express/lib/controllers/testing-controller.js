const DataLoader = require('../utils/data-loader');

function TestingController(storage) {
    this._storage = storage;
}

TestingController.prototype.setup = async function (res) {
    try {
        const videoGames = await DataLoader.load();

        this._storage.constructor();

        for (const videoGame of videoGames)
            this._storage.addVideoGame(videoGame);

        res.okText('Successfully loaded data.');
    } catch (error) {
        console.error(error);
        res.internalServerError();
    }
}

module.exports = TestingController;
