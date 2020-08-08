module.exports = {
    dbParams: function() {
        let result = {}
        result["host"] = process.env.X_MYSQL_HOST
        result["port"] = process.env.X_MYSQL_PORT
        result["user"] = process.env.X_MYSQL_USER
        result["password"] = process.env.X_MYSQL_PASSWORD
        result["name"] = process.env.X_MYSQL_DB_NAME
        return result
    },
    apiDSN: function() {
        const apiADDR = process.env.X_SERVER_LOCAL_HOST
        const apiPORT = process.env.X_SERVER_PORT
        let apiDSN = `${apiADDR}:${apiPORT}`

        if (!apiDSN) {
            console.log('Failed init api host:port. Abort!')
            return undefined
        }

        return apiDSN
    }
}
