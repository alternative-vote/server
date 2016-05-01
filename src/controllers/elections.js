'use strict';
const uuid = require('uuid');

module.exports = {
  upsert: upsert,
  read: read
};

const cache = [];

function upsert(opts) {

  const request = opts.request;
  const response = opts.response;

  response.setBody({hey: "you got routed to the right place!"})
}

function read(req, res) {

}
