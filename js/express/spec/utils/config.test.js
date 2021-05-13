const { expect, test } = require('@jest/globals');
const Config = require('../../lib/utils/config');

const DataLoader = require('../../lib/utils/data-loader');
const InvalidEnvironmentError = require('../../lib/errors/invalid-environment');

test('fromEnvironment should return configuration as json', async () => {
    // Arrange
    const expectedConfig = {};
    DataLoader.load = jest.fn(() => Promise.resolve(expectedConfig));

    // Act
    const configPromise = Config.fromEnvironment('DEV');

    // Assert
    await expect(configPromise).resolves.toBe(expectedConfig);

    expect(DataLoader.load).toHaveBeenCalledTimes(1);
    expect(DataLoader.load).toHaveBeenLastCalledWith('config.dev.json');
});

test('fromEnvironment should reject a data loading failure', async () => {
    // Arrange
    const dataLoadFailure = 'Failed to load data!';
    DataLoader.load = jest.fn(() => Promise.reject(dataLoadFailure));

    // Act
    const configPromise = Config.fromEnvironment('DEV');

    // Assert
    await expect(configPromise).rejects.toBe(dataLoadFailure);

    expect(DataLoader.load).toHaveBeenCalledTimes(1);
    expect(DataLoader.load).toHaveBeenLastCalledWith('config.dev.json');
});

test('fromEnvironment should reject an invalid environmnet error', () => {
    // Arrange
    const invalidEnvironment = 'POST';
    const expectedError = new InvalidEnvironmentError(invalidEnvironment);

    // Act
    const configPromise = Config.fromEnvironment(invalidEnvironment);

    // Assert
    expect(configPromise).rejects.toMatchObject(expectedError);    
});
