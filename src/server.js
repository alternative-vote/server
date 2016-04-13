'use strict';

var express = require('express');
var passport = require('./authentication');
const uuid = require('uuid');
const Producer = require('./producer');
const Cache = require('./cache');
const TYPES = require('./TYPES');
const _ = require('lodash');
const utils = require('./utils')


const ZOOKEEPER_HOST =  '192.168.99.100:2181';
let producer, cache;

// Create a new Express application.
var app = express();

// Use application-level middleware for common functionality, including
// logging, parsing, and session handling.
// app.use(require('morgan')('combined'));
app.use(require('cookie-parser')());
app.use(require('body-parser').urlencoded({ extended: true }));
app.use(require('body-parser').json());
app.use(require('express-session')({ secret: 'boggle at the situation', resave: false, saveUninitialized: false }));
// Initialize Passport and restore authentication state, if any, from the
// session.
app.use(passport.initialize());
app.use(passport.session());


// Define routes.
app.get('/',
  function(req, res) {
    res.send('root route');
  });

app.post('/login',
  passport.authenticate('local', { failureRedirect: '/'}),
  function (req, res) {
    res.redirect('/p')
  }
  );

app.get('/logout',
  function(req, res){
    req.logout();
    res.redirect('/');
  });

app.get('/p',
  needsAuth,
  function(req, res){
    res.send("in the private route");
  });

  app.get('/elections/:id', needsAuth, (req, res) => {
    const id = req.params.id;
    const election = _.cloneDeep(cache[id]);
    // console.log(cache);

    if (election == null) {
      return res.sendStatus(404);
    }

    election.id = id;

    election.candidates = denormalizeCollection(election.candidates);
    election.categories = denormalizeCollection(election.categories);

    _.forEach(election.categories, (category) => {
      category.ballots = denormalizeCollection(category.ballots);
      _.forEach(category.ballots, (ballot) => {
        ballot.votes = denormalizeCollection(ballot.votes);
      });
    });

    election.results = utils.getResults(election);

    res.json(election);
  });

app.post('/elections', needsAuth, (req, res) => {
  const id = uuid.v4();
  _.forEach(req.body, (v, k) => {
    producer.write(TYPES.election, id, k, '+', v)
  })
  res.status(201).json({id: id});
})

app.post('/elections/:id/candidates', needsAuth, (req, res) => {
  const electionId = req.params.id;
  const candidateId = uuid.v4();

  _.forEach(req.body, (v, k) => {
    producer.write(TYPES.candidate, candidateId, k, '+', v)
  })

  producer.write(TYPES.election, electionId, 'candidates', '+', candidateId)

  res.status(201).json({id: candidateId});
});

app.post('/elections/:id/categories', needsAuth, (req, res) => {
  const electionId = req.params.id;
  const categoryId = uuid.v4();

  _.forEach(req.body, (v, k) => {
    producer.write(TYPES.category, categoryId, k, '+', v)
  })

  producer.write(TYPES.election, electionId, 'categories', '+', categoryId)

  res.status(201).json({id: categoryId});
});

app.post('/elections/:electionId/categories/:categoryId/ballots', needsAuth, (req, res) => {
  const electionId = req.params.electionId; //TODO: this isn't needed
  const categoryId = req.params.categoryId;
  const ballotId = uuid.v4();

  console.log('req.params = ', req.params);
  console.log('req.body = ', req.body);

  _.forEach(req.body, (v, k) => {
    if( k == 'votes') {
      _.forEach(v, (vote) => {
         producer.write(TYPES.ballot, ballotId, k, '+', vote)
      });
    }
    else {
      producer.write(TYPES.ballot, ballotId, k, '+', v)
    }

  })

  producer.write(TYPES.category, categoryId, 'ballots', '+', ballotId)

  res.status(201).json({id: ballotId});
});





function needsAuth(req, res, next) {
  if (true || req.isAuthenticated()) { // TODO: don't short circut auth
    return next();
  }
  res.sendStatus(401);
}

producer = new Producer(ZOOKEEPER_HOST);
producer.on('ready', () => {
  cache = new Cache(ZOOKEEPER_HOST);
  app.listen(3000);
  console.log('server listening on port 3000');
})


function denormalizeCollection(collectionOfRefs) {
  const ret = [];

  _.forEach(collectionOfRefs, (ref, i) => {
    if(cache[ref] != null) {
      const obj = _.cloneDeep(cache[ref])
      obj.id = ref;
      ret.push(obj);
    }
  });

  return ret;
}
