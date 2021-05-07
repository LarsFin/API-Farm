const VideoGamesService = require('../../lib/services/video-games');

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

    // Act
    const query = videoGamesService.getAll();

    // Assert
    expect(query.code).toBe(0);
    expect(query.result).toBe(videoGames);
});
