const chai = require('chai');
chai.use(require('chai-shallow-deep-equal'));

const params = require('../../../params/params'),
    expect = require('chai').expect,
    supertest = require('supertest'),
    api = supertest(params.apiDSN());

describe('Get user with positive code - 200', function () {

    it("Get user with correct access_token", function (done) {
        api.post("/login")
            .set("Accept", "application/json")
            .send({
                "email": "dev@dev.dev"
            })
            .expect(200)
            .then(res => {
                // Need setup mockup
                // expect(res.body).to.shallowDeepEqual(
                //     {
                //         "access_token": "f4puO8lf2ZC19g4Sr5fsxccgb5EcEQ64"
                //     }
                // )
                done()
            }).catch(err => done(err))
    })
});