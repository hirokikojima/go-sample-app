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
            <label for="title">サービス名</label>
            <input type="text" id="title" class="form-control" v-model="title">
          </div>
          <div class="form-group">
            <label for="thumbnail">サムネイル</label>
            <div class="custom-file">
              <input type="file" id="thumbnail" class="custom-file-input" accept="image/jpeg, image/png" @change="onImageChange">
              <label class="custom-file-label" for="thumbnail">Choose file</label>
            </div>
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
  data() {
    return {
      title: null,
      thumbnail: null,
      body: null
    }
  },
  methods: {
    onImageChange: function(this: any, e: any) {
      const image = e.target.files || e.dataTransfer.files
    },
    submit: async function (this: any) {
      await this.$store.dispatch('service/create', {
        title: this.title,
        thumbnail: this.thumbnail,
        body: this.body
      })
    }
  }
}
</script>
