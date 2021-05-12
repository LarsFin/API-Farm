const clone = require('../../lib/utils/clone');

test('clone should return copy of object', () => {
    // Arrange
    const originalObj = {
        k1: {
            k1: 'v',
            k2: [1, 2, 3]
        },
        k2: false,
        k3: [
            { k: 'v' },
            { k1: 'v', k2: 100 }
        ]
    };

    // Act
    const cloneObj = clone(originalObj);

    // Assert
    expect(cloneObj).not.toBe(originalObj);
    expect(cloneObj).toMatchObject(originalObj);
});
