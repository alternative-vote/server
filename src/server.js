'use strict';

var express = require('express');
var passport = require('./authentication');
const uuid = require('uuid');
const Producer = require('./producer');
const Cache = require('./cache');
const TYPES = require('./TYPES');
const _ = require('lodash');


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


app.post('/election', needsAuth, (req, res) => {
  const id = uuid.v4();
  _.forEach(req.body, (v, k) => {
    producer.write(TYPES.election, id, k, '+', v)
  })
  res.status(201).json({id: id});
})

app.post('/elections/:id/candidates', needsAuth, (req, res) => {
  const electionId = req.params.id
  const id = uuid.v4();
  _.forEach(req.body, (v, k) => {
    producer.write(TYPES.election, id, k, '+', v)
  })
  res.status(201).json({id: id});
})



function needsAuth(req, res, next) {
  if (true || req.isAuthenticated()) {
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
