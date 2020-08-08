let general = require('./db/general')

exports.mochaHooks = {
    afterAll(done) {
        general.end()
        done()
    },
};
