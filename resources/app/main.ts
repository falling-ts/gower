import "simplebar"
import $ from "jquery"
import api from "./api"
import msg from "./lib/msg"
import util from "./lib/util"
import cookie from "js-cookie"
import hash from "./lib/hash"
import { themeChange } from "theme-change"
// @ts-ignore
import { createApp, ref, reactive, computed, onBeforeMount, onMounted, onBeforeUpdate, onUpdated, onBeforeUnmount, onUnmounted, watch, watchEffect } from "vue/dist/vue.esm-browser.prod"
import ResizeObserver from "resize-observer-polyfill"

window.ResizeObserver = ResizeObserver

$.fn.extend({
    handle: function (callback: Function) {
        callback(this)
        return this
    }
})

export {
    createApp,
    ref,
    reactive,
    computed,
    onBeforeMount,
    onMounted,
    onBeforeUpdate,
    onUpdated,
    onBeforeUnmount,
    onUnmounted,
    watch,
    watchEffect,
    $,
    api,
    msg,
    cookie,
    hash,
    util,
    themeChange
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
