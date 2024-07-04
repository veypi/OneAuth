// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  css: [
    '~/assets/css/app.scss',
  ],

  ssr: false,

  nitro: {
    devProxy: {
      "/api": {
        target: "http://127.0.0.1:4001/api",
        changeOrigin: true,
        ws: true,
      },
      '/fs': {
        target: 'http://127.0.0.1:4001/fs',
        changeOrigin: true,
        ws: true,
      },
      '/media': {
        target: 'http://127.0.0.1:4001/media',
        changeOrigin: true,
        ws: true,
      },
    },
  },

  app: {
    head: {
      meta: [
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'keywords', content: 'veypi, oneauth' }
      ],
      title: 'OneAuth',
      noscript: [
        { children: 'JavaScript is required' }
      ],
      link: [
        { rel: 'icon', type: 'image/ico', href: 'favicon.ico' }
      ],
      script: [
        { src: '/icon.js' },
      ]
    },
  },
  modules: ["@pinia/nuxt", "@nuxt/ui"]
})