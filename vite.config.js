import vue from '@vitejs/plugin-vue';
import replace from '@rollup/plugin-replace';

export default {
    plugins: [
        vue(),
        replace({
            'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV || 'development'),
            'process.env': JSON.stringify({}),
            preventAssignment: true
        })
    ],
    mode: process.env.NODE_ENV,
    publicDir: 'storage/app/public',
    build: {
        lib: {
            entry: 'resources/app/main.ts',
            name: 'main',
            fileName: 'main',
        },
        outDir: 'public/static',
        rollupOptions: {
            input: 'resources/app/main.ts',
            output: {
                assetFileNames: '[name][extname]'
            }
        }
    }
}
