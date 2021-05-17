const cloneObject = original => {
    const clone = (original instanceof Array ? [] : {});

    for (const [k, v] of Object.entries(original))
        clone[k] = (v instanceof Object ? cloneObject(v) : v);

    return clone;
};

exports.object = cloneObject;
