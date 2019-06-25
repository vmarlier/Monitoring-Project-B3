'use strict'
const UsersController = require(`../controllers/users`)
const Joi = require('joi')

const Users = {
  name: 'Users',
  version: '1.0.0',
  register: async function(server, options) {
    server.route([
      {
        method: 'POST',
        path: '/users/sudoers/add/{uid}',
        config: {
          validate: {
            params: {
              uid: Joi.string().required()
            }
          },
          handler: (req, h) => {
            return UsersController.addSudoer(req, h)
          },
          description: 'add sudoer role to user',
          tags: ['api', 'users', 'sudoer']
        }
      },
      {
        method: 'GET',
        path: '/users/list',
        config: {
          handler: (req, h) => {
            return UsersController.getUsers(req, h)
          },
          description: 'Get users list',
          tags: ['api', 'users']
        }
      },
      {
        method: 'POST',
        path: '/users/sudoers/remove/{uid}',
        config: {
          validate: {
            params: {
              uid: Joi.string().required()
            }
          },
          handler: (req, h) => {
            return UsersController.removeSudoer(req, h)
          },
          description: 'Remove sudoer role from user',
          tags: ['api', 'users', 'sudoer']
        }
      }
    ])
  }
}

module.exports = Users
