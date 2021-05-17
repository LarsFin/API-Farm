const ApiFarmDate = require('../utils/api-farm-date');
const Query = require('../utils/query');

function VideoGamesService (storage) {
    this._storage = storage;
}

const initVideoGame = () => {
    return {
        name: '',
        developers: [],
        publishers: [],
        directors: [],
        producers: [],
        designers: [],
        programmers: [],
        artists: [],
        composers: [],
        platforms: [],
        date_released: ''
    };
};

const mapProperties = (videoGame, data) => {
    const validProperties = Object.keys(initVideoGame());

    for (const [key, value] of Object.entries(data)) {
        if (!validProperties.includes(key))
            return Query.fail(400, `The provided data has an invalid attribute '${key}'.`);

        if (key === 'date_released' && !ApiFarmDate.isValid(value))
            return Query.fail(400, `The provided date_released '${value}' is invalid.`);

        videoGame[key] = value;
    }
};

VideoGamesService.prototype.get = function (id) {
    const videoGame = this._storage.getVideoGame(id);

    if (!videoGame)
        return Query.fail(404, `No video game with id '${id}' could be found.`);

    return Query.success(videoGame);
};

VideoGamesService.prototype.getAll = function () {
    const videoGames = this._storage.getAllVideoGames();

    return Query.success(videoGames);
};

VideoGamesService.prototype.add = function (data) {
    if (!data.name)
        return Query.fail(400, 'A name is required for a video game.');

    if (!data.date_released)
        return Query.fail(400, 'A date_released is required for a video game.');

    let videoGame = initVideoGame();

    const mappingIssue = mapProperties(videoGame, data);

    if (mappingIssue)
        return mappingIssue;

    videoGame = this._storage.addVideoGame(videoGame);

    return Query.success(videoGame);
};

VideoGamesService.prototype.update = function (id, data) {
    const query = this.get(id);

    if (query.code === 404)
        return query;

    let videoGame = query.result;

    const mappingIssue = mapProperties(videoGame, data);

    if (mappingIssue)
        return mappingIssue;

    videoGame = this._storage.updateVideoGame(id, videoGame);

    return Query.success(videoGame);
};

VideoGamesService.prototype.delete = function (id) {
    const videoGame = this._storage.deleteVideoGame(id);

    if (!videoGame)
        return Query.fail(404, `No video game with id '${id}' could be found.`);

    return Query.success(`Deleted video game with id '${id}'.`);
};

module.exports = VideoGamesService;
