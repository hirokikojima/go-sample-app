<template>
  <div>
    <div class="container" v-if="$route.params.log === undefined">
      <div class="row">
        <div class="col-md-12 section">
          <nav aria-label="breadcrumb">
            <ol class="breadcrumb">
              <li class="breadcrumb-item"><nuxt-link to="/">Home</nuxt-link></li>
              <li class="breadcrumb-item"><nuxt-link to="/user">User</nuxt-link></li>
              <li class="breadcrumb-item active" aria-current="page">Service</li>
            </ol>
          </nav>
        </div>
      </div>
      <div class="row">
        <div class="col-md-12">
          <div class="text-right">
            <button class="btn btn-primary">いいね</button>
            <button class="btn btn-primary">フォロー</button>
          </div>
          <div class="section-helf text-center bg-darkgray">
            <img src="~/assets/image/no_image.png" height="240px" style="background: #fff">
          </div>
          <div class="section-helf">
            <h4>サービス名</h4>
            <p>
              サービス説明テキストサービス説明テキストサービス説明テキストサービス説明テキストサービス説明テキスト<br/>
              サービス説明テキストサービス説明テキストサービス説明テキストサービス説明テキストサービス説明テキスト<br/>
              サービス説明テキストサービス説明テキストサービス説明テキストサービス説明テキストサービス説明テキスト
            </p>
            <p><a href="http://example.com/">http://example.com/</a></p>
          </div>
          <div class="section-helf">
            <h4 class="list-title">更新ログ</h4>
            <ul class="list-group">
              <nuxt-link class="list-group-item" tag="li" to="/user/service/log">
                <div class="list-item-log">
                  <div class="icon">
                    <img src="~/assets/image/no_image.png" width="62" height="62" />
                  </div>
                  <div class="body">
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
                <div class="list-item-log">
                  <div class="icon">
                    <img src="~/assets/image/no_image.png" width="62" height="62" />
                  </div>
                  <div class="body">
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
                <div class="list-item-log">
                  <div class="icon">
                    <img src="~/assets/image/no_image.png" width="62" height="62" />
                  </div>
                  <div class="body">
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
          <div class="section-helf">
      <h4 class="list-title">コメント</h4>
      <ul class="list-group">
        <nuxt-link class="list-group-item" tag="li" to="/user/service/log">
          <div class="list-item-comment">
            <div class="header">
              <div class="flex-left">
                <img class="user-icon" src="~/assets/image/no_image.png" width="62" height="62" />
                <span class="user-name">ユーザ名</span>
              </div>
              <div class="flex-right">
                <span class="date">2020-05-01</span>
              </div>
            </div>
            <div class="body">
              <p class="comment">
                コメントテキストコメントテキストコメントテキストコメントテキストコメントテキスト<br/>
                コメントテキストコメントテキストコメントテキストコメントテキストコメントテキスト<br/>
                コメントテキストコメントテキストコメントテキストコメントテキストコメントテキスト
              </p>
            </div>
          </div>
        </nuxt-link>
      </ul>
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