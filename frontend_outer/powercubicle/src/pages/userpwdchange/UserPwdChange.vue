<template>
    <div>
        <mt-header title="更改密码">
            <router-link to="/" slot="left">
            <mt-button icon="back">返回</mt-button>
            </router-link>
        </mt-header>
        <br>
        <mt-field label="工号" v-model="username" :readonly=true :disableClear=true></mt-field>
        <br>
        <mt-field label="新密码" v-model="userpwd" type="password" placeholder="请输入新的密码" :attr="{ maxlength: 25 }"></mt-field>
        <br>
        <mt-field label="确认密码" v-model="userpwdagain" type="password" placeholder="请再次输入新的密码" :attr="{ maxlength: 25 }"></mt-field>
        <br>
        <mt-button v-on:click="userPwdChange" type="primary" size="large">确认修改</mt-button>
    </div>
</template>
<script>
import { getCookie } from '../../localstorage'
const SUCCESS = 'SUCCESS'
// const ERROR = 'ERROR'
// const NAMEDUP = 'NAMEDUP'
// const NAMEERROR = 'NAMEERROR'
// const PWDERROR = 'PWDERROR'
// const RESERVED = 'RESERVED'
// const FORMATERROR = 'FORMATERROR'
const USERNAME = 'PCUserName'
export default {
  name: 'UserPwdChange',
  data () {
    return {
      username: '',
      userpwd: '',
      userpwdagain: ''
    }
  },
  mounted () {
    this.username = getCookie(USERNAME)
  },
  methods: {
    userPwdChange: function () {
      if (this.userpwd === this.userpwdagain) {
        var namereg = /^[0-9]{7}$/
        var pwdreg = /^(?=.*[a-zA-Z])(?=.*\d)(?=.*[~!@#$%^&*()_+`\-={}:";'<>?,./]).{6,16}$/
        if (!namereg.test(this.username)) {
          this.$toast({
            message: '工号格式不正确（7位数字）',
            duration: 800
          })
        } else if (!pwdreg.test(this.userpwd)) {
          this.$toast({
            message: '密码必须由6-16位字母，数字及符号组成！',
            duration: 800
          })
        } else {
          this.$usrPwdChange(this.userpwd)
            .then((response) => {
              if (response === SUCCESS) {
                this.$toast({
                  message: '密码修改成功！',
                  duration: 800
                })
                this.$router.push({
                  path: '/'
                })
              } else {
                this.$toast({
                  message: '密码修改失败！请重试',
                  duration: 800
                })
                this.userpwd = ''
                this.userpwdagain = ''
              }
            }, err => {
              if (err.response.status === 400) {
                this.$toast({
                  message: '密码修改失败！请重试',
                  duration: 800
                })
                this.userpwd = ''
                this.userpwdagain = ''
              }
            })
        }
      } else {
        this.userpwd = ''
        this.userpwdagain = ''
        alert('两次密码不一致')
      }
    }
  }
}
</script>
