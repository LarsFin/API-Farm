const Clone = require('../../lib/utils/clone');

test('clone should return copy of object', () => {
    // Arrange
    const original = {
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
    const clone = Clone.object(original);

    // Assert
    expect(clone).not.toBe(original);
    expect(clone.k1).not.toBe(original.k1);
    expect(clone.k3).not.toBe(original.k3);

    expect(clone).toMatchObject(original);
});
