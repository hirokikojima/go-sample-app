<template>
  <nav class="navbar navbar-dark bg-darkcyan">
    <div class="container">
      <nuxt-link class="navbar-brand" to="/">
        Go-Sample-App
      </nuxt-link>
      <div v-if="!me">
        <nuxt-link class="btn btn-light" to="/login">ログイン</nuxt-link>
        <nuxt-link class="btn btn-outline-light" to="/signup">ユーザ登録</nuxt-link>
      </div>
      <div v-else>
        <nuxt-link class="btn btn-outline-light" to="/user">{{ me.name }}</nuxt-link>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
export default ({
  computed: {
    me: function(this: any) {
      return this.$store.state.me
    }
  }
})
</script>