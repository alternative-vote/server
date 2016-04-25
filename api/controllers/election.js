'use strict';
const uuid = require()

module.exports = {
  upsert: upsert,
  read: read
};

const cache = []''

function upsert(req, res) {
  const id = uuid.v4();
  const election = {
    id: id,
    name: "hardcoded name",
    candidates: [],
    categories: []
  }

  res.json();
}
