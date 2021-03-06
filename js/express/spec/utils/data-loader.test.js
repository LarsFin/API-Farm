const DataLoader = require('../../lib/utils/data-loader');
const fs = require('fs');

test('load should read, parse and resolve data', async () => {
    // Arrange
    const dataPath = './path/to/data-file.json';
    const data = {};
    fs.readFile = jest.fn((_, __, callback) => callback(null, data));
    const sampleData = {};
    JSON.parse = jest.fn(() => sampleData);

    // Act
    const loadPromise = DataLoader.load(dataPath);

    // Assert
    await expect(loadPromise).resolves.toBe(sampleData);

    expect(fs.readFile).toHaveBeenCalledTimes(1);
    expect(fs.readFile).toHaveBeenCalledWith(dataPath, 'utf8', expect.anything());

    expect(JSON.parse).toHaveBeenCalledTimes(1);
    expect(JSON.parse).toHaveBeenCalledWith(data);
});

test('load should reject a file reading error', async () => {
    // Arrange
    const dataPath = './path/to/data-file.json';
    const err = {};
    fs.readFile = jest.fn((_, __, callback) => callback(err));

    // Act
    const loadPromise = DataLoader.load(dataPath);

    // Assert
    await expect(loadPromise).rejects.toBe(err);

    expect(fs.readFile).toHaveBeenCalledTimes(1);
    expect(fs.readFile).toHaveBeenCalledWith(dataPath, 'utf8', expect.anything());
});

test('load should reject a JSON parsing error', async () => {
    // Arrange
    const dataPath = './path/to/data-file.json';
    const data = {};
    fs.readFile = jest.fn((_, __, callback) => callback(null, data));
    const err = {};
    JSON.parse = jest.fn(() => {
        throw err;
    });

    // Act
    const loadPromise = DataLoader.load(dataPath);

    // Assert
    await expect(loadPromise).rejects.toBe(err);

    expect(fs.readFile).toHaveBeenCalledTimes(1);
    expect(fs.readFile).toHaveBeenCalledWith(dataPath, 'utf8', expect.anything());

    expect(JSON.parse).toHaveBeenCalledTimes(1);
    expect(JSON.parse).toHaveBeenCalledWith(data);
});
