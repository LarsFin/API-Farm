const Controller = require('../../lib/controllers/controller');

let controller;
let mockVideoGamesService;

beforeEach(() => {
    mockVideoGamesService = {};
    controller = new Controller(mockVideoGamesService);
});

// Get All

test('getAll should retrieve video games from service and respond 200', () => {
    // Arrange
    const res = {};
    res.ok = jest.fn();
    const videoGames = [];
    const query = {
        result: videoGames
    };
    mockVideoGamesService.getAll = jest.fn(() => query);

    // Act
    controller.getAll(res);

    // Assert
    expect(res.ok).toHaveBeenCalledTimes(1);
    expect(res.ok).toHaveBeenCalledWith(videoGames);
});

// Add

test('add should add body data to service and respond 200', () => {
    // Arrange
    const body = {};
    const req = {
        body
    };
    const res = {};
    res.created = jest.fn();
    const storedVideoGame = {};
    const query = {
        code: 0,
        result: storedVideoGame
    };
    mockVideoGamesService.add = jest.fn(() => query);

    // Act
    controller.add(req, res)

    // Assert
    expect(mockVideoGamesService.add).toHaveBeenCalledTimes(1);
    expect(mockVideoGamesService.add).toHaveBeenCalledWith(body);

    expect(res.created).toHaveBeenCalledTimes(1);
    expect(res.created).toHaveBeenCalledWith(storedVideoGame);
});

test('add should add body data to service and respond 400', () => {
    // Arrange
    const body = {};
    const req = {
        body
    };
    const res = {};
    res.badRequest = jest.fn();
    const failReason = 'INVALID VIDEO GAME DATA';
    const query = {
        code: 400,
        result: failReason
    };
    mockVideoGamesService.add = jest.fn(() => query);

    // Act
    controller.add(req, res)

    // Assert
    expect(mockVideoGamesService.add).toHaveBeenCalledTimes(1);
    expect(mockVideoGamesService.add).toHaveBeenCalledWith(body);

    expect(res.badRequest).toHaveBeenCalledTimes(1);
    expect(res.badRequest).toHaveBeenCalledWith(failReason);
});