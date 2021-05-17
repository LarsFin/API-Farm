const VideoGamesService = require('../../lib/services/video-games');

const ApiFarmDate = require('../../lib/utils/api-farm-date');
const Query = require('../../lib/utils/query');

let videoGamesService;
let mockStorage;

beforeEach(() => {
    mockStorage = {};
    videoGamesService = new VideoGamesService(mockStorage);
});

// Get

test('get should query storage and return video game', () => {
    // Arrange
    const id = 5;
    const videoGame = {};
    mockStorage.getVideoGame = jest.fn(() => videoGame);
    const successfulQuery = {};
    Query.success = jest.fn(() => successfulQuery);

    // Act
    const query = videoGamesService.get(id);

    // Assert
    expect(query).toBe(successfulQuery);

    expect(mockStorage.getVideoGame).toHaveBeenCalledTimes(1);
    expect(mockStorage.getVideoGame).toHaveBeenCalledWith(id);

    expect(Query.success).toHaveBeenCalledTimes(1);
    expect(Query.success).toHaveBeenCalledWith(videoGame);
});

test('get should query storage and return failure when video game not found', () => {
    // Arrange
    const id = 99;
    mockStorage.getVideoGame = jest.fn();
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.get(id);

    // Assert
    expect(query).toBe(failedQuery);

    expect(mockStorage.getVideoGame).toHaveBeenCalledTimes(1);
    expect(mockStorage.getVideoGame).toHaveBeenCalledWith(id);

    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(404, `No video game with id '${id}' could be found.`);
});

// Get All

test('getAll should query storage and return video games', () => {
    // Arrange
    const videoGames = [];
    mockStorage.getAllVideoGames = jest.fn(() => videoGames);
    const successfulQuery = {};
    Query.success = jest.fn(() => successfulQuery);

    // Act
    const query = videoGamesService.getAll();

    // Assert
    expect(Query.success).toHaveBeenCalledTimes(1);
    expect(Query.success).toHaveBeenCalledWith(videoGames);

    expect(query).toBe(successfulQuery);
});

// Add

test('add should instance a video game object and add to storage', () => {
    // Arrange
    const name = 'Video Game Quest II';
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const artists = ['VG ARTISTS'];
    const composers = ['VG COMPS'];
    const platforms = ['PLATFORM X'];
    const dateReleased = '15/04/2004';
    const data = {
        name,
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms,
        date_released: dateReleased
    };
    ApiFarmDate.isValid = jest.fn(() => true);
    const storedVideoGame = {};
    mockStorage.addVideoGame = jest.fn(() => storedVideoGame);
    const successfulQuery = {};
    Query.success = jest.fn(() => successfulQuery);

    // Act
    const query = videoGamesService.add(data);

    // Assert
    expect(mockStorage.addVideoGame).toHaveBeenCalledTimes(1);
    expect(mockStorage.addVideoGame).toHaveBeenCalledWith(
        expect.objectContaining({
            name,
            developers,
            publishers,
            directors,
            producers: [],
            designers: [],
            programmers: [],
            artists,
            composers,
            platforms,
            date_released: dateReleased
        })
    );

    expect(Query.success).toHaveBeenCalledTimes(1);
    expect(Query.success).toHaveBeenCalledWith(storedVideoGame);

    expect(query).toBe(successfulQuery);
});

test('add should fail when data has no name', () => {
    // Arrange
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const artists = ['VG ARTISTS'];
    const composers = ['VG COMPS'];
    const platforms = ['PLATFORM X'];
    const dateReleased = '15/04/2004';
    const data = {
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms,
        date_released: dateReleased
    };
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.add(data);

    // Assert
    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(400, 'A name is required for a video game.');

    expect(query).toBe(failedQuery);
});

test('add should fail when data has no date_released', () => {
    // Arrange
    const name = 'Video Game Quest II';
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const artists = ['VG ARTISTS'];
    const composers = ['VG COMPS'];
    const platforms = ['PLATFORM X'];
    const data = {
        name,
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms
    };
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.add(data);

    // Assert
    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(400, 'A date_released is required for a video game.');

    expect(query).toBe(failedQuery);
});

test('add should fail when data has an invalid date_released', () => {
    // Arrange
    const name = 'Video Game Quest II';
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const artists = ['VG ARTISTS'];
    const composers = ['VG COMPS'];
    const platforms = ['PLATFORM X'];
    const dateReleased = '04/15/2004';
    const data = {
        name,
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms,
        date_released: dateReleased
    };
    ApiFarmDate.isValid = jest.fn(() => false);
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.add(data);

    // Assert
    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(400, `The provided date_released '${dateReleased}' is invalid.`);

    expect(query).toBe(failedQuery);
});

test('add should fail when data has an invalid attribute', () => {
    // Arrange
    const name = 'Video Game Quest II';
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const artists = ['VG ARTISTS'];
    const composers = ['VG COMPS'];
    const platforms = ['PLATFORM X'];
    const testers = ['VG TESTERS'];
    const dateReleased = '15/04/2004';
    const data = {
        name,
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms,
        testers,
        date_released: dateReleased
    };
    ApiFarmDate.isValid = jest.fn(() => true);
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.add(data);

    // Assert
    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(400, 'The provided data has an invalid attribute \'testers\'.');

    expect(query).toBe(failedQuery);
});

// Update

test('update should update video game properties in storage and return successful query', () => {
    // Arrange
    const name = 'Video Game to Update';
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const dateReleased = '15/04/2004';
    const videoGame = {
        name,
        developers,
        publishers,
        directors,
        date_released: dateReleased
    };
    const id = 3;
    const updatedName = 'Updated Video Game';
    const updatedDirectors = ['NEW DIR'];
    const updatedDateReleased = '04/06/2004';
    const data = {
        name: updatedName,
        directors: updatedDirectors,
        date_released: updatedDateReleased
    };
    const successfulQuery1 = {
        code: 0,
        result: videoGame
    };
    videoGamesService.get = jest.fn(() => successfulQuery1);
    ApiFarmDate.isValid = jest.fn(() => true);
    const updatedVideoGame = {};
    mockStorage.updateVideoGame = jest.fn(() => updatedVideoGame);
    const successfulQuery2 = {
        code: 0,
        result: updatedVideoGame
    };
    Query.success = jest.fn(() => successfulQuery2);

    // Act
    const query = videoGamesService.update(id, data);

    // Assert
    expect(query).toBe(successfulQuery2);

    expect(videoGamesService.get).toHaveBeenCalledTimes(1);
    expect(videoGamesService.get).toHaveBeenCalledWith(id);

    expect(ApiFarmDate.isValid).toHaveBeenCalledTimes(1);
    expect(ApiFarmDate.isValid).toHaveBeenCalledWith(updatedDateReleased);

    expect(mockStorage.updateVideoGame).toHaveBeenCalledTimes(1);
    expect(mockStorage.updateVideoGame).toHaveBeenCalledWith(
        id,
        expect.objectContaining({
            name: updatedName,
            developers,
            publishers,
            directors: updatedDirectors,
            date_released: updatedDateReleased
        })
    );

    expect(Query.success).toHaveBeenCalledTimes(1);
    expect(Query.success).toHaveBeenCalledWith(updatedVideoGame);
});

test('update should return failed query when does not exist in storage', () => {
    // Arrange
    const id = 99;
    const failedQuery = {
        code: 404,
        result: 'Could not find video game!'
    };
    videoGamesService.get = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.update(id, {});

    // Assert
    expect(query).toBe(failedQuery);

    expect(videoGamesService.get).toHaveBeenCalledTimes(1);
    expect(videoGamesService.get).toHaveBeenCalledWith(id);
});

test('update should return failed query when data contains invalid attribute', () => {
    // Arrange
    const name = 'Video Game to Update';
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const dateReleased = '15/04/2004';
    const videoGame = {
        name,
        developers,
        publishers,
        directors,
        date_released: dateReleased
    };
    const id = 3;
    const updatedName = 'Updated Video Game';
    const updatedDirectors = ['NEW DIR'];
    const updatedTesters = ['INVALID TESTERS'];
    const updatedDateReleased = '04/06/2004';
    const data = {
        name: updatedName,
        directors: updatedDirectors,
        testers: updatedTesters,
        date_released: updatedDateReleased
    };
    const successfulQuery1 = {
        code: 0,
        result: videoGame
    };
    videoGamesService.get = jest.fn(() => successfulQuery1);
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.update(id, data);

    // Assert
    expect(query).toBe(failedQuery);

    expect(videoGamesService.get).toHaveBeenCalledTimes(1);
    expect(videoGamesService.get).toHaveBeenCalledWith(id);

    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(400, 'The provided data has an invalid attribute \'testers\'.');
});

test('update should return failed query when provided date_released is invalid', () => {
    // Arrange
    const name = 'Video Game to Update';
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const dateReleased = '15/04/2004';
    const videoGame = {
        name,
        developers,
        publishers,
        directors,
        date_released: dateReleased
    };
    const id = 3;
    const updatedName = 'Updated Video Game';
    const updatedDirectors = ['NEW DIR'];
    const updatedDateReleased = '12/30/2004';
    const data = {
        name: updatedName,
        directors: updatedDirectors,
        date_released: updatedDateReleased
    };
    const successfulQuery1 = {
        code: 0,
        result: videoGame
    };
    videoGamesService.get = jest.fn(() => successfulQuery1);
    ApiFarmDate.isValid = jest.fn(() => false);
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.update(id, data);

    // Assert
    expect(query).toBe(failedQuery);

    expect(videoGamesService.get).toHaveBeenCalledTimes(1);
    expect(videoGamesService.get).toHaveBeenCalledWith(id);

    expect(ApiFarmDate.isValid).toHaveBeenCalledTimes(1);
    expect(ApiFarmDate.isValid).toHaveBeenCalledWith(updatedDateReleased);

    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(400, `The provided date_released '${updatedDateReleased}' is invalid.`);
});

// Delete

test('delete should query storage and remove video game', () => {
    // Arrange
    const id = 2;
    const videoGame = {};
    mockStorage.deleteVideoGame = jest.fn(() => videoGame);
    const successfulQuery = {};
    Query.success = jest.fn(() => successfulQuery);

    // Act
    const query = videoGamesService.delete(id);

    // Assert
    expect(query).toBe(successfulQuery);

    expect(mockStorage.deleteVideoGame).toHaveBeenCalledTimes(1);
    expect(mockStorage.deleteVideoGame).toHaveBeenCalledWith(id);

    expect(Query.success).toHaveBeenCalledTimes(1);
    expect(Query.success).toHaveBeenCalledWith(`Deleted video game with id '${id}'.`);
});

test('delete should query storage and return failure when video game not found', () => {
    // Arrange
    const id = 99;
    mockStorage.deleteVideoGame = jest.fn();
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.delete(id);

    // Assert
    expect(query).toBe(failedQuery);

    expect(mockStorage.deleteVideoGame).toHaveBeenCalledTimes(1);
    expect(mockStorage.deleteVideoGame).toHaveBeenCalledWith(id);

    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(404, `No video game with id '${id}' could be found.`);
});
