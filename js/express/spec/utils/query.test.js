const Query = require('../../lib/utils/query');

test("success should return object with passed result and code of 0", () => {
    // Arrange
    const result = {};

    // Act
    const query = Query.success(result);

    // Assert
    expect(query.code).toBe(0);
    expect(query.result).toBe(result);
});