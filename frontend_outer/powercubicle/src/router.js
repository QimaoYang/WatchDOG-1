import Vue from 'vue'
import Router from 'vue-router'
import HallSeat from '@/pages/hallseat/HallSeat'
import Index from '@/pages/index/Index'
import UserLogin from '@/pages/userlogin/UserLogin'
import UserRegis from '@/pages/userregis/UserRegis'
Vue.use(Router)

export default new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes: [
    {
      path: '/seatlocinfo',
      name: 'HallSeat',
      component: HallSeat
    },
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/userlogin',
      name: 'UserLogin',
      component: UserLogin
    },
    {
      path: '/userregis',
      name: 'UserRegis',
      component: UserRegis
    }
  ]
})
