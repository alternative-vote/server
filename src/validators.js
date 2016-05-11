const validator = require('validator');


module.exports = function() {
  return {
    "uuid": validator.isUUID,
    "test": (val)=>val==="test",
  }
}
