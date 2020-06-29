export const actions = {
  async create(this: any, { commit }: any, { serviceId, title, body }: { serviceId: number, title: string, body: string}) {
    return await this.$axios.$post('/log', {
      service_id: serviceId,
      title: title,
      body: body
    })
  },
  async update(this: any, { commit }: any, { id, title, body }: { id: number, title: string, body: string}) {
    return await this.$axios.$put('/log', {
      id: id,
      title: title,
      body: body
    })
  }
}