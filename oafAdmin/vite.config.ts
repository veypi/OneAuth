import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            '@': resolve(__dirname, 'src'),
        },
    },
    server: {
        host: '127.0.0.1',
        port: 3000,
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:4001/',
                changeOrigin: true,
                ws: true,
            },
            '/media': {
                target: 'http://127.0.0.1:4001/',
                changeOrigin: true,
                ws: true,
            },
        },
    },
    plugins: [vue()],
    build: {
        outDir: './dist/',
        assetsDir: './',
        rollupOptions: {
            output: {
                entryFileNames: `static/[name].[hash].js`,
                chunkFileNames: `static/[name].[hash].js`,
                assetFileNames: `static/[name].[hash].[ext]`,
            },
        },
    },
})
