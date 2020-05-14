<template>
  <div class="container">
    <form @submit.prevent="login">
      <div class="form-group">
        <label for="email">メールアドレス</label>
        <input type="text" id="email" class="form-control" v-model="email">
      </div>
      <div class="form-group">
        <label for="password">パスワード</label>
        <input type="password" id="password" class="form-control" v-model="password">
      </div>
      <button type="submit">ログイン</button>
    </form>
  </div>
</template>

<script lang="ts">
export default {
  data() {
    return {
      email: null,
      password: null
    }
  },
  methods: {
    login: async function (this: any) {
      await this.$store.dispatch('login', {
        email: this.email,
        password: this.password
      })

      this.$nuxt.$emit('signedIn')
    }
  }
}
</script>
