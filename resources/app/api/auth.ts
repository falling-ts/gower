import client from "../lib/client";

interface Data {
    username: string
    password: string
}

async function register(data: Data, option: RequestInit = {}) {
    return await client.post("/auth/register", data, option)
}

async function login(data: Data, option: RequestInit = {}) {
    return await client.post("/auth/login", data, option)
}

async function logout(data: Data | object = {}, option: RequestInit = {}) {
    return await client.post("/auth/logout", data, option)
}

export default {
    register,
    login,
    logout
}
