const express = require('express')
const app = express()

app.get('/listen-for-database-schema', function (req, res) {
    res.send('Got database schema')
})

app.listen(3000)