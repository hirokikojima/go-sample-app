import Login from "~/types/models/Login"
import { EvalSourceMapDevToolPlugin } from "webpack"

interface IndexInterface {
    login: Login | null
}

export const state = (): IndexInterface => ({
    login: null
})

export const actions = {
    async register(this: any, { commit }: any, { email, password }: {email: string, password: string}) {
        const login = await this.$axios.$post('/users', {
            email: email,
            password: password
        })
        commit('setLogin', login)
    },
    async login(this: any, { commit }: any, { email, password }: {email: string, password: string}) {
        const login = await this.$axios.$get('/users', {
            email: email,
            password: password
        })
        commit('setLogin', login)
    }
}

export const mutations = {
    setLogin(state: IndexInterface, login: Login): void {
        state.login = login
    }
}