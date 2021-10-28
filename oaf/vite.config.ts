import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    // host: '0.0.0.0',
    host: '127.0.0.1',
    port: 8080,
    proxy: {
      '/api': 'http://127.0.0.1:4001/'
    }
  },
  build: {
    outDir: '../build/static/',
    assetsDir: 'assets'
  }
})
