'use strict';

const kafka = require('kafka-node');
const EventEmitter = require('events');
const _ = require('lodash');
const fs = require('fs');
const util = require('util');

const TOPIC = 'events'

const DB_FILE = './DB.txt';
const CACHE_FILE = './cache.txt';


function ensureLength(s, length) {
  if(s.length >= length){
    return s;
  }

  return ensureLength(s+' ', length);
}


class Cache extends EventEmitter {
  constructor(host) {
    super();
    // fs.writeFile(DB_FILE, "offset\tid\t\ttype\tkey\taction\tval\n");
    // fs.writeFile(CACHE_FILE, "");


    const client   = new kafka.Client(host);
    const consumer = new kafka.Consumer(client, [{
      topic : TOPIC,
      offset: 6017
    }], {fromOffset: true});

    consumer.on('message', (message) => {
      // console.log('received message:', message);

      try {
        const parsed = JSON.parse(message.value);
        // fs.appendFileSync(DB_FILE, `${message.offset}\t${ensureLength(parsed.id, 55)}\t${ensureLength(parsed.type, 15)}\t${ensureLength(parsed.key, 15)}\t${ensureLength(parsed.action, 5)}\t${parsed.val}\n`);

        _reduce(this, parsed)
        // fs.writeFile(CACHE_FILE, util.inspect(this,{depth: null }))
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
    username: fieldTypes.val,
    votes: fieldTypes.collectionOfRefs,
  }
}



function _reduce(root, message){
  const type = message.type;
  const id = message.id;
  const key = message.key;

  root[id] = _init(root, id)
  const oldVal = root[id][message.key];
  root[id][key] = _reduceModel(oldVal, cacheModels[type], message)

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
    oldVal.push(newVal);
    return oldVal;
  }

  if( action === '-' ) {
    return _.without(oldVal, newVal);
  }

  if( action === 'clear' ) {
    return [];
  }
}




module.exports = Cache
