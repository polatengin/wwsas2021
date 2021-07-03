const express = require('express');
const app = express();

app.get('/', function (req, res) {
  res.write("Hello world!")
});

app.get('/get-current', function (req, res) {
  res.write({
    title: "Great Campaign",
    pictureUrl: "https://picsum.photos/1920/1080",
    gotoUrl: "https://devopstips.net"
  })
});

const server = app.listen(7000, function () {
  const host = server.address().address
  const port = server.address().port
  console.log("Example app listening at http://%s:%s", host, port)
});
