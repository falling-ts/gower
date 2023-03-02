import "simplebar"
import $ from "jquery"
import api from "./api"
import { createApp } from "vue/dist/vue.esm-browser.prod"
import ResizeObserver from "resize-observer-polyfill"

window.ResizeObserver = ResizeObserver

export { createApp, $, api }

// Style
import "animate.css"
import "./styles/main.styl"
import "./styles/normalize.css"
import "tailwindcss/tailwind.css"
import "simplebar/dist/simplebar.css"

// Tailwindcss Dev Mode
import "./lib/tailwindcss.dev.js"
import "daisyui/dist/full.css"
