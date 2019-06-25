import request from '@/utils/request'

/**
 * add a sudoer
 * @param {String} uid user id
 * @returns {Promise<Object>}
 */
export function addSudoer(uid) {
  return request({
    url: `/users/sudoers/add/${uid}`,
    method: 'post'
  })
}
/**
 * Get users list
 * @returns {Promise<Object>} users list
 */
export function getUsers() {
  return request({
    url: '/users/list',
    method: 'get'
  })
}
/**
 * Remove a sudoer
 * @param {String} uid user id
 * @returns {Promise<Object>}
 */
export function removeSudoer(uid) {
  return request({
    url: `/users/sudoers/remove/${uid}`,
    method: 'post'
  })
}
