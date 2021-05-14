const InMemory = require('../../lib/storage/in-memory');

const Clone = require('../../lib/utils/clone');

let inMemory;

beforeEach(() => {
    inMemory = new InMemory();
});

// Get Video Game

test('getVideoGame should return video game with specified id', () => {
    // Arrange
    const videoGame1 = { id: 1 };
    const videoGame2 = { id: 3 };
    const videoGame3 = { id: 4 };
    inMemory._videoGames = [videoGame1, videoGame2, videoGame3];
    const clonedVideoGame = {};
    Clone.object = jest.fn(() => clonedVideoGame);

    // Act
    const videoGame = inMemory.getVideoGame(3);

    // Assert
    expect(videoGame).toBe(clonedVideoGame);

    expect(Clone.object).toHaveBeenCalledTimes(1);
    expect(Clone.object).toHaveBeenCalledWith(videoGame2);
});

test('getVideoGame should return null when the video game with id could not be found', () => {
    // Act
    const videoGame = inMemory.getVideoGame(5);
    Clone.object = jest.fn();

    // Assert
    expect(videoGame).toBe(undefined);
    expect(Clone.object).toHaveBeenCalledTimes(0);
});

// Get All Video Games

test('getAllVideoGames should return internal video games array', () => {
    // Arrange
    const clonedVideoGames = [];
    Clone.object = jest.fn(() => clonedVideoGames);

    // Act
    const videoGames = inMemory.getAllVideoGames();

    // Assert
    expect(videoGames).toBe(clonedVideoGames);

    expect(Clone.object).toHaveBeenCalledTimes(1);
    expect(Clone.object).toHaveBeenCalledWith(inMemory._videoGames);
});

// Add Video Game

test('addVideoGame should increment and set id, then push to video games array', () => {
    // Arrange
    const videoGame = {};
    const clonedVideoGame = {};
    Clone.object = jest.fn(() => clonedVideoGame);

    // Act
    const storedVideoGame = inMemory.addVideoGame(videoGame);

    // Assert
    expect(inMemory._videoGames.length).toBe(1);
    expect(inMemory._videoGames[0]).toBe(storedVideoGame);

    expect(videoGame.id).toBe(undefined);
    expect(storedVideoGame.id).toBe(1);

    expect(Clone.object).toHaveBeenCalledTimes(2);
    expect(Clone.object).toHaveBeenNthCalledWith(1, videoGame);
    expect(Clone.object).toHaveBeenNthCalledWith(2, clonedVideoGame);
});

// Update Video Game

test('updateVideoGame should find video game with passed id and replace with updated version', () => {
    // Arrange
    const videoGame1 = { id: 1 };
    const videoGame2 = { id: 4 };
    const videoGame3 = { id: 8 };
    inMemory._videoGames = [videoGame1, videoGame2, videoGame3];
    const updatedVideoGame = {};
    const clonedUpdatedVideoGame = {};
    Clone.object = jest.fn(() => clonedUpdatedVideoGame);

    // Act
    const storedUpdatedVideoGame = inMemory.updateVideoGame(4, updatedVideoGame);

    // Assert
    expect(storedUpdatedVideoGame).toBe(updatedVideoGame);
    expect(inMemory._videoGames[1]).toBe(clonedUpdatedVideoGame);

    expect(Clone.object).toHaveBeenCalledTimes(1);
    expect(Clone.object).toHaveBeenCalledWith(updatedVideoGame);
});

test('deleteVideoGame should find video game with passed id and delete it', () => {
    // Arrange
    const videoGame1 = { id: 1 };
    const videoGame2 = { id: 2 };
    const videoGame3 = { id: 3 };
    inMemory._videoGames = [videoGame1, videoGame2, videoGame3];

    // Act
    const videoGame = inMemory.deleteVideoGame(2);

    //Assert
    expect(videoGame).toBe(videoGame2);

    expect(inMemory._videoGames.includes(videoGame2)).toBe(false);
});

test('deleteVideoGame should return null when the video game with id could not be deleted', () => {
    // Act
    const videoGame = inMemory.deleteVideoGame(3);

    // Assert
    expect(videoGame).toBe(undefined);
});
