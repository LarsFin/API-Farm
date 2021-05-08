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
    mockRes.status = jest.fn();
    mockRes.header = jest.fn();
    mockRes.json = jest.fn();
    const videoGames = [];
    const query = {
        result: videoGames
    };
    mockVideoGamesService.getAll = jest.fn(() => query);

    // Act
    controller.getAll(mockRes);

    // Assert
    expect(mockRes.status).toHaveBeenCalledTimes(1);
    expect(mockRes.status).toHaveBeenCalledWith(200);

    expect(mockRes.header).toHaveBeenCalledTimes(1);
    expect(mockRes.header).toHaveBeenCalledWith('Content-Type', 'application/json');

    expect(mockRes.json).toHaveBeenCalledTimes(1);
    expect(mockRes.json).toHaveBeenCalledWith(videoGames);
});
