const path = require('path');
const SuperRouter = require('super-router');
const ContentNegotiation = SuperRouter.Middleware.ContentNegotiation;
const superSwag = require('super-swag');
const middleware = require('./middleware');


const app = new SuperRouter.App();
const superSwagOptions = {
  swaggerFileLocation: path.resolve(__dirname, '../swagger.yaml'),
  controllerDir: path.resolve(__dirname, 'controllers')
}
// call the superSwag constructor (that doesn't exist yet), and do a fs.readSync to load the yaml file.
// also pass in everything any middleware would need so you don't have to pass in options for each middleware you set up.


//the only thing you should go to the file system for is the yaml file.
//have the router take a hash like this (express train does this if you group stuff into controllers):

var mockRouting = {
  elections:{
    read: function (opts) {},
    upsert: function (opts) {},
  },
  users: {
    read: function (opts) {}
  }
};
app.then(ContentNegotiation.request);
// app.then(middleware.authorize(api));
app.then(superSwag.validate(superSwagOptions));
app.then(superSwag.route(superSwagOptions));
app.catch(middleware.error);
app.then(ContentNegotiation.response);


module.exports = app;
