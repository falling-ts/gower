import msg from "./msg"
import http from "./http"
import cookie from "js-cookie"
import * as store from "localforage"

interface Res {
    code: number
    msg: string
    data?: {} | undefined
    token: string
}

interface Client {
    baseUri: string
    option: RequestInit
    get(path: string, data: object, option: RequestInit): Promise<Res>
    post(path: string, data: object, option: RequestInit): Promise<Res>
    put(path: string, data: object, option: RequestInit): Promise<Res>
    del(path: string, data: object, option: RequestInit): Promise<Res>
}

class client implements Client {
    baseUri: string
    option: RequestInit
    constructor() {
        const headers = new Headers({
            Accept: "application/json",
        })
        this.option = {
            headers,
            credentials: "include",
        }
        this.baseUri = import.meta.env.VITE_APP_URL
    }
    setOption(option: RequestInit = {}) {
        const headers = new Headers({
            ...this.option.headers,
            ...option.headers
        })

        delete option.headers
        delete this.option.headers
        this.option = {
            headers,
            ...this.option,
            ...option
        }
    }
    async get(path: string, data: object = {}, option: RequestInit = {}): Promise<Res> {
        try {
            let url = `${this.baseUri}${path}`
            if (!isEmptyObject(data)) {
                url += `?${toQueryString(data)}`
            }

            await this._setToken()

            if (!isEmptyObject(option)) {
                this.setOption(option)
            }
            const response: Response = await fetch(url, this.option)
            const res: Res = await response.json()

            if (res.code) {
                msg.error(res.msg)
                console.log('Error : ', res)
                await unauthorized(res, path)
            }

            await handleToken(res)
            return res
        } catch (error) {
            msg.error(error as string)
            console.log('Error : ', error)
            throw error
        }
    }
    async post(path: string, data: object, option: RequestInit = {}): Promise<Res> {
        try {
            await this._setToken()
            const body: BodyInit = await this._body(data)

            if (!isEmptyObject(option)) {
                this.setOption(option)
            }
            const response: Response = await fetch(`${this.baseUri}${path}`, {
                ...this.option,
                method: "POST",
                body
            })
            const res: Res = await response.json()

            if (res.code) {
                msg.error(res.msg)
                console.log('Error : ', res)
                await unauthorized(res, path)
            }

            await handleToken(res)
            return res
        } catch (error) {
            msg.error(error as string)
            console.log('Error : ', error)
            throw error
        }
    }
    async put(path: string, data: object, option: RequestInit = {}): Promise<Res> {
        try {

            await this._setToken()
            const body: BodyInit = await this._body(data)
            if (!isEmptyObject(option)) {
                this.setOption(option)
            }

            const response: Response = await fetch(`${this.baseUri}${path}`, {
                ...this.option,
                method: "PUT",
                body
            })
            const res: Res = await response.json()

            if (res.code) {
                msg.error(res.msg)
                console.log('Error : ', res)
                await unauthorized(res, path)
            }

            await handleToken(res)
            return res
        } catch (error) {
            msg.error(error as string)
            console.log('Error : ', error)
            throw error
        }
    }
    async del(path: string, data: object = {}, option: RequestInit = {}): Promise<Res> {
        try {

            await this._setToken()
            const body: BodyInit = await this._body(data)
            if (!isEmptyObject(option)) {
                this.setOption(option)
            }

            const response: Response = await fetch(`${this.baseUri}${path}`, {
                ...this.option,
                method: "DELETE",
                body
            })
            const res: Res = await response.json()

            if (res.code) {
                msg.error(res.msg)
                console.log('Error : ', res)
                await unauthorized(res, path)
            }

            await handleToken(res)
            return res
        } catch (error) {
            msg.error(error as string)
            console.log('Error : ', error)
            throw error
        }
    }
    async _setToken() {
        let auth: string | null = await store.getItem("auth")
        if (auth === null) {
            auth = ""
        }

        this.setOption({
            headers: {
                Authorization: auth
            }
        })
    }
    async _body(data: object): Promise<BodyInit> {
        if (data instanceof FormData) {
            return data
        }

        this.setOption({
            headers: {
                "Content-Type": "application/json"
            }
        })
        return JSON.stringify(data)
    }
}

function isEmptyObject(obj: Record<string, any>): boolean {
    return Object.keys(obj).length === 0;
}

function toQueryString(params: Record<string, any>): string {
    return Object.entries(params)
        .map(([key, value]) => `${encodeURIComponent(key)}=${encodeURIComponent(value)}`)
        .join("&");
}

async function unauthorized(res: Res, path: string) {
    if (res.code == http.Unauthorized) {
        await cookie.remove("token")
        await cookie.remove("admin-token")
        await store.removeItem("auth")

        setTimeout(() => {
            const admin = new RegExp(`^\/admin`)
            if (admin.test(path)) {
                window.location.href = "/admin/auth/login"
                return
            }
            window.location.href = "/auth/login"
        }, 2000)
    }
}

async function handleToken(res: Res) {
    if (res.token !== "") {
        await store.setItem("auth", res.token)
    }
}

export default new client()
