import $ from "jquery"

interface Confirm {
    (...args: any[]): void
}

class message {
    default(msg: string) {
        alert("default", msg)
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

function alert(type, msg: string) {
    const elem = ((type, msg: string) => {
        return {
            "default": defaultElem(msg),
            "info": infoElem(msg),
            "success": successElem(msg),
            "warning": warningElem(msg),
            "error": errorElem(msg)
        }[type]
    })(type, msg), alert = $(elem)

    alert.appendTo('#alert-box').addClass("animate__fadeInDown").on("click", () => {
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
    const elem = confirmElem(msg), confirm: $ = $(elem)

    confirm.appendTo('#alert-box').addClass("animate__fadeInDown")
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
    <div>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info flex-shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
        <span>${msg}</span>
    </div>
    `

    return elem
}

function infoElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-info shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-current flex-shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
        <span>${msg}</span>
    </div>
    `
    return elem
}

function successElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-success shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div>
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current flex-shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span>${msg}</span>
    </div>
    `
    return elem
}

function warningElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-warning shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div>
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current flex-shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
        <span>${msg}</span>
    </div>
    `
    return elem
}

function errorElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-error shadow-lg animate__animated mt-10 flex justify-center"
    elem.innerHTML = `
    <div>
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current flex-shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m7-2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span>${msg}</span>
    </div>
    `
    return elem
}

function confirmElem(msg: string): HTMLElement {
    const elem = document.createElement("div")
    elem.className = "w-full h-20 alert alert-confirm shadow-lg animate__animated mt-10";
    elem.innerHTML = `
    <div>
        <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info flex-shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
        <span>${msg}</span>
    </div>
    <div class="flex-none">
        <button class="btn btn-sm btn-ghost">取消</button>
        <button class="btn btn-sm btn-primary">确定</button>
    </div>
  `;
    return elem
}

// function detailElem(title, msg: string): HTMLElement {
//     const elem = document.createElement("div")
//     elem.className = "w-full h-20 alert alert-detail shadow-lg animate__animated mt-10"
//     elem.innerHTML = `
//     <div>
//       <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" class="stroke-info flex-shrink-0 w-6 h-6"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"></path></svg>
//     <div>
//         <h3 class="font-bold">${title}</h3>
//         <div class="text-xs">${msg}</div>
//     </div>
//     </div>
//         <div class="flex-none">
//         <button class="btn btn-sm">查看</button>
//     </div>
//     `
//     return elem
// }

export default new message()
