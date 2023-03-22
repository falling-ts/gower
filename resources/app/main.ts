import "simplebar"
import $ from "jquery"
import api from "./api"
import msg from "./lib/msg"
import cookie from "js-cookie"
import * as store from "localforage"
import { createApp } from "vue/dist/vue.esm-browser.prod"
import ResizeObserver from "resize-observer-polyfill"

window.ResizeObserver = ResizeObserver

export {
    createApp,
    $,
    api,
    msg,
    cookie,
    store
}

// 样式
import "animate.css"
import "./styles/main.styl"
import "./styles/normalize.css"
import "tailwindcss/tailwind.css"
import "simplebar/dist/simplebar.css"

// Tailwindcss 开发模式
import "./lib/tailwindcss.dev.js"
import "daisyui/dist/full.css"
