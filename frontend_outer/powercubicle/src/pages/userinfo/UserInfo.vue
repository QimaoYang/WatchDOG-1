<template>
    <div v-if="username!==''" class="container">
        <p>工号：  {{username}}</p>
        <br>
        <div v-if="!hasSeat">
            <p>您还没有座位，请访问：
                10.124.108.253:3000
                获取座位二维码并使用微信扫码订座
            </p>
        </div>
        <div v-else>
            <p>当前工位：  {{userseat}}</p>
            <mt-button v-on:click="releaseSeat">释放座位</mt-button>
        </div>
    </div>
    <div v-else>
        <p>请先登录/注册</p>
        <button v-on:click="gotoLogin">登录</button>
        <button v-on:click="gotoRegis">注册</button>
    </div>
</template>
<script>
import { getCookie } from '../../localstorage'
const SUCCESS = 'SUCCESS'
const ERROR = 'ERROR'
// const NAMEDUP = "NAMEDUP"
// const NAMEERROR = 'NAMEERROR'
// const PWDERROR = 'PWDERROR'
// const RESERVED = "RESERVED"
// const FORMATERROR = "FORMATERROR"
const USERNAME = 'PCUserName'
// const USERSEAT = 'PCUserSeat'
export default {
  name: 'UserLogin',
  data () {
    return {
      username: '',
      userseat: '',
      hasSeat: false
    }
  },
  mounted () {
    this.getUserName()
    this.getUserSeat()
  },
  methods: {
    getUserName: function () {
      this.username = getCookie(USERNAME)
    },
    getUserSeat: function () {
      if (this.userseat === '') {
        this.$getCurrentSeat()
          .then((seatId) => {
            if (seatId !== ERROR) {
              this.userseat = seatId
              this.hasSeat = true
              console.log(this.seatId)
            }
          }, err => {
            console.log(err)
          })
      }
    },
    gotoLogin: function () {
      this.$router.push({
        path: '/userlogin'
      })
    },
    gotoRegis: function () {
      this.$router.push({
        path: '/userregis'
      })
    },
    releaseSeat: function () {
      this.$seatRelease()
        .then((result) => {
          if (result !== SUCCESS) {
            alert('释放失败，请联系管理员')
          } else {
            this.hasSeat = false
            this.getUserSeat()
          }
        }, err => {
          console.log(err)
        })
    }
  }
}
</script>
