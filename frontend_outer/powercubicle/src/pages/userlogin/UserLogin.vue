<template>
    <div class="container">
        <p>工号:</p>
        <br>
        <input v-model="username" type="text" pattern="[0-9]*">
        <p>密码:</p>
        <br>
        <input v-model="userpwd" type="password">
        <br>
        <button v-on:click="userLogin">登录</button>
        <button v-on:click="gotoRegis">注册</button>
    </div>
</template>
<script>
const SUCCESS = 'SUCCESS'
// const ERROR = 'ERROR'
// const NAMEDUP = "NAMEDUP"
const NAMEERROR = 'NAMEERROR'
const PWDERROR = 'PWDERROR'
// const RESERVED = "RESERVED"
// const FORMATERROR = "FORMATERROR"
export default {
  name: 'UserLogin',
  data () {
    return {
      username: '',
      userpwd: ''
    }
  },
  mounted () {
  },
  methods: {
    userLogin: function () {
      this.$usrLogin(this.username, this.userpwd)
        .then((response) => {
          if (response === SUCCESS) {
            // return
          } else if (response === NAMEERROR) {
            alert('工号不存在')
          } else if (response === PWDERROR) {
            alert('密码错误')
          } else {
            alert('登录失败，请联系管理员')
          }
        }, err => {
          console.log(err)
        })
    },
    gotoRegis: function () {
      this.$router.push({
        path: '/userregis'
      })
    }
  }
}
</script>
