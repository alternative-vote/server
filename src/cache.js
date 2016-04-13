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
      offset: 186
    }], {fromOffset: true});

    consumer.on('message', (message) => {
      // console.log('received message:', message);
      try {
        _reduce(this, JSON.parse(message.value))
      } catch(err) {
        console.log(err);
      }
    });

  }
  //
  // _reduce(root, message){
  //   console.log("in the reduce funcion");
  //   const id = message.id;
  //   if(root[id] == null){
  //     root[id] = {};
  //   }
  //
  //   _init(root, id)
  //
  //   root[id][message.key] = message.val;
  //
  //   console.log(root);
  //
  // }
}

const fieldTypes = {
  val: 'value',
  ref: 'reference',
  collectionOfVals : 'collectionOfVals',
  collectionOfRefs : 'collectionOfRefs',
}

const cacheModels = {
  election: {
    name: fieldTypes.val,
    candidates: fieldTypes.collectionOfRefs,
    categories: fieldTypes.collectionOfRefs,
  },
  candidate: {
    name: fieldTypes.val,
    description: fieldTypes.val,
  },
  category: {
    name: fieldTypes.val,
    ballots: fieldTypes.collectionOfRefs,
  },
  ballot: {
    votes : fieldTypes.collectionOfRefs,
  }
}



function _reduce(root, message){
  const type = message.type;
  const id = message.id;
  const key = message.key;
  // const action = message.action;
  // const val = message.val;

  root[id] = _init(root, id)
  const oldVal = root[id][message.key];
  root[id][key] = _reduceModel(oldVal, cacheModels[type], message)

  // root[id][message.key] = message.val;

  // console.log(root);
}

function _init(root, id) {
  if(root[id] == null){
    return {};
  }
  return root[id];

}

//model level
function _reduceModel(oldVal, cacheModel, message) {
  const key = message.key;
  const action = message.action;
  const val = message.val;

  if (cacheModel[key] == fieldTypes.val) {
      return _reduceVal(oldVal, action, val)
  }

  if (cacheModel[key] == fieldTypes.collectionOfRefs) {
      return _reduceCollectionOfRefs(oldVal, action, val)
  }

}


//field values
function _reduceVal(oldVal, action, newVal) {
  if( action === '+' ) {
    return newVal;
  }

  if( action === '-' || action === 'clear') {
    return undefined;
  }
}

function _reduceCollectionOfRefs(oldVal, action, newVal) {
  if(oldVal == null) {
    oldVal = [];
  }

  if( action === '+' ) {
    // console.log('oldVal = ', oldVal);
    // console.log('newVal = ', newVal);
    oldVal.push(newVal);
    return oldVal;
  }

  if( action === '-' ) {
    throw new Error('_reduceCollectionOfRefs() does not support - action yet');
  }

  if( action === 'clear' ) {
    return undefined;
  }
}




module.exports = Cache
