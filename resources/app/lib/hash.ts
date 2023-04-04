import JsSHA from "jssha"

export default {
    SHA256(str: string): string {
        const shaObj = new JsSHA(
            "SHA-256",
            "TEXT",
            { encoding: "UTF8" }
        )
        shaObj.update(str)
        return shaObj.getHash("HEX")
    }
}
