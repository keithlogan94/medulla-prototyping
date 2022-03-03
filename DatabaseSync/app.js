//const express = require('express')
//var bodyParser = require('body-parser')

//const app = express()

//app.use(bodyParser.json());       // to support JSON-encoded bodies
//app.use(bodyParser.urlencoded({     // to support URL-encoded bodies
//    extended: true
//}));

const { Sequelize, Model, DataTypes } = require('sequelize');

async function syncDatabase(data) {
    data = {
        data: '{"Databases":[{"Name":"test","Host":"localhost","Dialect":"mysql","UsernameSecretName":"root","PasswordSecretName":"test","Models":[{"Name":"User","Columns":[{"ColumnName":"firstName","Type":"VARCHAR(100)","AllowNull":false,"DefaultValue":"1","PrimaryKey":true,"Field":"test_field","Unique":"uniqueField","Comment":"This is a Test","Validate":null}]},{"Name":"Account","Columns":[{"ColumnName":"firstName","Type":"VARCHAR(100)","AllowNull":false,"DefaultValue":"1","PrimaryKey":true,"Field":"test_field","Unique":"uniqueField","Comment":"This is a Test","Validate":null},{"ColumnName":"lastName","Type":"VARCHAR(100)","AllowNull":false,"DefaultValue":"1","AutoIncrement":true,"PrimaryKey":true,"Field":"test_field","Unique":"uniqueField","Comment":"This is a Test","Validate":null}]}]}],"Metadata":{"annotations":null,"clusterName":null,"creationTimestamp":"2022-03-02T23:34:01Z","deletionGracePeriodSeconds":null,"deletionTimestamp":null,"finalizers":null,"generateName":null,"generation":1,"labels":null,"managedFields":[{"apiVersion":"medulla.recro.com/v1alpha1","fieldsType":"FieldsV1","fieldsV1":{"f:databases":{}},"manager":"kubectl-create","operation":"Update","subresource":null,"time":"2022-03-02T23:34:01Z"}],"name":"test-1","namespace":"default","ownerReferences":null,"resourceVersion":"472790","selfLink":null,"uid":"777666cd-e799-42be-846e-1ef00c858a7f"},"apiVersion":"medulla.recro.com/v1alpha1","kind":"Data"}'
    };
    data = JSON.parse(data.data)
    console.log(data);

    const databases = data.Databases;

    for (let i = 0; i < databases.length; i++) {
        const { Name, Host: host, Dialect: dialect, UsernameSecretName, PasswordSecretName, Models } = databases[i];
        console.log(Name, host, dialect, UsernameSecretName, PasswordSecretName)
        const sequelize = new Sequelize(Name, UsernameSecretName, PasswordSecretName, {
            host,
            dialect,
            port: 3306
        });

        try {
            await sequelize.authenticate();
            console.log('Connection has been established successfully.');
        } catch (error) {
            console.error('Unable to connect to the database:', error);
        }


        for (let modelIndex = 0; modelIndex < Models.length; modelIndex++) {
            console.log(Models[modelIndex]);

            let initObject = { firstName: "" };
            for (let columnIndex = 0; columnIndex < Models[modelIndex].Columns.length; columnIndex++) {
                initObject[Models[modelIndex].Columns[columnIndex].ColumnName]
                    = Models[modelIndex].Columns[columnIndex];
            }



            let keys = Object.keys(initObject);
            for (let keysIndex = 0; keysIndex < keys.length; keysIndex++) {

                let subKeys = Object.keys(initObject[keys[keysIndex]])

                for (let subKeysIndex = 0; subKeysIndex < subKeys.length; subKeysIndex++) {


                    const toLowerSubProperty = subKeys[subKeysIndex][0].toLowerCase() + subKeys[subKeysIndex].slice(1)
                    const deleteSubProperty = subKeys[subKeysIndex][0].toUpperCase() + subKeys[subKeysIndex].slice(1)

                    initObject[keys[keysIndex]][(toLowerSubProperty)]
                        = initObject[keys[keysIndex]][subKeys[subKeysIndex]]

                    delete initObject[keys[keysIndex]][(deleteSubProperty)]

                }

                
            }

            console.log("initializing table", Models[modelIndex].Name)
            console.log({ initObject })


            const Model = sequelize.define(Models[modelIndex].Name, initObject, {
                sequelize, // We need to pass the connection instance
                modelName: Models[modelIndex].Name // We need to choose the model name
                // Other model options go here
            });

            await Model.sync({ force: true })



        }





    }

    



}

(async () => {
    await syncDatabase()
})()


//app.post('/listen-for-database-schema', function (req, res) {
//    console.log(req.body);
//    res.send('Got database schema')
//})

//app.listen(3000)