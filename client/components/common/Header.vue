<template>
  <nav class="navbar navbar-dark bg-darkcyan">
    <div class="container">
      <nuxt-link class="navbar-brand" to="/">
        Go-Sample-App
      </nuxt-link>
      <div v-if="!isSignedIn">
        <nuxt-link class="btn btn-light" to="/login">ログイン</nuxt-link>
        <nuxt-link class="btn btn-outline-light" to="/signup">ユーザ登録</nuxt-link>
      </div>
      <div v-else>
        <a class="btn btn-outline-light" @click="me">Me</a>
      </div>
    </div>
  </nav>
</template>

<script lang="ts">
export default ({
  computed: {
    isSignedIn: function(this: any) {
      return !! this.$auth.getToken()
    }
  },
  methods: {
    me: function (this: any) {
      this.$store.dispatch('me')
    }
  }
})
</script>