'use strict'
const MetricsController = require(`../controllers/metrics`)

const Metrics = {
  name: 'Metrics',
  version: '1.0.0',
  register: async function(server, options) {
    server.route([
      {
        method: 'GET',
        path: '/metrics',
        config: {
          handler: (req, h) => {
            return MetricsController.getMetrics(req, h)
          },
          description: 'Get metrics',
          tags: ['api', 'metrics']
        }
      }
    ])
  }
}

module.exports = Metrics
