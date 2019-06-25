'use strict'

const path = require('path') // resolving path module
const ENTRY = path.join(global.appRoot, '../frontend/dist/') // if path change, only edit thi
const STATIC = path.join(global.appRoot, '../frontend/dist/static/') // if path change, only edit this

const staticConfig = {
  auth: false,
  handler: {
    directory: {
      path: [STATIC],
      redirectToSlash: true
    }
  },
  description: 'Serve frontend',
  notes: 'xxx',
  tags: ['front', 'assets']
}

const Front = {
  name: 'Front',
  version: '1.0.0',
  register: async function(server, options) {
    server.route([
      {
        method: 'GET',
        path: '/static/{param*}',
        config: staticConfig
      },
      {
        method: 'GET',
        path: '/static/static/{param*}',
        config: staticConfig
      },
      {
        method: 'GET',
        path: '/js/{param*}',
        config: {
          auth: false,
          handler: {
            directory: {
              path: [STATIC + '/js'],
              redirectToSlash: true
            }
          },
          description: 'Serve frontend',
          notes: 'xxx',
          tags: ['front', 'assets']
        }
      },
      {
        method: 'GET',
        path: '/css/{param*}',
        config: {
          auth: false,
          handler: {
            directory: {
              path: [STATIC + '/css'],
              redirectToSlash: true
            }
          },
          description: 'Serve frontend',
          notes: 'xxx',
          tags: ['front', 'assets']
        }
      },
      {
        method: 'GET',
        path: '/img/{param*}',
        config: {
          auth: false,
          handler: {
            directory: {
              path: [STATIC + '/img'],
              redirectToSlash: true
            }
          },
          description: 'Serve frontend',
          notes: 'xxx',
          tags: ['front', 'assets']
        }
      },
      {
        method: 'GET',
        path: '/{param*}',
        config: {
          auth: false,
          handler: {
            file: ENTRY + 'index.html'
          },
          description: 'Serve frontend',
          notes: 'xxx',
          tags: ['front', 'assets']
        }
      }
    ])
  }
}

module.exports = Front
