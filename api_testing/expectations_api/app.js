const express = require('express');
const app = express();
const port = 3000;

app.get('/ping', (_, res) => {
  res.send('pong')
});

app.use(express.static('expected_data'));

app.listen(port, () => {
  console.log(`Expectations api listening @ http://localhost:${port}`)
});