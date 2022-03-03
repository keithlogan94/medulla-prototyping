const express = require('express')
var bodyParser = require('body-parser')

const app = express()

app.use(bodyParser.json());       // to support JSON-encoded bodies
app.use(bodyParser.urlencoded({     // to support URL-encoded bodies
    extended: true
}));

app.post('/listen-for-database-schema', function (req, res) {
    console.log(req.body);
    res.send('Got database schema')
})

app.listen(3000)