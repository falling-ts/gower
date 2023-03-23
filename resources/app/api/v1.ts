import client from "../lib/client";

interface Data {
    test: string
    name: string
}

async function hello(data: Data, option: RequestInit = {}) {
    return await client.get("/api/v1/hello", data)
}

export default {
    hello
}
