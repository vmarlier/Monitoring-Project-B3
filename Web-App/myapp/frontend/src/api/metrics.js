import request from '@/utils/request'

/**
 * Get metrics
 * @returns {Promise<Object>} metrics
 */
export function getMetrics() {
  return request({
    url: '/metrics',
    method: 'get'
  })
}
