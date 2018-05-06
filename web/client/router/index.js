import Vue from 'vue'
import Router from 'vue-router'
// import Home from '../views/Home'
import Welcome from '../views/Welcome'
import Login from '../views/Login'
import Feedback from '../views/Feedback'
import Sessions from '../views/Sessions'
import SessionSingle from '../views/SessionSingle'

Vue.use(Router)

export default new Router({
  mode: 'hash',
  routes: [
    {
      path: '/',
      name: 'welcome',
      component: Welcome
    },
    {
      path: '/login',
      name: 'login',
      component: Login
    },
    {
      path: '/sessions',
      name: 'sessions',
      component: Sessions
    },
    {
      path: '/session',
      name: 'session',
      props: true,
      component: SessionSingle
    },
    {
      path: '/feedback',
      name: 'feedback',
      component: Feedback
    }
  ]
})
