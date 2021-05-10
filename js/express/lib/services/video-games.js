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

// date should always be in format DD/MM/YYYY
const isValidDate = dateReleased => {
    const dateElements = dateReleased.split('/');

    if (dateElements.length !== 3)
        return false;

    return !isNaN(new Date(`${dateElements[1]}/${dateElements[0]}/${dateElements[2]}`));
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

    if (!isValidDate(videoGameDateReleased))
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
