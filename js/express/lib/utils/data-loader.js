const fs = require('fs');

exports.load = dataSource => new Promise((resolve, reject) => {
    fs.readFile(dataSource, 'utf8', (err, data) => {
        if (err)
            reject(err);
        else
            try {
                const jsonData = JSON.parse(data);
                resolve(jsonData);
            } catch (err) {
                reject(err);
            }
    });
});
