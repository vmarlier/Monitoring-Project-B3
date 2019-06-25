'use strict'
const Joi = require('joi') // request validationb
const LoginController = require(`../controllers/login`)

const Login = {
  name: 'Login',
  version: '1.0.0',
  register: async function(server, options) {
    server.route([
      {
        method: 'POST',
        path: '/user/login',
        config: {
          auth: false,
          handler: (req, h) => {
            return LoginController.getUserToken(req, h)
          },
          description: 'Get user token',
          notes: 'Returns user token',
          tags: ['api', 'login'],
          validate: {
            payload: Joi.object().keys({
              username: Joi.string()
                .required()
                .description('username'),
              password: Joi.string().description('password')
            })
          }
        }
      },
      {
        method: 'POST',
        path: '/user/logout',
        config: {
          handler: (req, h) => {
            return LoginController.logout(req, h)
          },
          description: 'Logout',
          notes: 'Logout',
          tags: ['api', 'login']
        }
      }
    ])
  }
}

module.exports = Login
