const mongoose = require('mongoose')
const Schema = mongoose.Schema
const bcrypt = require('bcrypt-nodejs')

const userSchema = new Schema({
  email: {
    type: String,
    lowercase: true,
    required: true
  },
  username: {
    type: String,
    required: true
  },
  firstName: {
    type: String,
    required: true
  },
  lastName: {
    type: String,
    required: true
  },
  avatar: {
    type: String,
    default: ''
  },
  password: String,
  isDeactivated: {
    type: Boolean,
    required: true,
    default: false
  },
  dateCreated: {
    type: Date,
    required: true,
    default: Date.now
  },
  dateDeactivated: Date,
  roles: [{ type: String }],
  ips: {
    type: [{ type: String }],
    default: undefined
  }
})

userSchema.index({ email: 1, username: 1 })

// On Save Hook, encrypt password
// Before saving a model, run this function
userSchema.pre('save', function(next) {
  // get access to the user model
  const user = this

  // generate a salt then run callback
  bcrypt.genSalt(10, function(err, salt) {
    if (err) {
      return next(err)
    }

    // hash (encrypt) our password using the salt
    bcrypt.hash(user.password, salt, null, function(err, hash) {
      if (err) {
        return next(err)
      }

      // overwrite plain text password with encrypted password
      user.password = hash
      next()
    })
  })
})

userSchema.methods.comparePassword = async function(candidatePassword, callback) {
  const thisPassword = this.password

  return new Promise((resolve, reject) => {
    bcrypt.compare(candidatePassword, thisPassword, (err, isMatch) => {
      if (err) {
        return reject(err)
      }

      return resolve(isMatch)
    })
  })
}

/*
 * Override toJSON (called by JSON.stringify)
 * to prevent leaking entity details (such as BusinessManager).
 * when sending info to the client (with res.json() or JSON.stringify).
 * All private information is skipped.
 */
userSchema.method('toJSON', function() {
  return {
    username: this.username,
    firstName: this.firstName,
    lastName: this.lastName,
    email: this.email,
    avatar: this.avatar,
    roles: this.roles,
    isDeactivated: this.isDeactivated,
    dateCreation: this.dateCreation,
    dateDeactivated: this.dateDeactivated
  }
})

module.exports = mongoose.model('User', userSchema)
