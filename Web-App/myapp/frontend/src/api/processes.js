import request from '@/utils/request'

/**
 * Get processes list
 * @returns {Promise<Object>} processes list
 */
export function getProcesses() {
  return request({
    url: '/processes/list',
    method: 'get'
  })
}
/**
 * Kill a process
 * @param {Number} pid
 * @returns {Promise<Object>}
 */
export function killProcess(pid) {
  return request({
    url: `/processes/kill/${pid}`,
    method: 'post'
  })
}
