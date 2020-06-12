<template>
  <div class="container" v-if="user">
    <div class="row">
      <div class="col-md-12 section">
        <nav aria-label="breadcrumb">
          <ol class="breadcrumb">
            <li class="breadcrumb-item"><nuxt-link to="/">Home</nuxt-link></li>
            <li class="breadcrumb-item active" aria-current="page">{{ user.name }}</li>
          </ol>
        </nav>
      </div>
    </div>
    <div class="row">
      <div class="col-md-4 order-md-1 section">
        <div class="card profile">
          <div class="card-img-top img-area">
            <img class="img" src="~/assets/image/no_image.png">
          </div>
          <div class="card-body">
            <div class="row">
              <div class="col-md-6">
                <p class="text-center">1</p>
                <p class="text-center text-muted">フォロー</p>
              </div>
              <div class="col-md-6">
                <p class="text-center">1</p>
                <p class="text-center text-muted">フォロワー</p>
              </div>
              <div class="col-md-12">
                <nuxt-link tag="a" class="btn btn-outline-info btn-block" to="/user/profile/edit">フォローする</nuxt-link>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="col-md-8 order-md-2 section">
        <ServiceList :title="サービス" :services="services" />
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import ServiceList from '~/components/service/List.vue'

export default ({
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
    if (!this.$route.params.id) {
      // エラーページへ遷移させる
      this.$nuxt.error({
        statusCode: 500,
        message: 'ページが見つかりませんでした。'
      })
    }
      
    this.user = await this.$api.getUser(this.$route.params.id)
    this.services = await this.$api.getServicesByUserId(this.$route.params.id)
  }
})
</script>