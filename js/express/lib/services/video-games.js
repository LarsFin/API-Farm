const ApiFarmDate = require('../utils/api-farm-date');
const Query = require('../utils/query');

function VideoGamesService (storage) {
    this._storage = storage;
}

const newVideoGame = (name, dateReleased) => {
    return {
        name,
        developers: [],
        publishers: [],
        directors: [],
        producers: [],
        designers: [],
        programmers: [],
        artists: [],
        composers: [],
        platforms: [],
        date_released: dateReleased
    };
};

VideoGamesService.prototype.getAll = function () {
    const videoGames = this._storage.getAllVideoGames();

    return Query.success(videoGames);
};

VideoGamesService.prototype.add = function (data) {
    const videoGameName = data.name;

    if (!videoGameName)
        return Query.fail(400, 'A name is required for a video game.');

    const videoGameDateReleased = data.date_released;

    if (!videoGameDateReleased)
        return Query.fail(400, 'A date_released is required for a video game.');

    if (!ApiFarmDate.isValid(videoGameDateReleased))
        return Query.fail(400, `The provided date_released '${videoGameDateReleased}' is invalid.`);

    let videoGame = newVideoGame(videoGameName, videoGameDateReleased);

    for (const [key, value] of Object.entries(data)) {
        if (['name', 'date_released'].includes(key))
            continue;

        if (!Object.prototype.hasOwnProperty.call(videoGame, key))
            return Query.fail(400, `The provided data has an invalid attribute '${key}'.`);

        videoGame[key] = value;
    }

    videoGame = this._storage.add(videoGame);

    return Query.success(videoGame);
};

module.exports = VideoGamesService;
