import request from '@/utils/request'

/**
 * Send a command from terminal
 * @param {String} command
 * @returns {Promise<Object>}
 */
export function sendCommand(command) {
  return request({
    url: '/terminal/command',
    method: 'post',
    data: {
      command
    }
  })
}
