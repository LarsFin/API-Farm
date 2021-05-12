const Controller = require('../../lib/controllers/controller');

let controller;
let mockVideoGamesService;

beforeEach(() => {
    mockVideoGamesService = {};
    controller = new Controller(mockVideoGamesService);
});

// Get

test('get should retrieve video game with parametised id and respond 200', () => {
    // Arrange
    const rawId = '5';
    const req = { params: { id: rawId } };
    const id = 5;
    const videoGame = {};
    const query = {
        result: videoGame
    };
    mockVideoGamesService.get = jest.fn(() => query);
    const res = {};
    res.ok = jest.fn();

    // Act
    controller.get(req, res);

    // Assert
    expect(res.ok).toHaveBeenCalledTimes(1);
    expect(res.ok).toHaveBeenCalledWith(videoGame);

    expect(mockVideoGamesService.get).toHaveBeenCalledTimes(1);
    expect(mockVideoGamesService.get).toHaveBeenCalledWith(id);
});

test('get should not accept invalid id values and respond 400', () => {
    // Arrange
    const rawId = 'invalid!';
    const req = { params: { id: rawId } };
    const res = {};
    res.badRequest = jest.fn();

    // Act
    controller.get(req, res);

    // Assert
    expect(res.badRequest).toHaveBeenCalledTimes(1);
    expect(res.badRequest).toHaveBeenCalledWith(`The provided id '${rawId}' is invalid.`);
});

test('get should check for when a video game could not be found and respond 404', () => {
    // Arrange
    const rawId = '99';
    const req = { params: { id: rawId } };
    const id = 99;
    const failReason = 'FAILED TO FIND VIDEO GAME!';
    const query = {
        code: 404,
        result: failReason
    };
    mockVideoGamesService.get = jest.fn(() => query);
    const res = {};
    res.notFound = jest.fn();

    // Act
    controller.get(req, res);

    // Assert
    expect(res.notFound).toHaveBeenCalledTimes(1);
    expect(res.notFound).toHaveBeenCalledWith(failReason);

    expect(mockVideoGamesService.get).toHaveBeenCalledTimes(1);
    expect(mockVideoGamesService.get).toHaveBeenCalledWith(id);
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
    controller.add(req, res);

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
    controller.add(req, res);

    // Assert
    expect(mockVideoGamesService.add).toHaveBeenCalledTimes(1);
    expect(mockVideoGamesService.add).toHaveBeenCalledWith(body);

    expect(res.badRequest).toHaveBeenCalledTimes(1);
    expect(res.badRequest).toHaveBeenCalledWith(failReason);
});

// Put

test('put should make an update query and respond 200', () => {
    // Arrange
    const rawId = '5';
    const body = {};
    const req = {
        params: { id: rawId },
        body
    };
    const videoGame = {};
    const query = {
        code: 0,
        result: videoGame
    };
    mockVideoGamesService.update = jest.fn(() => query);
    const res = {};
    res.ok = jest.fn();

    // Act
    controller.put(req, res);

    // Assert
    expect(mockVideoGamesService.update).toHaveBeenCalledTimes(1);
    expect(mockVideoGamesService.update).toHaveBeenCalledWith(5, body);

    expect(res.ok).toHaveBeenCalledTimes(1);
    expect(res.ok).toHaveBeenCalledWith(videoGame);
});

test('put should not accept an invalid id and respond 400', () => {
    // Arrange
    const rawId = 'invalid!';
    const body = {};
    const req = {
        params: { id: rawId },
        body
    };
    const res = {};
    res.badRequest = jest.fn();

    // Act
    controller.put(req, res);

    // Assert
    expect(res.badRequest).toHaveBeenCalledTimes(1);
    expect(res.badRequest).toHaveBeenCalledWith(`The provided id '${rawId}' is invalid.`);
});

test('put should fail when invalid data is provided and respond 400', () => {
    // Arrange
    const rawId = '5';
    const body = {};
    const req = {
        params: { id: rawId },
        body
    };
    const failReason = 'Invalid data provided!';
    const query = {
        code: 400,
        result: failReason
    };
    mockVideoGamesService.update = jest.fn(() => query);
    const res = {};
    res.badRequest = jest.fn();

    // Act
    controller.put(req, res);

    // Assert
    expect(mockVideoGamesService.update).toHaveBeenCalledTimes(1);
    expect(mockVideoGamesService.update).toHaveBeenCalledWith(5, body);

    expect(res.badRequest).toHaveBeenCalledTimes(1);
    expect(res.badRequest).toHaveBeenCalledWith(failReason);
});

test('put should fail when video game not found and respond 404', () => {
    // Arrange
    const rawId = '5';
    const body = {};
    const req = {
        params: { id: rawId },
        body
    };
    const failReason = 'Video Game could not be found!';
    const query = {
        code: 404,
        result: failReason
    };
    mockVideoGamesService.update = jest.fn(() => query);
    const res = {};
    res.notFound = jest.fn();

    // Act
    controller.put(req, res);

    // Assert
    expect(mockVideoGamesService.update).toHaveBeenCalledTimes(1);
    expect(mockVideoGamesService.update).toHaveBeenCalledWith(5, body);

    expect(res.notFound).toHaveBeenCalledTimes(1);
    expect(res.notFound).toHaveBeenCalledWith(failReason);
});
