'use strict';

const http = require('http');

const server = http.createServer((req, res) => {
    console.log(Date.now(), 'request received');
    res.end('hello world');
});

server.listen(3000);
console.log('listening on 3000');
