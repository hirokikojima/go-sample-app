<template>
  <div class="container">
    <div class="row">
      <div class="col-md-6 order-md-2 section">
        <div class="section-helf">
          <h4>{{ title }}</h4>
          <div class="description" v-html="$md.render(body || '')" />
        </div>
      </div>
      <div class="col-md-6 order-md-1 section">
        <form @submit.prevent="submit">
          <div class="form-group">
            <label for="title">タイトル</label>
            <input type="text" id="title" class="form-control" v-model="title">
          </div>
          <div class="form-group">
            <label for="body">本文</label>
            <textarea rows="10" id="body" class="form-control" v-model="body"></textarea>
          </div>
          <button type="submit" class="btn btn-outline-darkcyan btn-block">下書き</button>
          <button type="submit" class="btn btn-darkcyan btn-block">追加</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  computed: {
    userId: function(this: any) {
      return this.$route.params.user_id
    },
    serviceId: function(this: any) {
      return this.$route.params.service_id
    }
  },
  data() {
    return {
      id: null,
      title: null,
      body: null
    }
  },
  methods: {
    submit: async function (this: any) {
      this.id ? this.update() : this.create()
    },
    create: async function (this: any) {
      await this.$store.dispatch('log/create', {
        serviceId: this.serviceId,
        title: this.title,
        body: this.body
      })
    },
    update: async function (this: any) {
      await this.$store.dispatch('log/update', {
        serviceId: this.serviceId,
        title: this.title,
        body: this.body
      })
    }
  },
  mounted(this: any) {
    if (!this.$route.params.user_id || !this.$route.params.service_id) {
      // エラーページへ遷移させる
      this.$nuxt.error({
        statusCode: 500,
        message: 'ページが見つかりませんでした。'
      })
    }
  }
}
</script>
