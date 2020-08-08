let params = require('../params/params')
let mysql = require('mysql')
let con;

_connection();

function _connection() {
    let dbParams = params.dbParams()
    console.log("PARAMS: ", dbParams)
    console.log("host: ", dbParams.host)
    con = mysql.createConnection({
        host: dbParams.host,
        port: dbParams.port,
        user: dbParams.user,
        password: dbParams.password,
        database: dbParams.name
    });
    con.connect(function(err) {
        if (err) {
            console.log("FAILED CONNECT TO DB. ERR: ", err)
            throw err;
        }
        console.log("Connected!");
    });
}

function _clear_table(table, done) {
    // Welcome to callback's hell
    con.query("SET FOREIGN_KEY_CHECKS = 0", err => {
        if (err) throw err;
        con.query("TRUNCATE TABLE " + table, err => {
            if (err) throw err;
            con.query("SET FOREIGN_KEY_CHECKS = 1", err => {
                if (err) throw err;
                console.log("\t...truncate table: " + table);
                if (done) {
                    done();
                }
            });
        });
    });
}

module.exports = {
    connection: function() {
        return con;
    },
    clear_all_tables: function(done) {
        _clear_table("publisher_app", () => {
            _clear_table("banner_localized_content", () => {
                _clear_table("banner_group", () => {
                    _clear_table("banner", () => {
                        _clear_table("game_localized_content", () => {
                            _clear_table("game", () => {
                                _clear_table("project", done);
                            });
                        });
                    });
                });
            });
        });
    },
    clear_table: (table, done) => {
        _clear_table(table, done)
    },
    connect: function(done) {
        con.connect(function(err) {
            if (err) throw err;
            if (done) done();
        });
    },
    end: function() {
        con.end();
    }
}
