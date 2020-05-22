
export default {
  mode: 'spa',
  /*
  ** Headers of the page
  */
  head: {
    title: process.env.npm_package_name || '',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: process.env.npm_package_description || '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress-bar color
  */
  loading: { color: '#fff' },
  /*
  ** Global CSS
  */
  css: [
    '@/assets/css/style.css'
  ],
  /*
  ** Plugins to load before mounting the App
  */
  plugins: [
    '@/plugins/auth.ts',
    '@/plugins/axios.ts'
  ],
  /*
  ** Nuxt.js dev-modules
  */
  buildModules: [
    '@nuxt/typescript-build',
  ],
  /*
  ** Nuxt.js modules
  */
  modules: [
    // Doc: https://bootstrap-vue.js.org
    'bootstrap-vue/nuxt',
    '@nuxtjs/axios',
    '@nuxtjs/dotenv',
    '@nuxtjs/markdownit',
  ],
  /*
  ** axios options
  */
  axios: {
    withCredentials: true,
    baseURL: 'http://localhost:1323'
  },
  /**
  * markdownit options
  */
  markdownit: {
    injected: true
  },
  /*
  ** Build configuration
  */
  build: {
    /*
    ** You can extend webpack config here
    */
    extend (config, ctx) {
    }
  },
  /**
  * Router configuration
  */
  router: {
    extendRoutes (routes, resolve) {
      /**
       * log editor
       */
      routes.push({
        path: '/users/:id/services/:id/logs/:id/editor',
        component: resolve(__dirname, 'pages/logs/editor.vue')
      })

      /**
       * log index
       */
      routes.push({
        path: '/users/:id/services/:id/logs/:id',
        component: resolve(__dirname, 'pages/logs/index.vue')
      })

      /**
       * service editor
       */
      routes.push({
        path: '/users/:id/services/:id/editor',
        component: resolve(__dirname, 'pages/services/editor.vue')
      })

      /**
       * service index
       */
      routes.push({
        path: '/users/:id/services/:id',
        component: resolve(__dirname, 'pages/services/index.vue')
      })

      /**
       * user editor
       */
      routes.push({
        path: '/users/:id/editor',
        component: resolve(__dirname, 'pages/users/editor.vue')
      })

      /**
       * user index
       */
      routes.push({
        path: '/users/:id',
        component: resolve(__dirname, 'pages/users/index.vue')
      })
    }
  }
}
