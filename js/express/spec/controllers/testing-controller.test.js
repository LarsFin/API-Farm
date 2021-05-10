const TestingController = require('../../lib/controllers/testing-controller');

const DataLoader = require('../../lib/utils/data-loader');

let testingController;
let mockStorage;

beforeEach(() => {
    mockStorage = {};
    testingController = new TestingController(mockStorage);
});

test('setup should load data into storage and respond success', async () => {
    // Arrange
    const videoGame1 = {};
    const videoGame2 = {};
    const videoGames = [videoGame1, videoGame2];
    DataLoader.load = jest.fn(() => Promise.resolve(videoGames));
    mockStorage.constructor = jest.fn();
    mockStorage.addVideoGame = jest.fn();
    const res = {};
    res.okText = jest.fn();

    // Act
    await testingController.setup(res);

    // Assert
    expect(DataLoader.load).toHaveBeenCalledTimes(1);

    expect(mockStorage.constructor).toHaveBeenCalledTimes(1);

    expect(mockStorage.addVideoGame).toHaveBeenCalledTimes(2);
    expect(mockStorage.addVideoGame).toHaveBeenNthCalledWith(1, videoGame1);
    expect(mockStorage.addVideoGame).toHaveBeenNthCalledWith(1, videoGame2);

    expect(res.okText).toHaveBeenCalledTimes(1);
    expect(res.okText).toHaveBeenCalledWith('Successfully loaded data.');
});

test('setup should respond failure when data couldn\'t load', async () => {
    // Arrange
    const error = {};
    DataLoader.load = jest.fn(() => Promise.reject(error));
    console.error = jest.fn();
    const res = {};
    res.internalServerError = jest.fn();

    // Act
    await testingController.setup(res);

    // Assert
    expect(DataLoader.load).toHaveBeenCalledTimes(1);

    expect(console.error).toHaveBeenCalledTimes(1);
    expect(console.error).toHaveBeenCalledWith(error);

    expect(res.internalServerError).toHaveBeenCalledTimes(1);
});
