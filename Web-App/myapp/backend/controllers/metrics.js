'use strict'

const axios = require('axios')

class Metrics {
  async getMetrics(request, h) {
    return axios
      .get('http://89.234.183.248:8002/metrics')
      .then(response => {
        return {
          code: 20000,
          data: response.data
        }
      })
      .catch(error => {
        return {
          code: 50000,
          data: null,
          error: error.message
        }
      })
  }
}

module.exports = new Metrics()
