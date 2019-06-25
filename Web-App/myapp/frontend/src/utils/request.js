import axios from 'axios'
import store from '../store'

// Create axios instance
const service = axios.create({
  baseURL: process.env.BASE_API,
  timeout: 15000
})

// request interceptors
service.interceptors.request.use(
  config => {
    if (store.getters.token) {
      config.headers['Authorization'] = localStorage.getItem('user-token') // Let each request carry a custom token Please modify it according to the actual situation
    }
    return config
  },
  error => {
    // Do something with request error
    Promise.reject(error)
  }
)

// response interceptors
service.interceptors.response.use(
  response => {
    /**
     * The code is non-20000 error-free
     */
    if (!response || !response.data) {
      return Promise.reject('Empty response')
    }
    const res = response.data
    if (res.code !== 20000 && res.code !== 20100) {
      switch (res.code) {
        // Silent error
        case 40066:
          return response.data
        // Invalid username
        case 40316:
          return Promise.reject('invalid username')
        // User disconnected
        case 50008:
        case 50012:
        case 50014:
          store.dispatch('FedLogOut').then(() => {
            location.reload() // To re-instantiate the vue-router object Avoid bugs
          })
          break
        // Default error message
        default:
          return response.data
      }
    } else {
      return response.data
    }
  },
  error => {
    console.error('request error debug : ', error) // for debug
    return Promise.reject(error)
  }
)

export default service
