'use strict'

const axios = require('axios')

class Terminal {
  async sendCommand(request, h) {
    if (request.payload.command) {
      return axios
        .post(`http://89.234.183.248:8002/terminal`, { command: request.payload.command })
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
    return {
      code: 50000,
      data: null
    }
  }
}

module.exports = new Terminal()
