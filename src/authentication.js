var passport = require('passport');
var Strategy = require('passport-local').Strategy;

passport.use(new Strategy(
  function(username, password, cb) {
      process.nextTick(function () {
        if(username == 'test' && password == 'test') {
          return cb(null, {id: '123', username: 'test'});
        }

        if(username == 'test2' && password == 'test') {
          return cb(null, {id: '456', username: 'test2'});
        }

        if(username == 'test3' && password == 'test') {
          return cb(null, {id: '789', username: 'test3'});
        }

        return cb(null, false);
      })
  }));

  passport.serializeUser(function(user, cb) {
    cb(null, user);
  });

  passport.deserializeUser(function(user, cb) {
    cb(null, user);
  });

module.exports = passport;
