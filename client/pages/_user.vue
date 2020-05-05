<template>
  <div>
    <div class="container" v-if="$route.params.service === undefined">
      <div class="row">
        <div class="col-md-12 section">
          <nav aria-label="breadcrumb">
            <ol class="breadcrumb">
              <li class="breadcrumb-item"><nuxt-link to="/">Home</nuxt-link></li>
              <li class="breadcrumb-item active" aria-current="page">User</li>
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
          <h4 class="list-title">サービス</h4>
          <div class="card section-helf">
            <div class="card-body">
              <nuxt-link tag="a" to="/user/service">
                <div class="list-item-service" style="margin-bottom: 0.5rem;">
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

              <ul class="list-group">
                <nuxt-link class="list-group-item" tag="li" to="/user/service/log">
                  <div class="list-item-log">
                    <div class="body">
                      <h4 class="title">ログ名</h4>
                      <p class="description">
                        <span class="date">2020-05-01</span>
                      </p>
                    </div>
                  </div>
                </nuxt-link>
                <nuxt-link class="list-group-item" tag="li" to="/user/service/log">
                  <div class="list-item-log">
                    <div class="body">
                      <h4 class="title">ログ名</h4>
                      <p class="description text-right">
                        <span class="date">2020-05-01</span>
                      </p>
                    </div>
                  </div>
                </nuxt-link>
              </ul>
            </div>
          </div>

          <div class="card section-helf">
            <div class="card-body">
              <nuxt-link tag="a" to="/user/service">
                <div class="list-item-service" style="margin-bottom: 0.5rem;">
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

              <ul class="list-group">
                <nuxt-link class="list-group-item" tag="li" to="/user/service/log">
                  <div class="list-item-log">
                    <div class="body">
                      <h4 class="title">ログ名</h4>
                      <p class="description">
                        <span class="date">2020-05-01</span>
                      </p>
                    </div>
                  </div>
                </nuxt-link>
                <nuxt-link class="list-group-item" tag="li" to="/user/service/log">
                  <div class="list-item-log">
                    <div class="body">
                      <h4 class="title">ログ名</h4>
                      <p class="description text-right">
                        <span class="date">2020-05-01</span>
                      </p>
                    </div>
                  </div>
                </nuxt-link>
              </ul>
            </div>
          </div>
        </div>
      </div>
    </div>

    <!-- 
      nuxtでネストしたルートへルーティングできない
      $route.nameが一致した場合以外は子ページへ流す
     -->
    <nuxt-child v-else />
  </div>
</template>

<script lang="ts">
import consola from 'consola'
export default {
  mounted: function(this: any) {
    consola.log(this.$route)
  }
}
</script>