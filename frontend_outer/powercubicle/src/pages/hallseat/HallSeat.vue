<template>
  <div class="wapper">
     <!--排期详情和座位上方示例图 结束-->
    <seat-area :propThumbnailAreaWidth="thumbnailBoxWidth" :propThumbnailAreaHeight="thumbnailBoxHeight"
    :propYMax="yMax" :propSeatScale="seatScale" :propSeatHeight="positionDistin" :propSeatToolArr="seatToolArr"
    :propSeatAreaWidthRem="seatAreaWidthRem" :propSeatAreaHeightRem ="seatAreaHeightRem"
    :propSeatBoxHeight="seatBoxHeight" :propMiddleLine="middleLine" :propHorizontalLine="horizontalLine" :propAreaId="areaId" ref="seatArea">
      <!--以下为缩略座位图具名插槽 开始-->
      <template slot="thumbnail-seat-solt">
         <template v-for="seatItem in seatList" >
            <div class="thumbnailSeatClass" :key="'thumbnail'+seatItem.id" :style="{height:thumbnailHeight +'rem',
            width:thumbnailWidth +'rem',background: thumbnailBackgroud(seatItem),
            top:seatItem.gRow * thumbnailPositionDistin +'rem',left:seatItem.gCol * thumbnailPositionDistin +'rem'}">
            </div>
          </template>
      </template>
      <!--以上为缩略座位图具名插槽 结束-->
      <!--以下为座位图具名插槽 开始-->
      <template slot="seat-area-solt">
        <div class="seatBox" :style="{transform: 'scale('+seatScale+')',height:seatBoxHeight +'rem',
        width:seatBoxWidth +'rem',marginLeft:seatBoxCenterMargin+'rem'}">
          <template v-for="(seatItem,index) in seatList" >
            <div class="seatClass" :key="seatItem.id" :style="{height:height +'rem',width:width +'rem',
            top:seatItem.gRow * positionDistin +'rem',left:seatItem.gCol * positionDistin +'rem'}"
            >
              <img class="seatImgClass" :seatId="seatItem.id" :seatIndex="index" :src="seatItem.nowIcon"/>
            </div>
          </template>
        </div>
      </template>
      <!--以上为座位图具名插槽 结束-->
    </seat-area>
    <loading :load="load"></loading>
  </div>
</template>
<script>
import SeatArea from './component/SeatArea'
import Loading from '@/components/loading'
export default {
  name: 'HallSeat',
  data () {
    return {
      seatList: [], // 座位对象list
      seatTypeList: [], // 座位类型list
      positionDistin: 0.5, // 每个座位偏移距离
      width: 0.5, // 每个座位的宽
      height: 0.5, // 每个座位的高
      thumbnailWidth: 0.1, // 缩略图每个座位的宽
      thumbnailHeight: 0.1, // 缩略图每个座位的高
      thumbnailPositionDistin: 0.15, // 缩略图每个座位偏移距离
      seatAreaWidthRem: 10, // 座位区域横向rem最大值 用于和 seatAreaHeightRem 共同计算区域缩放比例
      load: false, // 加载dom的控制
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
    // 请求影院列表数据
    getSeatList: function () {
      console.log(this.areaId)
      this.$getSeatSitu(this.areaId)
        .then((seatInfo) => {
          this.resSeatInfo = seatInfo
          console.log('qimao', this.resSeatInfo)
          this.$get('/mock/seatInfo.json')
            .then((seatGeoInfo) => {
              if (seatGeoInfo.errorCode !== 0) {
                alert(seatGeoInfo.errorMsg)
                this.$router.go(-1)
              }
              var resSeatList = seatGeoInfo['seatList' + this.areaId]
              // var resSeatList = seatGeoInfo.seatListA
              // 座位处理 -------开始
              if (!resSeatList) {
                return
              }
              resSeatList.forEach(element => {
                // 获取座位的类型的首字母
                let isAva = '1'
                // 加载座位的图标
                // console.log('element===================', element)
                for (const item of this.resSeatInfo) {
                  if (Object.keys(item)[0] === element.id) {
                    if (item[element.id] === 'availble') {
                      isAva = '0'
                      break
                    } else {
                      isAva = '1'
                      break
                    }
                  }
                }
                let seatType = seatGeoInfo.seatTypeList
                for (const key in seatType) {
                  if (isAva === seatType[key].type) {
                    element.nowIcon = seatType[key].icon
                  }
                }
              })
              this.seatList = resSeatList
              this.seatTypeList = seatGeoInfo.seatTypeList
            }, err => {
              console.log(err)
            })
        }, err => {
          console.log(err)
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
    // 座位区域高度rem
    seatAreaHeightRem: function () {
      let screenRem = (document.body.clientWidth || window.innerWidth || document.documentElement.clientWidth) / 10
      let height = document.documentElement.clientHeight || window.innerHeight || document.body.clientHeight
      // 除了座位区域的大小
      let otherDom = 1.08// 头+排期信息+座位示例 +底部快捷选择+确认按钮
      // 剩下的座位区域的大小
      return height / screenRem - otherDom
    },
    // 取最大横坐标
    xMax: function () {
      let i = 0
      for (let index in this.seatList) {
        if (this.seatList[index].gCol > i) {
          i = this.seatList[index].gCol
        }
      }
      return i
    },
    // 取最大纵坐标
    yMax: function () {
      let i = 0
      for (let index in this.seatList) {
        if (this.seatList[index].gRow > i) {
          i = this.seatList[index].gRow
        }
      }
      return i
    },
    // 竖中轴线
    middleLine: function () {
      return ((this.xMax / 2) + 1) * this.positionDistin - 0.025
    },
    // 横中轴线
    horizontalLine: function () {
      return ((this.yMax / 2) + 1) * this.positionDistin - 0.025
    },
    // 根据影厅的大小缩放比例(需要把影厅全部显示出来)
    seatScale: function () {
      let seatScaleX = 1
      let seatScaleY = 1
      seatScaleX = this.seatAreaWidthRem / this.seatBoxWidth
      seatScaleY = this.seatAreaHeightRem / this.seatBoxHeight
      return seatScaleX < seatScaleY ? seatScaleX : seatScaleY
    },
    // 让影厅居中展示的偏移值
    seatBoxCenterMargin: function () {
      return -(this.seatBoxWidth * this.seatScale) / 2
    },
    // class 为 seatBox 的高度值 单位为rem
    seatBoxHeight: function () {
      return (this.yMax + 1) * this.positionDistin + this.height
    },
    // class 为 seatBox 的宽度值 单位为rem
    seatBoxWidth: function () {
      return (this.xMax + 1) * this.positionDistin + this.width
    },
    // 缩略图宽 rem
    thumbnailBoxWidth: function () {
      return ((this.xMax + 1) * this.thumbnailPositionDistin + this.thumbnailWidth)
    },
    // 缩略图高 rem
    thumbnailBoxHeight: function () {
      return ((this.yMax + 1) * this.thumbnailPositionDistin + this.thumbnailHeight)
    },
    // 座位左边栏的数组
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
</style>
