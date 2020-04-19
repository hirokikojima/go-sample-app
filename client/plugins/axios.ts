export let axios: any

export default ({ store, $axios }: { store: any, $axios: any }) => {
  $axios.onRequest((config: any) => {
    // config.headers.common['Authorization'] = `Bearer token`
    // config.headers.common['Accept'] = 'application/json'
  })

  $axios.onResponse((response: any) => {
    return Promise.resolve(response);
  })

  $axios.onError((error: any) => {
    return Promise.reject(error.response);
  })

  axios = $axios;
}