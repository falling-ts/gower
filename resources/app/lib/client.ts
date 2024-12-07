import msg from "./msg"
import http from "./http"
import cookie from "js-cookie"

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

const admin = new RegExp(`^\/admin`)

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
        // @ts-ignore
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

            await this._setToken(path)

            if (!isEmptyObject(option)) {
                this.setOption(option)
            }
            const response: Response = await fetch(url, this.option)
            const res: Res = await response.json()

            if (res.code) {
                msg.error(res.msg)
                console.log("Error : ", res)
                await unauthorized(res, path)
            }

            await handleToken(res, path)
            return res
        } catch (error) {
            msg.error(error as string)
            console.log("Error : ", error)
            throw error
        }
    }
    async post(path: string, data: object, option: RequestInit = {}): Promise<Res> {
        try {
            await this._setToken(path)
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
                console.log("Error : ", res)
                await unauthorized(res, path)
            }

            await handleToken(res, path)
            return res
        } catch (error) {
            msg.error(error as string)
            console.log("Error : ", error)
            throw error
        }
    }
    async put(path: string, data: object, option: RequestInit = {}): Promise<Res> {
        try {

            await this._setToken(path)
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
                console.log("Error : ", res)
                await unauthorized(res, path)
            }

            await handleToken(res, path)
            return res
        } catch (error) {
            msg.error(error as string)
            console.log("Error : ", error)
            throw error
        }
    }
    async del(path: string, data: object = {}, option: RequestInit = {}): Promise<Res> {
        try {

            await this._setToken(path)
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
                console.log("Error : ", res)
                await unauthorized(res, path)
            }

            await handleToken(res, path)
            return res
        } catch (error) {
            msg.error(error as string)
            console.log("Error : ", error)
            throw error
        }
    }
    async _setToken(path: string) {
        if (admin.test(path)) {
            return this.setOption({
                headers: {
                    "Admin-Authorization": localStorage.getItem("admin-auth") || ""
                }
            })
        }

        this.setOption({
            headers: {
                Authorization: localStorage.getItem("auth") || ""
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
        setTimeout(() => {
            if (admin.test(path)) {
                cookie.remove("admin-token")
                localStorage.removeItem("admin-auth")
                window.location.href = "/admin/auth/login"
                return
            }

            cookie.remove("token")
            localStorage.removeItem("auth")
            window.location.href = "/auth/login"
        }, 1000)
    }
}

async function handleToken(res: Res, path: string) {
    if (!res.token) {
        return
    }
    if (admin.test(path)) {
        return localStorage.setItem("admin-auth", res.token)
    }
    localStorage.setItem("auth", res.token)
}

export default new client()
