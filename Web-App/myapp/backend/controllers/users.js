'use strict'

const axios = require('axios')

class Users {
  async addSudoer(request, h) {
    if (request.params.uid) {
      return axios
        .post(`http://89.234.183.248:8002/users/sudoers/add?uid=${request.params.uid}`, { uid: request.params.uid })
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

  async getUsers(request, h) {
    return axios
      .get('http://89.234.183.248:8002/users/list')
      .then(response => {
        if (response && response.data && response.data.Body) {
          return {
            code: 20000,
            data: response.data.Body
          }
        } else {
          return {
            code: 50000,
            data: null
          }
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

  async removeSudoer(request, h) {
    if (request.params.uid) {
      return axios
        .post(`http://89.234.183.248:8002/users/sudoers/remove?uid=${request.params.uid}`, { uid: request.params.uid })
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

module.exports = new Users()
