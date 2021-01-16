<template>
    <div class="container">
        <p>工号:</p>
        <br>
        <input v-model="username" type="text" pattern="[0-9]*">
        <p>密码:</p>
        <br>
        <input v-model="userpwd" type="password">
        <p>确认密码:</p>
        <br>
        <input v-model="userpwdagain" type="password">
        <br>
        <button v-on:click="userRegis">注册</button>
    </div>
</template>
<script>
const SUCCESS = 'SUCCESS'
// const ERROR = 'ERROR'
const NAMEDUP = 'NAMEDUP'
// const NAMEERROR = 'NAMEERROR'
// const PWDERROR = 'PWDERROR'
// const RESERVED = 'RESERVED'
// const FORMATERROR = 'FORMATERROR'
export default {
  name: 'UserRegis',
  data () {
    return {
      username: '',
      userpwd: '',
      userpwdagain: ''
    }
  },
  mounted () {
  },
  methods: {
    userRegis: function () {
      if (this.userpwd === this.userpwdagain) {
        this.$usrRegis(this.username, this.userpwd)
          .then((response) => {
            if (response === SUCCESS) {
              // return
            } else if (response === NAMEDUP) {
              alert('工号已存在')
            } else {
              alert('注册失败，请联系管理员')
            }
          }, err => {
            console.log(err)
          })
      } else {
        this.userpwd = ''
        this.userpwdagain = ''
        alert('两次密码不一致')
      }
    }
  }
}
</script>
