import daisyui from 'daisyui';
import { addDynamicIconSelectors } from "@iconify/tailwind"

export default {
    content: [
        './resources/app/main.ts',
        './resources/**/*.{js,ts,tmpl}'
    ],
    theme: {
        extend: {},
    },
    plugins: [
        daisyui,
        addDynamicIconSelectors()
    ],
    daisyui: {
        themes: [
            "light",
            "dark",
            "cupcake",
            "bumblebee",
            "emerald",
            "corporate",
            "synthwave",
            "retro",
            "cyberpunk",
            "valentine",
            "halloween",
            "garden",
            "forest",
            "aqua",
            "lofi",
            "pastel",
            "fantasy",
            "wireframe",
            "black",
            "luxury",
            "dracula",
            "cmyk",
            "autumn",
            "business",
            "acid",
            "lemonade",
            "night",
            "coffee",
            "winter",
            "dim",
            "nord",
            "sunset",
        ],
    },
};
