import Vue from 'vue'
import App from './App.vue'
import router from './router'
import 'amfe-flexible'
import './assets/stylus/reset.styl'
import { post, get, patch, put } from './http'
import { getSeatSitu, getSeatNum, usrLogin, usrRegis } from './sqlserver'
import Vconsole from 'vconsole'
var VueTouch = require('vue-touch')
Vue.config.productionTip = false
if (process.env.NODE_ENV === 'development') {
  // eslint-disable-next-line no-unused-vars
  let vConsole = new Vconsole()
}
Vue.use(VueTouch, { name: 'v-touch' })
// 定义全局变量
Vue.prototype.$post = post
Vue.prototype.$get = get
Vue.prototype.$patch = patch
Vue.prototype.$put = put
Vue.prototype.$getSeatSitu = getSeatSitu
Vue.prototype.$getSeatNum = getSeatNum
Vue.prototype.$usrLogin = usrLogin
Vue.prototype.$usrRegis = usrRegis
Vue.prototype.$isLogged = false
Vue.prototype.$userName = ''
Vue.prototype.$userSeat = ''
new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
