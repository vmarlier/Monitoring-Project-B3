import { login, logout } from '@/api/login'

const user = {
  state: {
    token: localStorage.getItem('user-token'),
    firstName: '',
    lastName: '',
    username: '',
    email: '',
    history: []
  },

  mutations: {
    SET_TOKEN: (state, token) => {
      state.token = token
    },
    SET_NAME: (state, { firstName, lastName }) => {
      state.firstName = firstName
      state.lastName = lastName
    },
    SET_USERNAME: (state, username) => {
      state.username = username
    },
    SET_EMAIL: (state, email) => {
      state.email = email
    },
    ADD_TO_HISTORY: (state, command) => {
      state.history.push(command)
    }
  },

  actions: {
    // Login
    Login({ commit }, { userInfo, shouldRememberUser }) {
      return new Promise((resolve, reject) => {
        login(userInfo.username, userInfo.password)
          .then(response => {
            if (response && response.data) {
              const data = response.data
              const user = data.user
              const storage = shouldRememberUser ? localStorage : sessionStorage
              storage.setItem('user-token', data.token)
              storage.setItem('user-info', JSON.stringify(user))
              commit('SET_TOKEN', data.token)
              commit('SET_USERNAME', user.username)
              commit('SET_EMAIL', user.email)
              commit('SET_NAME', {
                firstName: user.firstName,
                lastName: user.lastName
              })
              commit('SET_AVATAR', user.avatar)
              resolve()
            } else if (response && response.error) {
              resolve(response)
            } else {
              return reject('Empty response')
            }
          })
          .catch(error => {
            localStorage.removeItem('user-token')
            localStorage.removeItem('user-info')
            reject(error)
          })
      })
    },
    // Get user information
    GetInfo({ commit }) {
      return new Promise((resolve, reject) => {
        const user = JSON.parse(localStorage.getItem('user-info'))
        commit('SET_USERNAME', user.username)
        commit('SET_EMAIL', user.email)
        commit('SET_NAME', {
          firstName: user.firstName,
          lastName: user.lastName
        })
      })
    },
    // Signout
    LogOut({ commit, state }) {
      return new Promise((resolve, reject) => {
        logout(state.token)
          .finally(() => {
            commit('SET_TOKEN', '')
            commit('SET_USERNAME', '')
            commit('SET_EMAIL', '')
            commit('SET_NAME', {
              firstName: '',
              lastName: ''
            })
            localStorage.removeItem('user-token')
            localStorage.removeItem('user-info')
            resolve()
          })
          .catch(error => {
            reject(error)
          })
      })
    },

    // Front end
    FedLogOut({ commit }) {
      return new Promise(resolve => {
        commit('SET_TOKEN', null)
        localStorage.removeItem('user-token')
        localStorage.removeItem('user-info')
        resolve()
      })
    },

    // Add command to history
    addToHistory({ commit }, command) {
      commit('ADD_TO_HISTORY', command)
    }
  }
}

export default user
