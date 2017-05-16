"use strict";

const express = require('express');
const cors = require('cors');
const morgan = require('morgan');
const zips = require('./zips.json');

const zipCityIndex = zips.reduce((index, record) => {
    let cityLower = record.city.toLowerCase()
    let zipsForCity = index[cityLower]
    if (!zipsForCity) {
        index[cityLower] = zipCityIndex = []
    }
    zipsForCity.push(record);
    return index
}, {})


const app = express();

const port = process.env.PORT || 80;
const host = process.env.HOST || '';

app.use(morgan('dev'));
app.use(cors());

app.get('/zips/city/:cityName', (req, res) => {
    let zipsForCity = zipsForCity[req.params.cityName.toLowerCase()];
    if (!zipsForCity) {
        res.status(404).send('invalid city name')
    } else {
        res.json(zipsForCity)
    }
})

app.get('/hello/:name', (req, res) => {
    res.send(`hello ${req.params.name}!`);
});

app.get()

app.listen(port, host, () => {
    console.log(`server is listening at http://${host}:${port}`)
});