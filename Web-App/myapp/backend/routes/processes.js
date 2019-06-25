'use strict'
const ProcessesController = require(`../controllers/processes`)
const Joi = require('joi')

const Processes = {
  name: 'Processes',
  version: '1.0.0',
  register: async function(server, options) {
    server.route([
      {
        method: 'GET',
        path: '/processes/list',
        config: {
          handler: (req, h) => {
            return ProcessesController.getProcesses(req, h)
          },
          description: 'Get processes list',
          tags: ['api', 'processes']
        }
      },
      {
        method: 'POST',
        path: '/processes/kill/{pid}',
        config: {
          validate: {
            params: {
              pid: Joi.string().required()
            }
          },
          handler: (req, h) => {
            return ProcessesController.killProcess(req, h)
          },
          description: 'Kill a process',
          tags: ['api', 'processes']
        }
      }
    ])
  }
}

module.exports = Processes
