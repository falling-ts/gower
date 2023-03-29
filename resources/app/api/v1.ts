import client from "../lib/client";

interface Data {
    test: string
    name: string
}

async function hello(data: Data, option: RequestInit = {}) {
    return await client.get("/api/v1/hello", data)
}

async function image(data: FormData, option: RequestInit = {}) {
    return await client.post("/api/v1/upload/image", data)
}

export default {
    hello,
    image
}
