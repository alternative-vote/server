'use strict';
const uuid = require('uuid');

module.exports = function () {

  const cache = {};

  class ElectionController {

    static upsert(opts){
      const request = opts.request;
      const response = opts.response;
      const election = request.body;

      if (election.id == null){
        election.id = uuid.v4();
      }

      cache[election.id] = election;

      response.setBody(election)
    }

    static read(opts){
      const request = opts.request;
      const response = opts.response;
      const id = request.routeParams.id;

      if (cache[id] == null){
        throw { statusCode: 404, message: `could not find an election with id ${id}` };
      }

      response.setBody(cache[id]);
    }
  }

  return ElectionController;
}
