<template>
    <div class="container">
        <li v-for="item in seatInfo" :key="item[0]">
          {{ Object.keys(item)[0] }} 区域： 剩余座位
          <a :href="'/seatlocinfo?areaId='+Object.keys(item)[0]">
          {{ item[Object.keys(item)[0]] }}
          </a>
        </li>
    </div>
</template>
<script>
export default {
  name: 'Index',
  data () {
    return {
      seatInfo: [], // 座位对象list
      seatTypeList: [], // 座位类型list
      load: false, // 加载dom的控制
      encryptCode: ''
    }
  },
  mounted () {
    this.encryptCode = this.getQueryString('encryptCode')
    if (this.encryptCode !== null) {
      this.regisSeat(this.encryptCode)
    }
    this.getSeatResNum()
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
          console.log(this.seatInfo)
        }, err => {
          console.log(err)
        })
    },
    regisSeat: function (encryptCode) {
      this.$seatRegis(encryptCode)
        .then((seatId) => {
          if (seatId !== '') {
            alert('座位注册成功！')
          } else {
            alert('座位注册失败，请联系管理员')
          }
        })
    }
  }
}
</script>
