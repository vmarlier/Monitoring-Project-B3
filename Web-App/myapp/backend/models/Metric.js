const mongoose = require('mongoose')
const Schema = mongoose.Schema

const metricSchema = new Schema({
  hostname: {
    type: String,
    required: true
  },
  ipv4: String,
  os: String,
  uptime: String,
  process: String,
  memory: {
    total: String,
    free: String,
    usedPerc: String
  },
  disk: [
    {
      path: String,
      total: String,
      free: String,
      used: String,
      usedPerc: String
    }
  ],
  dateCreated: {
    type: Date,
    default: Date.now()
  }
})

metricSchema.index({ hostname: 1 })

module.exports = mongoose.model('Metric', metricSchema)
