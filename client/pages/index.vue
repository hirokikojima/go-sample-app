<template>
  <div class="container">
    <div class="row">
      <div class="col-md-4 order-md-2 section">
        <div>
          <h4 class="list-title">人気サービスランキング</h4>
          <ul class="list-group">
            <nuxt-link class="list-group-item" tag="li" to="/user/service">
              <div class="list-item-service">
                <div class="icon">
                  <img src="~/assets/image/no_image.png" width="62" height="62" />
                </div>
                <div class="body">
                  <h4 class="title">サービス名</h4>
                  <p class="description">
                    <span class="date">2020-05-01</span>
                    <span class="user-name">ユーザ名</span>
                  </p>
                </div>
              </div>
            </nuxt-link>
            <nuxt-link class="list-group-item" tag="li" to="/user/service">
              <div class="list-item-service">
                <div class="icon">
                  <img src="~/assets/image/no_image.png" width="62" height="62" />
                </div>
                <div class="body">
                  <h4 class="title">サービス名</h4>
                  <p class="description">
                    <span class="date">2020-05-01</span>
                    <span class="user-name">ユーザ名</span>
                  </p>
                </div>
              </div>
            </nuxt-link>
          </ul>
        </div>
        <div class="section-helf">
          <h4 class="list-title">アクティブユーザランキング</h4>
          <ul class="list-group">
            <nuxt-link class="list-group-item" tag="li" to="/user">
              <div class="list-item-user">
                <div class="icon">
                  <img src="~/assets/image/no_image.png" width="62" height="62" />
                </div>
                <div class="body">
                  <h4 class="title">ユーザ名</h4>
                </div>
              </div>
            </nuxt-link>
          </ul>
        </div>
      </div>
      <div class="col-md-8 order-md-1 section">
        <ServiceList title="タイムライン" :services="services" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import ServiceList from '~/components/service/List.vue'
export default({
  components: {
    ServiceList
  },
  data() {
    return {
      user: null,
      services: null
    }
  },
  async mounted(this: any) {
    this.services = await this.$api.getServices()
  }
})
</script>