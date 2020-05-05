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