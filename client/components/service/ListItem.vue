<template>
  <div class="card section-helf">
    <div class="card-body">
      
      <div class="list-item-service" style="margin-bottom: 0.5rem;">
        <div class="icon">
          <img src="~/assets/image/no_image.png" width="62" height="62" />
        </div>
        <div class="body">
          <div class="d-flex justify-content-between">
            <h4 class="title">
              <nuxt-link tag="a" :to="`/users/${service.UserID}/services/${service.ID}`">{{ service.Title }}</nuxt-link>
            </h4>
            <b-dropdown variant="outline-dark" size="sm" right v-if="editable">
              <b-dropdown-item :to="`/users/${service.UserID}/services/${service.ID}/edit`">編集</b-dropdown-item>
              <b-dropdown-item>削除</b-dropdown-item>
              <b-dropdown-divider></b-dropdown-divider>
              <b-dropdown-item :to="`/users/${service.UserID}/services/${service.ID}/logs/new`">ログを追加</b-dropdown-item>
            </b-dropdown>
          </div>
          <p class="description">
            <span class="date">{{ service.CreatedAt }}</span>
            <span class="user-name">
              <nuxt-link tag="a" :to="`/users/${service.UserID}`">{{ service.User.Name }}</nuxt-link>
            </span>
          </p>
        </div>
      </div>

      <ul class="list-group">
        <template v-for="log in service.Logs">
          <nuxt-link class="list-group-item" tag="li" :key="log.ID" :to="`/users/${service.UserID}/services/${service.ID}/logs/${log.ID}`">
            <div class="list-item-log">
              <div class="body">
                <h4 class="title">{{ log.Title }}</h4>
                <p class="description">
                  <span class="date">{{ log.CreatedAt }}</span>
                </p>
              </div>
            </div>
          </nuxt-link>
        </template>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
export default ({
  props: ['service', 'editable']
})
</script>