<template>
  <div class="wrapper">
    <header>
      <Header />
    </header>
    <main>
      <nuxt />
    </main>
    <footer>
      <Footer />
    </footer>
  </div>
</template>

<script lang="ts">
import Header from '~/components/common/Header.vue'
import Footer from '~/components/common/Footer.vue'

export default ({
  components: {
    Header,
    Footer
  }
})
</script>