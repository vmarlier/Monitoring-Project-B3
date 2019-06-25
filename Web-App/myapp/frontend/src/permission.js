import router from './router'
import store from './store'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style

NProgress.configure({
  showSpinner: false
}) // NProgress Configuration

const whiteList = ['/login'] // no redirect whitelist

router.beforeEach((to, from, next) => {
  NProgress.start() // start progress bar
  if (localStorage.getItem('user-token')) {
    // if the user has a token in local storage
    if (!store.getters.username) {
      store.dispatch('GetInfo')
    }
    if (RegExp(/^((\/\w+){0,2}(\/login))$/).test(to.path)) {
      next({ path: '/' })
      NProgress.done()
    } else {
      next()
      NProgress.done()
    }
    // else if (RegExp(/^(\/\w+){0,2}\/404/).test(to.path)) {
    //   // 404 page
    //   next()
    // }
  } else {
    /* has no token*/
    if (whiteList.indexOf(to.path) !== -1) {
      // In the log-in white list, enter directly
      next()
    } else {
      next({ path: '/login' }) // Otherwise all redirect to login page
      NProgress.done()
    }
  }
})

router.afterEach(() => {
  NProgress.done() // finish progress bar
})
