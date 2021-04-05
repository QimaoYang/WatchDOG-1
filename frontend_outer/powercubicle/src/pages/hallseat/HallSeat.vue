<template>
  <div>
    <div>
      <mt-header v-bind:title="'区域' + areaId">
        <router-link to="/" slot="left">
          <mt-button icon="back">返回</mt-button>
        </router-link>
      </mt-header>
    </div>
    <div class="wapper">
      <seat-area :propThumbnailAreaWidth="thumbnailBoxWidth" :propThumbnailAreaHeight="thumbnailBoxHeight"
      :propYMax="yMax" :propSeatScale="seatScale" :propSeatHeight="positionDistin" :propSeatToolArr="seatToolArr"
      :propSeatAreaWidthRem="seatAreaWidthRem" :propSeatAreaHeightRem ="seatAreaHeightRem"
      :propSeatBoxHeight="seatBoxHeight" :propMiddleLine="middleLine" :propHorizontalLine="horizontalLine" :propAreaId="areaId" ref="seatArea">
        <template slot="thumbnail-seat-solt">
          <template v-for="seatItem in seatList" >
              <div class="thumbnailSeatClass" :key="'thumbnail'+seatItem.id" :style="{height:thumbnailHeight +'rem',
              width:thumbnailWidth +'rem',background: thumbnailBackgroud(seatItem),
              top:seatItem.gRow * thumbnailPositionDistin +'rem',left:seatItem.gCol * thumbnailPositionDistin +'rem'}">
              </div>
            </template>
        </template>
        <template slot="seat-area-solt">
          <div class="seatBox" :style="{transform: 'scale('+seatScale+')',height:seatBoxHeight +'rem',
          width:seatBoxWidth +'rem',marginLeft:seatBoxCenterMargin+'rem'}">
            <template v-for="(seatItem,index) in seatList" >
              <div class="seatClass" :key="seatItem.id" :style="{height:height +'rem',width:width +'rem',
              top:seatItem.gRow * positionDistin +'rem',left:seatItem.gCol * positionDistin +'rem'}"
              >
                <img :class="seatItem.iconClass" :seatId="seatItem.id" :seatIndex="index" :src="seatItem.nowIcon"/>
              </div>
            </template>
          </div>
        </template>
      </seat-area>
      <loading :load="load"></loading>
    </div>
  </div>
</template>
<script>
import SeatArea from './component/SeatArea'
import Loading from '@/components/loading'
export default {
  name: 'HallSeat',
  data () {
    return {
      seatList: [],
      seatTypeList: [],
      positionDistin: 0.5,
      width: 0.5,
      height: 0.5,
      thumbnailWidth: 0.1,
      thumbnailHeight: 0.1,
      thumbnailPositionDistin: 0.15,
      seatAreaWidthRem: 10,
      load: false,
      resSeatInfo: [],
      areaId: ''
    }
  },
  components: {
    SeatArea,
    Loading
  },
  mounted () {
    this.loading(true)
    this.areaId = this.getQueryString('areaId')
    this.getSeatList()
    this.loading(false)
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
    getSeatList: function () {
      this.$getSeatSitu(this.areaId)
        .then((seatInfo) => {
          this.resSeatInfo = seatInfo
          this.$get('/mock/seatInfo.json')
            .then((seatGeoInfo) => {
              if (seatGeoInfo.errorCode !== 0) {
                alert(seatGeoInfo.errorMsg)
                this.$router.go(-1)
              }
              var resSeatList = seatGeoInfo['seatList' + this.areaId]
              if (!resSeatList) {
                return
              }
              resSeatList.forEach(element => {
                if (element.team !== 'public') {
                  element.nowIcon = require('../../assets/images/' + element.team + 'Home.png')
                } else {
                  let isAva = '1'
                  if (this.resSeatInfo[element.id]['state'] === 'availble') {
                    isAva = '0'
                  }
                  if (isAva === '0') {
                    element.nowIcon = require('../../assets/images/avaSeat.png')
                  } else {
                    element.nowIcon = require('../../assets/images/unAvaSeat.png')
                  }
                  // for (const key in seatType) {
                  //   if (isAva === seatType[key].type) {
                  //     element.nowIcon = require(seatType[key].icon)
                  //   }
                  // }
                }
              })
              this.seatList = resSeatList
              this.seatTypeList = seatGeoInfo.seatTypeList
            }, err => {
              alert('获取区域座位情况失败！')
              console.log(err)
              this.$router.push({
                path: '/'
              })
            })
        }, err => {
          alert('获取区域座位情况失败！')
          console.log(err)
          this.$router.push({
            path: '/'
          })
        })
    },
    back: function () {
      this.$router.go(-1)
    },
    loading: function (value) {
      this.load = value
    },
    thumbnailBackgroud: function (seatItem) {
      if (seatItem.nowIcon === seatItem.selectedIcon) {
        return 'green'
      } else if (seatItem.nowIcon === seatItem.soldedIcon) {
        return 'red'
      } else if (seatItem.nowIcon === seatItem.fixIcon) {
        return 'red'
      } else {
        return 'white'
      }
    }
  },
  computed: {
    seatAreaHeightRem: function () {
      let screenRem = (document.body.clientWidth || window.innerWidth || document.documentElement.clientWidth) / 10
      let height = document.documentElement.clientHeight || window.innerHeight || document.body.clientHeight
      let otherDom = 1.08
      return height / screenRem - otherDom
    },
    xMax: function () {
      let i = 0
      for (let index in this.seatList) {
        if (this.seatList[index].gCol > i) {
          i = this.seatList[index].gCol
        }
      }
      if (this.areaId === 'A') {
        i = i + 2
      }
      return i
    },
    yMax: function () {
      let i = 0
      for (let index in this.seatList) {
        if (this.seatList[index].gRow > i) {
          i = this.seatList[index].gRow
        }
      }
      return i
    },
    middleLine: function () {
      return ((this.xMax / 2) + 1) * this.positionDistin - 0.025
    },
    horizontalLine: function () {
      return ((this.yMax / 2) + 1) * this.positionDistin - 0.025
    },
    seatScale: function () {
      let seatScaleX = 1
      let seatScaleY = 1
      seatScaleX = this.seatAreaWidthRem / this.seatBoxWidth
      seatScaleY = this.seatAreaHeightRem / this.seatBoxHeight
      return seatScaleX < seatScaleY ? seatScaleX : seatScaleY
    },
    seatBoxCenterMargin: function () {
      return -(this.seatBoxWidth * this.seatScale) / 2
    },
    seatBoxHeight: function () {
      return (this.yMax + 1) * this.positionDistin + this.height
    },
    seatBoxWidth: function () {
      return (this.xMax + 1) * this.positionDistin + this.width
    },
    thumbnailBoxWidth: function () {
      return ((this.xMax + 1) * this.thumbnailPositionDistin + this.thumbnailWidth)
    },
    thumbnailBoxHeight: function () {
      return ((this.yMax + 1) * this.thumbnailPositionDistin + this.thumbnailHeight)
    },
    seatToolArr: function () {
      let seatToolArr = []
      let yMax = this.yMax
      for (let i = 1; i <= yMax; i++) {
        let el = this.seatList.find((item) => (
          item.gRow === i
        ))
        if (el) {
          seatToolArr.push(el.row)
        } else {
          seatToolArr.push('')
        }
      }
      return seatToolArr
    }
  }
}
</script>
<style lang="stylus" rel="stylesheet/stylus" scoped="scoped">
  .wapper
    width: 270px
    width 750px
    background #f3f4f6
    .seat-detail-item
      display flex
      align-content center
      .seatTypeClass
        display block
        height 35px
        line-height 35px
        white-space:nowrap
    .thumbnailSeatClass
      position absolute
    .seatBox
      position absolute
      left 50%
      transform-origin 0rem 0rem 0rem
      .middle-line
        position absolute
        border-right 0.05rem rgba(0,0,0,0.2) dashed
      .seatClass
        position absolute
        .seatImgClass
          position absolute
          left 0
          top 0
          height 100%
        .seatDoubleHeightImgClass
          position absolute
          left 0
          top 0
          height 200%
        .seatTriHeightImgClass
          position absolute
          left 0
          top 0
          height 300%
        .seatSixHeightImgClass
          position absolute
          left 0
          top 0
          height 600%
</style>
