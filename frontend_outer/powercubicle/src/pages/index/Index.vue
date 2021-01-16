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
      load: false // 加载dom的控制
    }
  },
  mounted () {
    this.getSeatResNum()
  },
  methods: {
    getSeatResNum: function () {
      this.$getSeatNum()
        .then((resSeatInfo) => {
          this.seatInfo = resSeatInfo
          console.log(this.seatInfo)
        }, err => {
          console.log(err)
        })
    }
  }
}
</script>
