// date should always be in format DD/MM/YYYY
exports.isValid = date => {
    const dateElements = date.split('/');

    if (dateElements.length !== 3)
        return false;

    return !isNaN(new Date(`${dateElements[1]}/${dateElements[0]}/${dateElements[2]}`));
}
