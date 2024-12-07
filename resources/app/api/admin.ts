import client from "../lib/client";

interface Data {
    username: string
    password: string
}

async function login(data: Data, option: RequestInit = {}) {
    return await client.post("/admin/auth/login", data, option)
}

async function logout(data: Data | object = {}, option: RequestInit = {}) {
    return await client.post("/admin/auth/logout", data, option)
}

async function image(data: FormData, option: RequestInit = {}) {
    return await client.post("/admin/upload/image", data, option)
}

export default {
    login,
    logout,
    image
}
