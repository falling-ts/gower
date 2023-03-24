import msg from "./msg"
import http from "./http"
import cookie from "js-cookie"
import * as store from "localforage"

interface Res {
    code: number
    msg: string
    data?: {} | undefined
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
            "Content-Type": "application/json",
            Authorization: "Bearer ",
        })
        this.option = {
            headers,
            credentials: "include",
        }
        this.baseUri = "http://127.0.0.1:8080"
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
            if (!isEmptyObject(option)) {
                this.setOption(option)
            }

            const response: Response = await fetch(url, this.option)
            const res: Res = await response.json()

            if (res.code) {
                msg.error(res.msg)
                console.log('Error : ', res)
                await unauthorized(res)
            }
            return res
        } catch (error) {
            msg.error(error as string)
            console.log('Error : ', error)
            throw error
        }
    }
    async post(path: string, data: object, option: RequestInit = {}): Promise<Res> {
        try {
            if (!isEmptyObject(option)) {
                this.setOption(option)
            }

            const response: Response = await fetch(`${this.baseUri}${path}`, {
                ...this.option,
                method: "POST",
                body: JSON.stringify(data)
            })
            const res: Res = await response.json()

            if (res.code) {
                msg.error(res.msg)
                console.log('Error : ', res)
                await unauthorized(res)
            }
            return res
        } catch (error) {
            msg.error(error as string)
            console.log('Error : ', error)
            throw error
        }
    }
    async put(path: string, data: object, option: RequestInit = {}): Promise<Res> {
        try {
            if (!isEmptyObject(option)) {
                this.setOption(option)
            }

            const response: Response = await fetch(`${this.baseUri}${path}`, {
                ...this.option,
                method: "PUT",
                body: JSON.stringify(data)
            })
            const res: Res = await response.json()

            if (res.code) {
                msg.error(res.msg)
                console.log('Error : ', res)
                await unauthorized(res)
            }
            return res
        } catch (error) {
            msg.error(error as string)
            console.log('Error : ', error)
            throw error
        }
    }
    async del(path: string, data: object = {}, option: RequestInit = {}): Promise<Res> {
        try {
            if (!isEmptyObject(option)) {
                this.setOption(option)
            }

            const response: Response = await fetch(`${this.baseUri}${path}`, {
                ...this.option,
                method: "DELETE",
                body: JSON.stringify(data)
            })
            const res: Res = await response.json()

            if (res.code) {
                msg.error(res.msg)
                console.log('Error : ', res)
                await unauthorized(res)
            }
            return res
        } catch (error) {
            msg.error(error as string)
            console.log('Error : ', error)
            throw error
        }
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

async function unauthorized(res: Res) {
    if (res.code == http.Unauthorized) {
        await cookie.remove("auth")
        await store.removeItem("auth")

        setTimeout(() => {
            window.location.href = "/auth/login"
        }, 2000)
    }
}

export default new client()
