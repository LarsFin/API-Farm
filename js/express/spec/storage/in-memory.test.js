const InMemory = require('../../lib/storage/in-memory');

let inMemory;

beforeEach(() => {
    inMemory = new InMemory();
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
