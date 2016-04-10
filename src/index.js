'use strict';

const http  = require('http');
const url   = require('url');
const kafka = require('kafka-node');

const TOPIC_NAME = 'queries';

const client   = new kafka.Client('zookeeper:2181');
const producer = new kafka.Producer(client);

const messages = [];

producer.on('error', (err) => {
  console.error('producer error:', err);
});

producer.on('ready', () => {
  console.log('producer ready');
  producer.createTopics([TOPIC_NAME], false, (err, data) => {
    console.error(err);
    console.log(data);
  });

  const consumer = new kafka.Consumer(client, [{
    topic : TOPIC_NAME
  }]);

  consumer.on('message', (message) => {
    console.log('received message:', message);
    try {
      messages.push(JSON.parse(message.value));
    } catch(err) {}
  });

  const server = http.createServer((req, res) => {
    const query = url.parse(req.url, true).query;
    if(query && Object.keys(query).length > 0) {
      const message = JSON.stringify(query);
      producer.send([{
        topic    : TOPIC_NAME,
        messages : [message]
      }], (err, data) => {
        console.log(`wrote message: ${message}`)
        console.error('err', err);
        console.log('data', data);
      });
    }
    res.writeHead(200, {"Content-Type": "application/json"});
    res.end(JSON.stringify(messages));
  });

  server.listen(3000);
  console.log('listening on 3000');

});






