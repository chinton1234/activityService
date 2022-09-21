const express = require("express");
const app = express();
const mongoose = require('mongoose');

mongoose.connect("mongodb+srv://activity_service:123456B@activity.lglfdxc.mongodb.net/?retryWrites=true&w=majority");

const port = process.env.PORT || 3000;

const db = mongoose.connection;
db.on('error', (error) => console.error(error));
db.once('open', () => console.log('Connected to Database'));

pp.use(express.json());

const activityRouter = require('./schema/activity');
app.use('/activity', activityRouter);

app.listen(port, () => {
  console.log("Starting node.js at port " + port);
});