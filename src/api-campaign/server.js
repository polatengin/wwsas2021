const express = require('express');
const app = express();

app.get('/', function (req, res) {
  res.write("Hello world!")
});

const server = app.listen(7000, function () {
});
