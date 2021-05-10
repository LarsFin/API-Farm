const VideoGamesService = require('../../lib/services/video-games');

const Query = require('../../lib/utils/query');

let videoGamesService;
let mockStorage;

beforeEach(() => {
    mockStorage = {};
    videoGamesService = new VideoGamesService(mockStorage);
});

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
