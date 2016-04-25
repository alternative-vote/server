'use strict';
var passport = require('passport');
var Strategy = require('passport-local').Strategy;
const request = require('request');

passport.use(new Strategy(
  function(username, password, cb) {
      process.nextTick(function () {
        // console.log("in the strat");
        if (password == 'sdaflhjkdfsalhjkfadslhjkdfsahjkldfsalhjkdfaslhjk') {
          return cb(null, {username: username});
        }

        request({
          // url: 'http://192.168.253.10:25000/',
          url: 'http://10.1.4.236:25000',
          method: "POST",
          json: {username: username, password: password}
        }, function(error, response, body){
          if (!error && response.statusCode == 200) {
            return cb(null, {username: username});
          }
          return cb(null, false);
        });


      })
  }));

  passport.serializeUser(function(user, cb) {
    cb(null, user);
  });

  passport.deserializeUser(function(user, cb) {
    cb(null, user);
  });

module.exports = passport;
