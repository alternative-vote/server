'use strict';

const http  = require('http');
const url   = require('url');
const kafka = require('kafka-node');

const TOPIC_NAME = 'queries';

const client   = new kafka.Client('zookeeper:2181');
const producer = new kafka.Producer(client);


producer.on('ready', () => {
  producer.createTopics([TOPIC_NAME], false, (err, data) => {
    console.error(err);
    console.log(data);
  });

  const consumer = new kafka.Consumer(client, [{
    topic : TOPIC_NAME
  }]);

  const messages = [];
  client.on('message', (message) => {
    messages.push(JSON.parse(message));
  });

  const server = http.createServer((req, res) => {
    const query = url.parse(req.url).query;
    producer.send({
      topic    : TOPIC_NAME,
      messages : [JSON.stringify(query)]
    });
    res.end(messages);
  });

  server.listen(3000);
  console.log('listening on 3000');

});

function startServer() {

}





