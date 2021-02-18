<template>
  <div class="hello">
    <h1>欢迎宝宝们进行选座 ^.^</h1>
    <p style = "font-size:20px;">请输入您最心仪的座位号:</p>
    <input type = "text" v-model = "site_num">
    <button v-on:click="greet" style = "font-size:13px;">生成心动的二维码</button>
    <button v-on:click="clean" style = "font-size:13px;">清空</button>
    <p style = "font-size:18px;">{{result}}</p>
    <vue-qr :logoSrc="imageUrl" :text="qr_value" :size="200" v-show="code_show == 1"></vue-qr>
  </div>
</template>

<script>
import vueQr from 'vue-qr'
import axios from 'axios'
export default {
  name: 'HelloWorld',
  data () {
    return {
      site_num: 'WS02.',
      result: '未生成有效二维码',
      response_body: '',
      ip_valid: 'false',
      qr_value: '',
      code_show: 0,
      encrypt_data: 'testme',
      seat_num: '',
      imageUrl: require('../assets/vxrail.png')
    }
  },
  components: {
    vueQr
  },
  methods: {
    greet: function (event) {
      if (event) {
        if (this.site_num.length === 10 && this.site_num < 'WS02.02233' && this.site_num > 'WS02.02004') {
          this.site_string = this.site_num.slice(5)
          this.seat_num = {
            'seat_number': this.site_string
          }
          axios({
            url: '/api/powercubicle/v1/seat/encrypt',
            method: 'POST',
            data: JSON.stringify(this.seat_num),
            headers: {
              'Content-Type': 'application/json'
            }
          }).then((res) => {
            console.log(res.data)
            this.encrypt_data = res.data
            console.log(this.encrypt_data)
            console.log(JSON.stringify(this.seat_num))
          }).catch((error) => {
            console.log(error)
            this.result = '未生成有效二维码'
            this.code_show = 0
            alert('cannot obtain code string')
          })
        } else {
          this.result = '未生成有效二维码'
          this.code_show = 0
          this.qr_value = 'test'
          alert('Please input 10 valid characters! [WS02.02005 ~ WS02.02232]')
        }
      }
    },
    clean: function (event) {
      if (event) {
        this.site_num = 'WS02.'
        this.result = '未生成有效二维码'
        this.code_show = 0
        this.qr_value = 'test'
      }
    }
  },
  watch: {
    encrypt_data () {
      this.response_body = this.encrypt_data.encrypt_text
      this.ip_valid = this.encrypt_data.msg
      if (this.ip_valid.indexOf('true') !== -1) {
        this.qr_value = 'http://dellemc.evenhidata.com:8080/?encryptCode=' + this.response_body
        this.result = '已生成可扫描二维码'
        console.log(this.qr_value)
        this.code_show = 1
      } else {
        this.result = '未生成有效二维码'
        this.code_show = 0
        alert('please arrive at the office building')
      }
    }
  }
}
</script>
