import { EvalSourceMapDevToolPlugin } from "webpack"

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
    }
}