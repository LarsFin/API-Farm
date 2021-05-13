const { expect } = require('@jest/globals');
const Config = require('../../lib/utils/config');

const DataLoader = require('../../lib/utils/data-loader');
const InvalidEnvironmentError = require('../../lib/errors/invalid-environment');

test('fromEnvironment should return configuration as json', () => {
    // Arrange
    const expectedConfig = {};
    DataLoader.load = jest.fn(() => expectedConfig);

    // Act
    const config = Config.fromEnvironment('DEV');

    // Assert
    expect(config).toBe(expectedConfig);

    expect(DataLoader.load).toHaveBeenCalledTimes(1);
    expect(DataLoader.load).toHaveBeenLastCalledWith('config.dev.json');
});

test('fromEnvironment should throw an error when passed an invalid environmnet', () => {
    // Arrange
    const invalidEnvironment = 'POST';
    const expectedError = new InvalidEnvironmentError(invalidEnvironment);

    // Assert
    expect(() => {
        Config.fromEnvironment(invalidEnvironment);
    }).toThrowError(expectedError);
});
