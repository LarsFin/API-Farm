const VideoGamesService = require('../../lib/services/video-games');

const Query = require('../../lib/utils/query');

let videoGamesService;
let mockStorage;

beforeEach(() => {
    mockStorage = {};
    videoGamesService = new VideoGamesService(mockStorage);
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
    const date_released = '15/04/2004';
    const data = {
        name,
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms,
        date_released
    };
    const storedVideoGame = {};
    mockStorage.add = jest.fn(() => storedVideoGame);
    const successfulQuery = {};
    Query.success = jest.fn(() => successfulQuery);

    // Act
    const query = videoGamesService.add(data);

    // Assert
    expect(mockStorage.add).toHaveBeenCalledTimes(1);
    expect(mockStorage.add).toHaveBeenCalledWith(
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
            date_released
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
    const date_released = '15/04/2004';
    const data = {
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms,
        date_released
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

test('add should fail when data has an invalid date_released (1)', () => {
    // Arrange
    const name = 'Video Game Quest II';
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const artists = ['VG ARTISTS'];
    const composers = ['VG COMPS'];
    const platforms = ['PLATFORM X'];
    const date_released = '04/15/2004';
    const data = {
        name,
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms,
        date_released
    };
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.add(data);

    // Assert
    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(400, `The provided date_released '${date_released}' is invalid.`);

    expect(query).toBe(failedQuery);
});

test('add should fail when data has an invalid date_released (2)', () => {
    // Arrange
    const name = 'Video Game Quest II';
    const developers = ['VG DEVS'];
    const publishers = ['VG PUBS'];
    const directors = ['VG DIR'];
    const artists = ['VG ARTISTS'];
    const composers = ['VG COMPS'];
    const platforms = ['PLATFORM X'];
    const date_released = 'Last Monday';
    const data = {
        name,
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms,
        date_released
    };
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.add(data);

    // Assert
    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(400, `The provided date_released '${date_released}' is invalid.`);

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
    const date_released = '15/04/2004';
    const data = {
        name,
        developers,
        publishers,
        directors,
        artists,
        composers,
        platforms,
        testers,
        date_released
    };
    const failedQuery = {};
    Query.fail = jest.fn(() => failedQuery);

    // Act
    const query = videoGamesService.add(data);

    // Assert
    expect(Query.fail).toHaveBeenCalledTimes(1);
    expect(Query.fail).toHaveBeenCalledWith(400, 'The provided data has an invalid attribute \'testers\'.');

    expect(query).toBe(failedQuery);
});