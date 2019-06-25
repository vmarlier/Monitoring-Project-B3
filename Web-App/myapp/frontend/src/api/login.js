import request from '@/utils/request'

/**
 * Login of the user
 * @param {String} email
 * @param {String} password
 * @returns {Promise<Object>} token of the user
 */
export function login(username, password) {
  return request({
    url: '/user/login',
    method: 'post',
    data: {
      username,
      password
    }
  })
}

/**
 * Log out the user
 */
export function logout() {
  return request({
    url: '/user/logout',
    method: 'post'
  })
}
