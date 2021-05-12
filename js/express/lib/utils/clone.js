const clone = (originalObj) => {
    const cloneObj = (originalObj instanceof Array ? [] : {});

    for (const [k, v] of Object.entries(originalObj))
        cloneObj[k] = (v instanceof Object ? clone(v) : v);

    return cloneObj;
};

module.exports = clone;
