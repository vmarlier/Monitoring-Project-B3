'use strict'

require('dotenv').config({ path: process.env.DOTENV || '.env' })

const User = require('../models/User')

const mongoose = require('mongoose')
const MONGO_URI = process.env.MONGO_URI || 'mongodb://0.0.0.0:27017/monitoring'
mongoose.connect(MONGO_URI)

const addAdmin = async () => {
  await User.create({
    email: 'admin@admin.com',
    username: 'admin',
    firstName: 'admin',
    lastName: 'admin',
    password: 'admin'
  })

  mongoose.disconnect()
  process.exit()
}

addAdmin()
