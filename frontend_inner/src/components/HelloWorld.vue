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
export default {
  name: 'HelloWorld',
  data () {
    return {
      site_num: 0,
      result: '未生成有效二维码',
      qr_value: 'test',
      code_show: 0,
      encrypt_data: '',
      imageUrl: require('../assets/vxrail.png')
    }
  },
  components: {
    vueQr
  },
  methods: {
    greet: function (event) {
      if (event) {
        if (this.site_num.length === 5) {
          this.site_string = 'WS02.' + this.site_num
          this.$http.post('http://127.0.0.1:12076/powercubicle/v1/seat/encrypt', {'seat_number': this.site_string}, {emulateJSON: true}).then(function (res) {
            this.encrypt_data = res.body.msg
            console.log(this.encrypt_data)
          }, function (res) {
            console.log(res.status)
          })
          alert(this.encrypt_data)
          this.result = '已生成可扫描二维码'
          this.qr_value = this.encrypt_data
          this.code_show = 1
        } else {
          this.result = '未生成有效二维码'
          this.code_show = 0
          this.qr_value = 'test'
          alert('Please input 5 characters!')
        }
      }
    },
    clean: function (event) {
      if (event) {
        this.site_num = ''
        this.result = '未生成有效二维码'
        this.code_show = 0
        this.qr_value = 'test'
      }
    }
  },
  watch: {
    encrypt_data () {
      if (this.encrypt_data === '') {
        return
      }
      alert(this.encrypt_data)
      this.encrypt_data = ''
    }
  }
}
</script>
