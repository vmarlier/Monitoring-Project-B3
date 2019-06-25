const getters = {
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  firstName: state => state.user.firstName,
  lastName: state => state.user.lastName,
  username: state => state.user.username,
  email: state => state.user.email,
  history: state => state.user.history
}
export default getters
