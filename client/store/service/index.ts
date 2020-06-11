import Service from "~/types/models/Service"
import { EvalSourceMapDevToolPlugin } from "webpack"

interface ServiceInterface {
    service: Service | null
}

export const state = (): ServiceInterface => ({
    service: null
})

export const actions = {
    async create(this: any, { commit }: any, { title, thumbnail, body }: {title: string, thumbnail: File, body: string}) {
        let formData = new FormData
        formData.append('title', title)
        formData.append('thumbnail', thumbnail)
        formData.append('body', body)
        
        const response = await this.$axios.$post('/service', formData, {
            headers: { 'Content-Type': 'multipart/form-data' }
        })

        return response
    },
    async getServices(this: any, { commit }: any, { userId }: { userId: number }) {
        const response = await this.$axios.$get('/service/' + userId)
    }
}