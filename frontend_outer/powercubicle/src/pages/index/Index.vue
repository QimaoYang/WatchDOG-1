<template>
    <div>
      <mt-navbar v-model="selected">
        <mt-tab-item id="index">首页</mt-tab-item>
        <mt-tab-item id="userinfo">用户</mt-tab-item>
      </mt-navbar>
      <mt-tab-container v-model="selected">
        <mt-tab-container-item id="index">
          <div class="seat_info" :style="seatinfoDiv">
            <img src="../../assets/images/topview.png" width="100%" alt="seats" usemap="#seatmap"/>
              <a :href="'/seatlocinfo?areaId=A'" class="areaACube" :style="{'width': getClientRem * 308 + 'px', 'left': getClientRem * 86 + 'px', 'top': getClientRem * 206 + 'px', 'height': getClientRem * 424 + 'px', 'border-radius': getClientRem * 60 + 'px', 'background-color': areaColorA}">
                <p class="textStyle" style="font-size: 20px">剩余<br>
                {{seatInfo[0]['A']}}
                </p>
              </a>
              <a :href="'/seatlocinfo?areaId=B'" class="areaBCube" :style="{'width': getClientRem * 615 + 'px', 'left': getClientRem * 437 + 'px', 'top': getClientRem * 377 + 'px', 'height': getClientRem * 253 + 'px', 'border-radius': getClientRem * 60 + 'px', 'background-color': areaColorB}">
                <p class="textStyle" style="font-size: 20px">剩余<br>
                {{seatInfo[1]['B']}}
                </p>
              </a>
              <a :href="'/seatlocinfo?areaId=C'" class="areaCCube" :style="{'width': getClientRem * 239 + 'px', 'left': getClientRem * 26 + 'px', 'top': getClientRem * 900 + 'px', 'height': getClientRem * 414 + 'px', 'border-radius': getClientRem * 60 + 'px', 'background-color': areaColorC}">
                <p class="textStyle" style="font-size: 20px">剩余<br>
                {{seatInfo[2]['C']}}
                </p>
              </a>
            <!-- <map name="seatmap">
              <area shape="rect" coords="0,0,800,600" alt="Sun" href="http://www.baidu.com" target="_blank"/>
              <area shape="circle" coords="90,58,3" alt="Mercury" href="http://www.163.com" target="_parent"/>
              <area shape="circle" coords="124,58,3" alt="Venus" href="http://www.qq.com" target="_top"/>
            </map> -->
            <!-- <li v-for="item in seatInfo" :key="item[0]">
              {{ Object.keys(item)[0] }} 区域： 剩余座位
              <a :href="'/seatlocinfo?areaId='+Object.keys(item)[0]">
              {{ item[Object.keys(item)[0]] }}
              </a>
            </li> -->
          </div>
        </mt-tab-container-item>
        <mt-tab-container-item id="userinfo">
          <div v-if="usernametoshow!==''">
              <mt-field label="工号" v-model="usernametoshow" :readonly=true :disableClear=true></mt-field>
              <br>
              <div v-if="!hasSeat">
                  <p class="indict_info" style="font-size: 20px">您还没有座位，请访问：</p>
                  <p class="indict_info" style="font-size: 20px">10.124.108.253:3000</p>
                  <p class="indict_info" style="font-size: 20px">获取座位二维码并使用微信扫码订座</p>
                  <p class="indict_info" style="font-size: 20px">提示：Home of Team 无需占座</p>
              </div>
              <div v-else>
                  <mt-field label="当前工位" v-model="userseat" :readonly=true :disableClear=true></mt-field>
                  <br>
                  <mt-button v-on:click="releaseSeat" type="danger" size="large">释放座位</mt-button>
              </div>
              <br>
              <mt-button v-on:click="gotoChangePwd" size="large">更改密码</mt-button>
          </div>
          <div v-else>
              <br>
              <div class="container">
                  <mt-field label="工号" v-model="username" type="tel" placeholder="请输入工号" :attr="{ maxlength: 7 }"></mt-field>
                  <br>
                  <mt-field label="密码" v-model="userpwd" type="password" placeholder="请输入密码" :attr="{ maxlength: 25 }"></mt-field>
                  <br>
                  <mt-button v-on:click="userLogin" type="primary" size="large">登录</mt-button>
              </div>
              <br>
              <mt-button v-on:click="gotoRegis" size="large">注册</mt-button>
          </div>
        </mt-tab-container-item>
      </mt-tab-container>
    </div>
</template>
<script>
import { getCookie, setCookie } from '../../localstorage'
const SUCCESS = 'SUCCESS'
const ERROR = 'ERROR'
// const NAMEDUP = "NAMEDUP"
const NAMEERROR = 'NAMEERROR'
const PWDERROR = 'PWDERROR'
// const RESERVED = "RESERVED"
// const FORMATERROR = "FORMATERROR"
const TOKEN = 'PCSessionToken'
const USERNAME = 'PCUserName'
const USERSEAT = 'PCUserSeat'
// const USERSEAT = 'PCUserSeat'
export default {
  name: 'Index',
  data () {
    return {
      selected: 'index',
      seatInfo: [ { 0: { 'A': 0 } }, { 1: { 'B': 0 } }, { 2: { 'C': 0 } } ], // 座位对象list
      seatTypeList: [], // 座位类型list
      load: false, // 加载dom的控制
      encryptCode: '',
      usernametoshow: '',
      userseat: '',
      hasSeat: false,
      username: '',
      userpwd: '',
      seatinfoDiv: {
        backgroundImage: require('../../assets/images/topview.png')
      },
      areaColorA: '',
      areaColorB: '',
      areaColorC: ''
      // screenrem: 10
    }
  },
  mounted () {
    this.encryptCode = this.getQueryString('encryptCode')
    if (this.encryptCode !== null) {
      if (getCookie(USERNAME) === '' || getCookie(TOKEN) === '') {
        if (confirm('尚未登录，前去登录')) {
          this.$router.push({
            path: '/'
          })
          this.selected = 'userinfo'
        }
      } else {
        this.$getCurrentSeat()
          .then((seatId) => {
            if (seatId !== '' && seatId !== ERROR) {
              this.userseat = seatId
              this.hasSeat = true
              this.$toast({
                message: '请勿重复占座',
                duration: 800
              })
              this.$router.push({
                path: '/'
              })
            } else {
              this.regisSeat(this.encryptCode)
            }
          }, err => {
            if (err.response.status === 401) {
              if (confirm('登录已过期，前去登录')) {
                this.usernametoshow = ''
                setCookie(USERNAME, '', 1000)
                setCookie(TOKEN, '', 1000)
                setCookie(USERSEAT, '', 1000)
                this.selected = 'userinfo'
              }
            }
          })
      }
    }
    this.getSeatResNum()
    this.getUserName()
    this.getUserSeat()
  },
  computed: {
    getClientRem: function () {
      let rem = (document.body.clientWidth || window.innerWidth || document.documentElement.clientWidth) / 1080
      return rem
    }
  },
  methods: {
    getQueryString: function (name) {
      let reg = `(^|&)${name}=([^&]*)(&|$)`
      let r = window.location.search.substr(1).match(reg)
      if (r != null) {
        return unescape(r[2])
      } else {
        return null
      }
    },
    getSeatResNum: function () {
      this.$getSeatNum()
        .then((resSeatInfo) => {
          this.seatInfo = resSeatInfo
          this.seatInfo.forEach(item => {
            var areaColor = '#e00000'
            if (item[Object.keys(item)[0]] > 5) {
              areaColor = '#05d700'
            } else if (item[Object.keys(item)[0]] <= 5 && item[Object.keys(item)[0]] > 0) {
              areaColor = '#ffea00'
            }
            switch (Object.keys(item)[0]) {
              case 'A':
                this.areaColorA = areaColor
                break
              case 'B':
                this.areaColorB = areaColor
                break
              case 'C':
                this.areaColorC = areaColor
            }
          })
        }, err => {
          if (err.response.data.indexOf('expired') !== -1) {
            if (confirm('登录已过期，前去登录')) {
              this.usernametoshow = ''
              setCookie(USERNAME, '', 1000)
              setCookie(TOKEN, '', 1000)
              setCookie(USERSEAT, '', 1000)
              this.selected = 'userinfo'
            }
          }
        })
    },
    regisSeat: function (encryptCode) {
      this.$seatRegis(encryptCode)
        .then((seatId) => {
          if (seatId !== '') {
            alert('占座成功！')
            this.userseat = seatId
            this.hasSeat = true
            this.$router.push({
              path: '/'
            })
            this.selected = 'userinfo'
          } else {
            alert('占座失败，请联系管理员')
            this.$router.push({
              path: '/'
            })
          }
        }, err => {
          if (err.response.data.indexOf('expired') !== -1) {
            if (confirm('登录已过期，前去登录')) {
              this.usernametoshow = ''
              setCookie(USERNAME, '', 1000)
              setCookie(TOKEN, '', 1000)
              setCookie(USERSEAT, '', 1000)
              this.selected = 'userinfo'
            }
          } else if (err.response.data.indexOf('booked') !== -1) {
            alert('该座位已被使用')
            this.$router.push({
              path: '/'
            })
          }
        })
    },
    userLogin: function () {
      var namereg = /^[0-9]{7}$/
      var pwdreg = /^(?=.*[a-zA-Z])(?=.*\d)(?=.*[~!@#$%^&*()_+`\-={}:";'<>?,./]).{6,16}$/
      if (!namereg.test(this.username)) {
        this.$toast({
          message: '工号格式不正确',
          duration: 800
        })
      } else if (!pwdreg.test(this.userpwd)) {
        this.$toast({
          message: '密码格式不正确',
          duration: 800
        })
      } else {
        this.$usrLogin(this.username, this.userpwd)
          .then((response) => {
            if (response === SUCCESS) {
              this.usernametoshow = this.username
              this.getUserSeat()
            } else if (response === NAMEERROR) {
              if (confirm('工号不存在，是否前去注册')) {
                this.$router.push({
                  path: '/userregis'
                })
              }
            } else if (response === PWDERROR) {
              alert('密码错误')
            } else {
              alert('登录失败，请联系管理员')
            }
          }, err => {
            if (err.response.data.indexOf('username') >= 0) {
              if (confirm('工号不存在，是否前去注册')) {
                this.$router.push({
                  path: '/userregis'
                })
              }
            } else if (err.response.data.indexOf('password') >= 0) {
              alert('密码错误')
              this.userpwd = ''
            } else {
              alert('登录失败，请联系管理员')
            }
          })
      }
    },
    getUserName: function () {
      this.usernametoshow = getCookie(USERNAME)
    },
    getUserSeat: function () {
      this.$getCurrentSeat()
        .then((seatId) => {
          if (seatId !== ERROR) {
            this.userseat = seatId
            this.hasSeat = true
          }
        }, err => {
          if (err.response.status === 401) {
            if (confirm('登录已过期，前去登录')) {
              this.usernametoshow = ''
              setCookie(USERNAME, '', 1000)
              setCookie(TOKEN, '', 1000)
              setCookie(USERSEAT, '', 1000)
              this.selected = 'userinfo'
            }
          }
        })
    },
    gotoRegis: function () {
      this.$router.push({
        path: '/userregis'
      })
    },
    gotoChangePwd: function () {
      this.$router.push({
        path: '/userpwdchange'
      })
    },
    releaseSeat: function () {
      this.$seatRelease()
        .then((result) => {
          if (result !== SUCCESS) {
            alert('释放失败，请联系管理员')
          } else {
            this.userseat = ''
            this.hasSeat = false
            this.getUserSeat()
          }
        }, err => {
          if (err.response.status === 401) {
            if (confirm('登录已过期，前去登录')) {
              this.usernametoshow = ''
              setCookie(USERNAME, '', 1000)
              setCookie(TOKEN, '', 1000)
              setCookie(USERSEAT, '', 1000)
              this.selected = 'userinfo'
            }
          }
        })
    }
  }
}
</script>
<style lang="stylus" rel="stylesheet/stylus" scoped="scoped">
.btn
  position fixed
  bottom 15px
  left 50%
.indict_info
  text-align center
  color #666
  padding-bottom 5px
.seat_info
  width:100%
  height: 100%
  z-index: 10
  text-align: center
  margin-bottom:50px
  background-repeat: no-repeat
  background-position: center
  position: relative
.areaACube
  position: absolute
  top: 30px
  left: 30px
  height: 390px
  text-align: center
  display: table
  margin: auto
  border-radius: 30px
  color: #fff
  cursor: pointer
.areaBCube
  position: absolute
  top: 30px
  left: 390px
  height: 390px
  text-align: center
  display: table
  margin: auto
  border-radius: 30px
  color: #fff
  cursor: pointer
.areaCCube
  position: absolute
  top: 480px
  left: 30px
  height: 390px
  text-align: center
  display: table
  margin: auto
  border-radius: 30px
  color: #fff
  cursor: pointer
.textStyle
  display: table-cell
  vertical-align: middle
</style>
