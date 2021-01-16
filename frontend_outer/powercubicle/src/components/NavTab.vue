<template>
  <div class="footer">
      <footer class="fixBottomBox">
        <ul>
            <router-link tag="li" to="/" class="tabItem">
              <a href="javascript:;" class="tab-item-link" v-if="is_shops_index">
                <img src="../assets/images/index-norm.png" alt="" class="tabbar-logo">
                <p class="tabbar-text" style="color: rgba(234, 49, 6, 0.66);">首页</p>
              </a>
              <a href="javascript:;" class="tab-item-link" else>
                <img src="../assets/images/index-norm.png" alt="" class="tabbar-logo">
                <p class="tabbar-text">首页</p>
              </a>
            </router-link>
            <router-link tag="li" to="/shops/category" class="tabItem">
                <a href="javascript:;" class="tab-item-link" v-if="is_category">
                    <img src="../assets/images/user-info-norm.png" alt="" class="tabbar-logo">
                    <p class="tabbar-text" style="color: rgba(234, 49, 6, 0.66);">分类</p>
                </a>
                <a href="javascript:;" class="tab-item-link" else>
                  <img src="../assets/images/user-info-norm.png" alt="" class="tabbar-logo">
                  <p class="tabbar-text">分类</p>
                </a>
            </router-link>
        </ul>
    </footer>
  </div>
</template>
<script>
export default{
  data () {
    return {
    }
  },
  props: {
    is_shops_index: Boolean,
    is_category: Boolean,
    is_cart: Boolean,
    is_mine: Boolean,
  },
  mounted () {
  },
  computed: {
  },
  methods: {
  }
}
</script>