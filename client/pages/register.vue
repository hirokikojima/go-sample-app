<template>
  <div class="container">
    <form @submit.prevent="register">
      <div class="form-group">
        <label for="email">メールアドレス</label>
        <input type="text" id="email" class="form-control">
      </div>
      <div class="form-group">
        <label for="password">パスワード</label>
        <input type="password" id="password" class="form-control">
      </div>
      <button type="submit">登録</button>
    </form>
  </div>
</template>

<script lang="ts">
export default {
  date() {
    return {
      email: null,
      password: null
    }
  },
  methods: {
    register: function (this: any) {
      this.$store.dispatch('register', {
        email: this.email,
        password: this.password
      })
    }
  }
}
</script>
