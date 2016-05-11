const path = require('path');
const SuperRouter = require('super-router');
const ContentNegotiation = SuperRouter.Middleware.ContentNegotiation;
const SuperSwag = require('super-swag');
const middleware = require('./middleware');
const _ = require('lodash');


module.exports = function(controllers, validators) {

  const app = new SuperRouter.App();
  const superSwag = new SuperSwag({
    swaggerFileLocation: path.resolve(__dirname, '../swagger/swagger.yaml'),
    controllers: controllers,
    customValidators: validators //register custom format validators that aren't part of the swagger spec
  });


  app.then(ContentNegotiation.request);
  // app.then(middleware.authorize); //TODO: implement this in super-swag
  app.then(superSwag.meta);
  app.then(superSwag.validate);
  app.then(superSwag.route);
  app.catch(middleware.error);
  app.then(ContentNegotiation.response);

  return app;
}
