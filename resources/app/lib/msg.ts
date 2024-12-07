import $ from "jquery"

interface Confirm {
    (...args: any[]): void
}

class message {
    default(msg: string) {
        alert("default", msg)
    }
    loading(msg: string) {
        return alert("loading", msg)
    }
    info(msg: string) {
        alert("info", msg)
    }
    success(msg: string) {
        alert("success", msg)
    }
    warning(msg: string) {
        alert("warning", msg)
    }
    error(msg: string) {
        alert("error", msg)
    }

    confirm(msg: string, ok: Confirm) {
        confirm(msg, ok)
    }
}

function alert(type: string, msg: string) {
    const elem = ((type, msg): HTMLElement => {
        let result = {
            "default": defaultElem(msg),
            "info": infoElem(msg),
            "success": successElem(msg),
            "warning": warningElem(msg),
            "error": errorElem(msg),
            "loading": loadingElem(msg)
        }[type]
        if (result === undefined) {
            result = defaultElem(msg)
        }

        return result
    })(type, msg), alert = $(elem)

    if (type === "loading") {
        alert.appendTo("#alert-box").addClass("animate__fadeInDown")
        const mask = $(`<div class="fixed inset-0 z-10 bg-base-300 opacity-80"></div>`)
        mask.appendTo("#app")
        return () => {
            alert.detach()
            mask.detach()
        }
    }

    alert.appendTo("#alert-box").addClass("animate__fadeInDown").on("click", () => {
        alert.addClass("animate__rotateOutDownLeft")
        setTimeout(() => {
            alert.detach()
        }, 1000)
    })

    setTimeout(() => {
        alert.addClass("animate__rotateOutDownLeft")
        setTimeout(() => {
            alert.detach()
        }, 1000)
    }, 2000)
}

function confirm(msg: string, ok: Confirm) {
    const elem = confirmElem(msg), confirm: JQuery<HTMLElement> = $(elem)

    confirm.appendTo("#alert-box").addClass("animate__fadeInDown")
    confirm.on("click", ".btn-ghost", () => {
        confirm.addClass("animate__rotateOutDownLeft")
        setTimeout(() => {
            confirm.detach()
        }, 1000)
    })
    confirm.on("click", ".btn-primary", () => {
        ok()
    })

}

function defaultElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-default shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div class="flex items-center">
        <i class="icon-[line-md--brake-abs-twotone] text-lg" role="img" aria-hidden="true"></i>
        <span class="ml-2">${msg}</span>
    </div>
    `

    return elem
}

function loadingElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-default shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div class="flex items-center">
        <i class="icon-[svg-spinners--pulse-3]" role="img" aria-hidden="true"></i>
        <span class="ml-2">${msg}</span>
    </div>
    `
    return elem
}

function infoElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-info shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div class="flex items-center">
        <i class="icon-[line-md--chat] text-lg" role="img" aria-hidden="true"></i>
        <span class="ml-2">${msg}</span>
    </div>
    `
    return elem
}

function successElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-success shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div class="flex items-center">
        <i class="icon-[line-md--confirm-square] text-lg" role="img" aria-hidden="true"></i>
        <span class="ml-2">${msg}</span>
    </div>
    `
    return elem
}

function warningElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-warning shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div class="flex items-center">
        <i class="icon-[line-md--alert-loop] text-lg" role="img" aria-hidden="true"></i>
        <span class="ml-2">${msg}</span>
    </div>
    `
    return elem
}

function errorElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-error shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div class="flex items-center">
        <i class="icon-[line-md--close-circle] text-lg" role="img" aria-hidden="true"></i>
        <span class="ml-2">${msg}</span>
    </div>
    `
    return elem
}

function confirmElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-confirm shadow-lg animate__animated mt-10 flex justify-between items-center";
    elem.innerHTML = `
    <div class="flex items-center">
        <i class="icon-[svg-spinners--tadpole]" role="img" aria-hidden="true"></i>
        <span class="ml-2">${msg}</span>
    </div>
    <div class="flex">
        <button class="btn btn-sm btn-ghost">取消</button>
        <button class="btn btn-sm btn-primary">确定</button>
    </div>
  `;
    return elem
}





export default new message()
