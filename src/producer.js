'use strict';

const kafka = require('kafka-node');
const EventEmitter = require('events');

const TOPIC = 'events'

class Producer extends EventEmitter {
  constructor(host) {
    super();
    const client = new kafka.Client(host);
    this.producer = new kafka.Producer(client);

    this.producer.on('ready', () =>{
      this.producer.createTopics(['events'], false, (err, data) => {
        // console.log('ready in producer.js');
        this.emit('ready');
      });
    });
  }

  write(type, id, key, action, val) {
    const message = {
      type: type,
      id: id,
      key: key,
      action: action,
      val: val
    }
    // console.log("writing message:");
    // console.log(message);
    this.producer.send([{
      topic    : TOPIC,
      messages : [JSON.stringify(message)]
    }], (err, data) => {
      // console.log(`wrote message: ${message}`)
      // console.error('err', err);
      // console.log('data', data);
    });
  }
}


module.exports = Producer;
