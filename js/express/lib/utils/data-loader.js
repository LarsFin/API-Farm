const fs = require('fs');

exports.load = () => new Promise((resolve, reject) => {
    fs.readFile('./data.json', 'utf8', (err, data) => {
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
