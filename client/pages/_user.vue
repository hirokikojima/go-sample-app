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