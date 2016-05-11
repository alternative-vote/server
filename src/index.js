'use strict';

const train    = require('express-train');

module.exports = () => {
  const tree = train({
    base   : __dirname,
    config : false,
    files  : [
      './**/*.js',
      '../package.json',
      { pattern : './**/*Controller.js', aggregateOn : 'controllers' },
    ],
  });

  return tree;

};
