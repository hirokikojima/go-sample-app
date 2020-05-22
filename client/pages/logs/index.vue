<template>
  <div class="container">
    <div class="row">
      <div class="col-md-12 section">
        <nav aria-label="breadcrumb">
          <ol class="breadcrumb">
            <li class="breadcrumb-item"><nuxt-link to="/">Home</nuxt-link></li>
            <li class="breadcrumb-item"><nuxt-link to="/user">User</nuxt-link></li>
            <li class="breadcrumb-item"><nuxt-link to="/service">Service</nuxt-link></li>
            <li class="breadcrumb-item active" aria-current="page">Log</li>
          </ol>
        </nav>
      </div>
    </div>
    <div class="row">
      <div class="col-md-12">
        <div class="text-right">
          <button class="btn btn-primary">いいね</button>
        </div>
        <div class="section-helf">
          <h4>サービス名</h4>
          <p>
            サンプルテキストサンプルテキストサンプルテキストサンプルテキスト<br/>
            サンプルテキストサンプルテキストサンプルテキストサンプルテキスト<br/>
            サンプルテキストサンプルテキストサンプルテキストサンプルテキスト
          </p>
        </div>
      </div>
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
</template>

<script lang="ts">
import consola from 'consola'
export default {
  mounted: function(this: any) {
    consola.log(this.$route)
  }
}
</script>