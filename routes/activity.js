const express = require('express')
const router = express.Router()
const Activity = require('../schema/activity')

// Getting all
router.get('/', async (req, res) => {
    try {
      const activity = await Activity.find()
      res.json(activity)
    } catch (err) {
      res.status(500).json({ message: err.message })
    }
  })

  router.get('/:id', get, (req, res) => {
    res.json(res.subscriber)
  })
  






module.exports = router