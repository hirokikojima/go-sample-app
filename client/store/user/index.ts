import User from "~/types/models/User"
import { EvalSourceMapDevToolPlugin } from "webpack"

interface UserInterface {
    user: User | null
}

export const state = (): UserInterface => ({
    user: null
})

export const actions = {
    async getUser(this: any, {commit}: any, { userId }: { userId: number}) {
        const response = await this.$axios.$get('/user/' + userId)
        
        commit('setUser', response)
    }
}

export const mutations = {
    setUser(state: UserInterface, user: User): void {
        state.user = user
    }
}