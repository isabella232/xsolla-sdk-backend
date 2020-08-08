module.exports = {
    select: function (sql, values, callback) {
        return _select(sql, values, callback);
    },
    insert: function (sql, values, callback) {
        return _insert(sql, values, callback);
    },
};

var tools = require('./general');
var con = tools.connection();

function _select(sql, values, callback) {
    con.query(sql, values, function (err, result) {
        if (err) throw err;
        callback(result)
    });
}

function _insert(sql, values, callback) {
    con.query(sql, values, function (err, result) {
        if (err) throw err;
        callback(err, result)
    });
}
