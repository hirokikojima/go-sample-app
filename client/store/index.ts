import User from "~/types/models/User"
import { EvalSourceMapDevToolPlugin } from "webpack"

interface IndexInterface {
    me: User | null
}

export const state = (): IndexInterface => ({
    me: null
})

export const actions = {
    async signup(this: any, { commit }: any, { name, email, password }: {name: string, email: string, password: string}) {
        const response = await this.$axios.$post('/signup', {
            name: name,
            email: email,
            password: password
        })
        this.$auth.signin(response.token)
    },
    async login(this: any, { commit }: any, { email, password }: {email: string, password: string}) {
        const response = await this.$axios.$post('/login', {
            email: email,
            password: password
        })
        this.$auth.signin(response.token)
    },
    async logout(this: any, { commit }: any) {
        this.$auth.signout()
    },
    async me(this: any, { commit }: any) {
        const user = await this.$axios.$get('/me')

        commit('setMe', user)
    }
}

export const mutations = {
    setMe(state: IndexInterface, me: User): void {
        state.me = me
    }
}