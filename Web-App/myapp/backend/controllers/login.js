const User = require(`../models/User`)

const JWT = require('jsonwebtoken')
const TOKEN_KEY = process.env.TOKEN_KEY

class Login {
  async getUserToken(req, h) {
    console.log('getUserToken')
    let isMatch = false
    try {
      const user = await User.findOne({ username: req.payload.username }).exec()
      if (!user) {
        throw new Error('admin not existing')
      }

      isMatch = await user.comparePassword(req.payload.password)

      if (!isMatch) {
        throw new Error('wrong credentials')
      }

      const token = await this.createUserJWTToken(TOKEN_KEY, user)
      const response = h.response({
        code: 20000,
        data: {
          token: token,
          user: user
        }
      })
      console.log(token)
      response.header('Authorization', token)
      return response
    } catch (error) {
      console.warn(`Error on login : ${error.message}`)
      return {
        code: 50000,
        data: null,
        error: error.message
      }
    }
  }

  /**
   * Generate a signed admin token from TOKEN_KEY and admin object
   * @param {String} TOKEN_KEY
   * @param {Object} user
   * @return {String} Signed admin token
   */
  async createUserJWTToken(TOKEN_KEY, user) {
    const session = {}
    session.id = user.id
    session.username = user.username
    session.roles = user.roles
    return JWT.sign(session, TOKEN_KEY)
  }

  /**
   * Get Users info
   */

  logout(req, h) {
    return {
      code: 20000
    }
  }
}

module.exports = new Login()
