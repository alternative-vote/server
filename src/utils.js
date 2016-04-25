'use strict';
const altVote = require('./altVote')
const _ = require('lodash');


const KEY = 'name'

class utils {
  static getResults(election){
    // console.log("get results called");

    //set up simple collections with just the info we need to calculate results
    let candidates = _.map(election.candidates, candidate => candidate[KEY]);
    const simpleBallotsByCategory = {};
    _.forEach(election.categories, (category)=>{
      const simpleBallots = _.map(category.ballots, ballot => _.map(ballot.votes, vote => vote[KEY]));
      if(simpleBallots.length > 0) {
        simpleBallotsByCategory[category[KEY]] = simpleBallots;
      }
    });


    const finalResults = {}
    while(_.keys(simpleBallotsByCategory).length > 0) {
      const results = {};
      // console.log("starting loop");
      // console.log("candidates:");
      // console.log(candidates);
      // console.log('simpleBallotsByCategory:');
      // console.log(simpleBallotsByCategory);

      //go through each category and calculate the winner
      _.forEach(simpleBallotsByCategory, (simpleBallots, categoryName) => {
        console.log("calculating results of category: ", categoryName, "with candidates:");
        console.log(candidates);
        const categoryResults = altVote(candidates, simpleBallots);

        //reduce for easy to read output
        _.forEach(categoryResults, (round) => {
            _.forEach(round, (ballots, candidate) => {
                round[candidate] = ballots.length;
            });
          });

        results[categoryName] = categoryResults;
      })

      const categoryWinners = [];
      _.forEach(results, (rounds, category)=>{
        const lastRound = _.last(rounds);
        var winner = _.maxBy(_.keys(lastRound), o => lastRound[o]);
        categoryWinners.push({
          category: category,
          winner: winner,
          winningVoteCount: lastRound[winner],
        });
      });

      if(categoryWinners.length > 0) { //TODO: REMOVE THIS CHECK?
        //find the category that was won the hardest and save that off
        // console.log(categoryWinners);
        const categoryThatWasWonTheHardest = _.maxBy(categoryWinners, o => o.winningVoteCount); //picks the highest index if there's a tie, so that's cool
        console.log(`${categoryThatWasWonTheHardest.category} was won the hardest by ${categoryThatWasWonTheHardest.winner} with ${categoryThatWasWonTheHardest.winningVoteCount} votes`);
        finalResults[categoryThatWasWonTheHardest.category] = _.cloneDeep(results[categoryThatWasWonTheHardest.category]);
        //remove that category
        delete simpleBallotsByCategory[categoryThatWasWonTheHardest.category];
        //as well as that candidate
        candidates = _.without(candidates, categoryThatWasWonTheHardest.winner);
        //redo calculations on remaining categories and candidates

      }


   }

    return finalResults;




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
