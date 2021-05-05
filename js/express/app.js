const express = require('express')
const app = express()
const port = 8080

app.get('/ping', (req, res) => {
  res.send('pong')
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})
