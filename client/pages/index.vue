<template>
  <div class="container">
    <div class="row">
      <div class="col-md-4 order-md-2 section">
        <h4 class="list-title">人気サービスランキング</h4>
        <ul class="list-group">
          <nuxt-link class="list-group-item" tag="li" to="/user/service">
            <div class="row service">
              <div class="col-md-4 icon">
                <img src="~/assets/image/no_image.png" width="62" height="62" />
              </div>
              <div class="col-md-8">
                <h4 class="title">サービス名</h4>
                <p class="description">
                  <span class="date">2020-05-01</span>
                  <span class="user-name">ユーザ名</span>
                </p>
              </div>
            </div>
          </nuxt-link>
          <nuxt-link class="list-group-item" tag="li" to="/user/service">
            <div class="row service">
              <div class="col-md-4 icon">
                <img src="~/assets/image/no_image.png" width="62" height="62" />
              </div>
              <div class="col-md-8">
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
      <div class="col-md-8 order-md-1 section">
        <h4 class="list-title">新着情報</h4>
        <ul class="list-group">
          <nuxt-link class="list-group-item" tag="li" to="/user/service/log">
            <div class="row log">
              <div class="col-md-2 icon">
                <img src="~/assets/image/no_image.png" width="62" height="62" />
              </div>
              <div class="col-md=10">
                <h4 class="title">ログ名</h4>
                <p class="description">
                  <span class="date">2020-05-01</span>
                  <span class="service-name">サービス名</span>
                  <span class="user-name">ユーザ名</span>
                </p>
              </div>
            </div>
          </nuxt-link>
          <nuxt-link class="list-group-item" tag="li" to="/user/service/log">
            <div class="row log">
              <div class="col-md-2 icon">
                <img src="~/assets/image/no_image.png" width="62" height="62" />
              </div>
              <div class="col-md=10">
                <h4 class="title">ログ名</h4>
                <p class="description">
                  <span class="date">2020-05-01</span>
                  <span class="service-name">サービス名</span>
                  <span class="user-name">ユーザ名</span>
                </p>
              </div>
            </div>
          </nuxt-link>
          <nuxt-link class="list-group-item" tag="li" to="/user/service/log">
            <div class="row log">
              <div class="col-md-2 icon">
                <img src="~/assets/image/no_image.png" width="62" height="62" />
              </div>
              <div class="col-md=10">
                <h4 class="title">ログ名</h4>
                <p class="description">
                  <span class="date">2020-05-01</span>
                  <span class="service-name">サービス名</span>
                  <span class="user-name">ユーザ名</span>
                </p>
              </div>
            </div>
          </nuxt-link>
        </ul>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
})
</script>