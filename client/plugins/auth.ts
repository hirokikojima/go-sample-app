export const tokenKey = 'jwt-token'

export class Auth {
    signin(token: string) {
        return localStorage.setItem(tokenKey, token)
    }
    signout() {
        return localStorage.removeItem(tokenKey)
    }
    getToken() {
        return localStorage.getItem(tokenKey)
    }
}

export default ({ app }: { app: any }, inject: any) => {
    const auth = new Auth()

    inject('auth', auth)
}