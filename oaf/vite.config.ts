import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import path from "path"

// https://vitejs.dev/config/
export default defineConfig({
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "src"),
        },
    },
    plugins: [vue()],
    server: {
        host: '0.0.0.0',
        // host: '127.0.0.1',
        port: 8080,
        proxy: {
            '/api': {
                target: 'http://127.0.0.1:4001/',
                changeOrigin: true,
                ws: true
            },
            '/media': {
                target: 'http://127.0.0.1:4001/',
                changeOrigin: true,
                ws: true
            }
        }
    },
    build: {
        outDir: '../sub/static/',
        assetsDir: './',
        rollupOptions: {
            output: {
                // 重点在这里哦
                entryFileNames: `static/[name].[hash].js`,
                chunkFileNames: `static/[name].[hash].js`,
                assetFileNames: `static/[name].[hash].[ext]`
            }
        }
    }
})
