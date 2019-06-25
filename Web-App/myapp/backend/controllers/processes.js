'use strict'

const axios = require('axios')

class Processes {
  async getProcesses(request, h) {
    return axios
      .get('http://89.234.183.248:8002/process/list')
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
  async killProcess(request, h) {
    if (request.params.pid) {
      return axios
        .post(`http://89.234.183.248:8002/process/kill?pid=${request.params.pid}`, { pid: request.params.pid })
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

module.exports = new Processes()
