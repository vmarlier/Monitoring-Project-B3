import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      component: () => import('@/views/Layout/index'),
      children: [
        {
          path: '',
          name: 'Dashboard',
          component: () => import('@/views/Dashboard/index')
        },
        {
          path: 'processes',
          name: 'Processes',
          component: () => import('@/views/Processes/index')
        },
        {
          path: 'users',
          name: 'Users',
          component: () => import('@/views/Users/index')
        },
        {
          path: 'terminal',
          name: 'Terminal',
          component: () => import('@/views/Terminal/index')
        }
      ]
    },
    {
      path: '/login',
      name: 'Login',
      component: () => import('@/views/Login/index')
    },
    {
      path: '*',
      redirect: '/'
    }
  ]
})
