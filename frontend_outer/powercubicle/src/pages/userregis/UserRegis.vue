<template>
    <div>
        <mt-header title="注册">
            <router-link to="/" slot="left">
            <mt-button icon="back">返回</mt-button>
            </router-link>
        </mt-header>
        <br>
        <mt-field label="工号" v-model="username" type="tel" placeholder="请输入工号" :attr="{ maxlength: 7 }"></mt-field>
        <br>
        <mt-field label="密码" v-model="userpwd" type="password" placeholder="请输入密码" :attr="{ maxlength: 25 }"></mt-field>
        <br>
        <mt-field label="确认密码" v-model="userpwdagain" type="password" placeholder="请再次输入密码" :attr="{ maxlength: 25 }"></mt-field>
        <br>
        <mt-button v-on:click="userRegis" type="primary" size="large">注册</mt-button>
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
          this.$usrRegis(this.username, this.userpwd)
            .then((response) => {
              if (response === SUCCESS) {
                this.$router.push({
                  path: '/'
                })
              } else if (response === NAMEDUP) {
                alert('工号已存在')
                this.username = ''
                this.userpwd = ''
                this.userpwdagain = ''
              } else {
                alert('注册失败，请联系管理员')
                this.username = ''
                this.userpwd = ''
                this.userpwdagain = ''
              }
            }, err => {
              if (err.response.status === 400) {
                if (confirm('工号已存在，是否前去登录')) {
                  this.$router.go(-1)
                }
              }
              console.log(err)
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
