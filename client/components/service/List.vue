<template>
  <div>
    <h4 class="list-title">
      <span>{{ title }}</span>
    </h4>
    <template v-for="service in services">
      <ListItem :key="service.ID" :service="service" :editable="editable"/>
    </template>
  </div>
</template>

<script lang="ts">
import ListItem from '~/components/service/ListItem.vue'

export default ({
  components: {
    ListItem
  },
  props: ['title', 'services', 'editable']
})
</script>