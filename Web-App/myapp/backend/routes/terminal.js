'use strict'
const TerminalController = require(`../controllers/terminal`)
const Joi = require('joi')

const Terminal = {
  name: 'Terminal',
  version: '1.0.0',
  register: async function(server, options) {
    server.route([
      {
        method: 'POST',
        path: '/terminal/command',
        config: {
          validate: {
            payload: {
              command: Joi.string().required()
            }
          },
          handler: (req, h) => {
            return TerminalController.sendCommand(req, h)
          },
          description: 'Send a command from terminal',
          tags: ['api', 'terminal', 'command']
        }
      }
    ])
  }
}

module.exports = Terminal
