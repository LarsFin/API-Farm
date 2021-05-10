const ApiFarmDate = require('../../lib/utils/api-farm-date');

test('isValid should return true when date in format \'DD/MM/YYYY\'', () => {
    // Arrange
    const strDate = '22/01/2021';

    // Act
    const isValid = ApiFarmDate.isValid(strDate);

    // Assert
    expect(isValid).toBe(true);
});

test('isValid should return false when date in format \'MM/DD/YYYY\'', () => {
    // Arrange
    const strDate = '01/22/2021';

    // Act
    const isValid = ApiFarmDate.isValid(strDate);

    // Assert
    expect(isValid).toBe(false);
});

test('isValid should return false when date in format \'DD-MM-YYYY\'', () => {
    // Arrange
    const strDate = '22-01-2021';

    // Act
    const isValid = ApiFarmDate.isValid(strDate);

    // Assert
    expect(isValid).toBe(false);
});
