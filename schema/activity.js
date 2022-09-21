const mongoose = require('mongoose')

const activitySchema = new mongoose.Schema({
  activityName: {
    type: String,
    required: true
  },
  imageProfile: {
    type: String,
    required: false,
    default: ""
  },
  activityType: {
    type: Array,
    required: true,
    default: []
  },
  activityType: {
    type: Array,
    required: true,
    default: []
  },
  ownerId: {
    type: String,
    required: true
  },
  location: {
    type: String,
    required: true
  },
  participant: {
    type: Array,
    required: true
  },
  chatId: {
    type: String,
    required: true
  }
})

module.exports = mongoose.model('activity', activitySchema)