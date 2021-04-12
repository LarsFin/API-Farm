const express = require('express');
const app = express();
const port = 3000;

app.get('/ping', (_, res) => {
  res.send('pong')
});

app.use(express.static('expected_data'));

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
});