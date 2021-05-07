const InMemory = require('../../lib/storage/in-memory');

test('getAllVideoGames should return internal video games array', () => {
    // Arrange
    const inMemory = new InMemory();

    // Act
    const videoGames = inMemory.getAllVideoGames();

    // Assert
    expect(videoGames).toBe(inMemory._videoGames);
});
