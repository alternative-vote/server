'use strict';
const altVote = require('./altVote')
const _ = require('lodash');


const KEY = 'name'

class utils {
  static getResults(election){
    // console.log("get results called");

    //get array of candidates
    const candidates = _.map(election.candidates, candidate => candidate[KEY]);
    const simpleBallotsByCategory = {};
    _.forEach(election.categories, (category)=>{
      simpleBallotsByCategory[category.name] = _.map(category.ballots, ballot => _.map(ballot.votes, vote => vote[KEY]));
    });

    return 1;

    return _calculate(candidates, simpleBallotsByCategory);

    // const results = {}
    // _.forEach(election.categories, (category)=>{
    //   // console.log("working with category = ", category);
    //
    //   //get array of simple ballots
    //   const simpleBallots = _.map(category.ballots, ballot => _.map(ballot.votes, vote => vote[KEY]));
    //
    //   //crunch the numbers
    //   const categoryResults = altVote(candidates, simpleBallots);
    //
    //   reduce
    //   _.forEach(categoryResults, (round) => {
    //     _.forEach(round, (ballots, candidate) => {
    //         round[candidate] = ballots.length;
    //     });
    //   });
    //
    //   results[category.name] = categoryResults;
    // });
    //
    //



    return results
  }

}


// [
//   category: 'Best Hair',
//   winner : 'A',
//   ballots: []
// }

function _calculate(candidates, simpleBallotsByCategory, results) {
  _.forEach(simpleBallotsByCategory, (bins, category) => {

  });

  if (_uniqueWinners(results)){
    return results;
  }
  return results
}

function _uniqueWinners(results) {
  return (_.uniqBy(results, 'winner')).length == results.length;
}



function _getWinner(results) {
  const lastRound = _.last(results);
  let winner, winnerVoteCount;
  _.forEach(lastRound, (v, k) => {
    if (!winner || v.length > winnerVoteCount ) {
      winner = k;
      winnerVoteCount = v.length;
    }
  });
  return {name: winner, voteCount: winnerVoteCount};
}

module.exports = utils;
