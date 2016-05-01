'use strict';

const http = require('http');
const _ = require('lodash');
const app = require('./app');

// Setup a transport binding to http server
const server = http.createServer((req, res) => {
  const request = ({
    path    : req.url,
    method  : req.method,
    headers : req.headers,
    body    : req
  });

  app.processRequest(request).then((response) => {
    res.statusCode = response.statusCode;
    _.each(response.headers, (value, key) => {
      res.setHeader(key, value);
    });
    response.body.pipe(res);
  }).catch((err) => {
    res.statusCode = 500;
    console.log("Error made it out of super-router");
    console.log(err.stack);
    res.end(err.message);
  }).done();
});

server.listen(3000);
console.log('listening on port 3000');
