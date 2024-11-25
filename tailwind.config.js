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
    ]
};
