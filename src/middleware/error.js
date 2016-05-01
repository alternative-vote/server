const _ = require('lodash');
const JD = require('junk-drawer');

module.exports = {
  handler: (opts) => {
      const error = opts.error;
      const response = opts.response;
      const logger = console;

      if (error.statusCode == null) {
        error.details = error.message;
        error.statusCode = 500;
        error.name = 'InternalServerError';
        error.message = 'Internal Server Error';
      }

      if (error.stack != null || error.details != null || JD.is500(error.statusCode)) {
        logger.error(error.stack || error.details);
      }

      if (error.upstreamError != null) {
        logger.error(`UpstreamError: `, error.upstreamError);
      }

      // get the message out of the error object
      const message = error.message;
      const plain = _.toPlainObject(error);
      delete plain.message;
      plain.message = message;
      if (!_.isString(plain.details)) {
         // This is to prevent circular references if this is an error object
         // if it was an error, it already got logged.
        delete plain.details;
      }

      response.statusCode = error.statusCode;
      response.setBody(plain);
  },
  errorHandler : (opts) => {
    const error = opts.error;
    const response = opts.response;

    console.error('Something very unexpected happened in the ErrorStack middleware: ', error.stack);

    const simplestErrorPossible = {
      statusCode : 500,
      name       : 'InternalServerError',
      message    : 'Internal Server Error'
    };
    response.setBody(JSON.stringify(simplestErrorPossible));
    response.end();
  }
};
