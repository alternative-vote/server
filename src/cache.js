'use strict';

const kafka = require('kafka-node');
const EventEmitter = require('events');

const TOPIC = 'events'

class Cache extends EventEmitter {
  constructor(host) {
    super();

    const client   = new kafka.Client(host);
    const consumer = new kafka.Consumer(client, [{
      topic : TOPIC,
      offset: 0
    }], {fromOffset: true});

    consumer.on('message', (message) => {
      console.log('received message:', message);
      try {
        _reduce(this, JSON.parse(message.value))
      } catch(err) {

        console.log(err);
      }
    });

  }

  _reduce(root, message){
    console.log("in the reduce funcion");
    const id = message.id;
    if(root[id] == null){
      root[id] = {};
    }

    _init(root, id)

    root[id][message.key] = message.val;

    console.log(root);

  }
}

function _reduce(root, message){
  console.log("in the reduce funcion");
  const type = message.type;
  const id = message.id;
  const key = message.key;
  const action = message.action;
  const val = message.val;


  _init(root, id)

  root[id][message.key] = message.val;

  console.log(root);

}

function _init(root, id) {
  if(root[id] == null){
    root[id] = {};
  }
}

module.exports = Cache
