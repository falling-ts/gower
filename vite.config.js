export default {
    build: {
        rollupOptions: {
            input: 'resources/js/main.js',
            output: {
                dir: 'public/static',
                entryFileNames: "[name].js",
                assetFileNames: "[name][extname]"
            }
        }
    }
}