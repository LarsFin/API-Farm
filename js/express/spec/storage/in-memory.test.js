const InMemory = require('../../lib/storage/in-memory');

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

    // Act
    const videoGame = inMemory.getVideoGame(3);

    // Assert
    expect(videoGame).toBe(videoGame2);
});

test('getVideoGame should return null when the video game with id could not be found', () => {
    // Act
    const videoGame = inMemory.getVideoGame(5);

    // Assert
    expect(videoGame).toBe(undefined);
});

// Get All Video Games

test('getAllVideoGames should return internal video games array', () => {
    // Act
    const videoGames = inMemory.getAllVideoGames();

    // Assert
    expect(videoGames).toBe(inMemory._videoGames);
});

// Add Video Game

test('addVideoGame should increment and set id, then push to video games array', () => {
    // Arrange
    const videoGame = {};

    // Act
    const storedVideoGame = inMemory.addVideoGame(videoGame);

    // Assert
    expect(inMemory._videoGames.length).toBe(1);
    expect(inMemory._videoGames[0]).toBe(videoGame);
    expect(storedVideoGame.id).toBe(1);
});
