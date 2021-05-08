const Controller = require('../../lib/controllers/controller');

let controller;
let mockVideoGamesService;

beforeEach(() => {
    mockVideoGamesService = {};
    controller = new Controller(mockVideoGamesService);
});

test('getAll should retrieve video games from service and respond 200', () => {
    // Arrange
    const mockRes = {};
    mockRes.ok = jest.fn();
    const videoGames = [];
    const query = {
        result: videoGames
    };
    mockVideoGamesService.getAll = jest.fn(() => query);

    // Act
    controller.getAll(mockRes);

    // Assert
    expect(mockRes.ok).toHaveBeenCalledTimes(1);
    expect(mockRes.ok).toHaveBeenCalledWith(videoGames);
});
