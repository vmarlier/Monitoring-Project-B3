'use stric'

module.exports = function(mongoose, server) {
  const MONGO_URI = process.env.MONGO_URI || 'mongodb://0.0.0.0:27017/monitoring'

  // Connect ro mongodb
  // Good way to make sure mongoose never stops trying to reconnect
  // The Number.MAX_VALUE property represents the maximum numeric value representable in JavaScript.
  // reconnectTries suggestion from http://mongoosejs.com/docs/connections.html
  // For long running applications, it is often prudent to enable keepAlive with a number of milliseconds. Without it, after some period of time you may start to see "connection closed" errors
  // http://mongoosejs.com/docs/connections.html
  const mongoOptions = {
    auto_reconnect: true,
    keepAlive: 2000,
    connectTimeoutMS: 30000,
    reconnectTries: Number.MAX_VALUE,
    useCreateIndex: true,
    useFindAndModify: false,
    useNewUrlParser: true
  }

  mongoose.connect(MONGO_URI, mongoOptions)

  // If the connection throws an error
  mongoose.connection.on('error', function(err) {
    console.error('MongoDB connection error: ' + err)
  })

  mongoose.connection.on('reconnected', function() {
    console.debug('MongoDB reconnected')
  })

  mongoose.connection.on('connected', function() {
    console.debug('MongoDB connected on', MONGO_URI)
  })

  // When the connection is disconnected, Log the error
  mongoose.connection.on('disconnected', function() {
    console.warn('MongoDB disconnected')
  })

  const cleanup = () => {
    mongoose.connection.close(() => {
      process.exit(0)
    })
  }

  process.on('SIGINT', cleanup)
  process.on('SIGTERM', cleanup)
}
