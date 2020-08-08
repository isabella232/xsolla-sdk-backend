const chai = require('chai');
chai.use(require('chai-shallow-deep-equal'));

const params = require('../../../params/params'),
    supertest = require('supertest'),
    api = supertest(params.apiDSN());

describe('Get healthcheck with positive code - 200', function () {

    it("Get healthcheck", function (done) {
        api.get("/healthcheck")
            .set("Accept", "application/json")
            .expect(200)
            .then(res => done())
            .catch(err => done(err))
    })
});