import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
import handlebars from 'vite-plugin-handlebars'

export default defineConfig({
  root: 'src',
  build: {
    outDir: '../dist',
    emptyOutDir: true,
  },
  server: {
    port: 80,
    host: true,
  },
  plugins: [
    tailwindcss(),
    handlebars({
      partialDirectory: 'src/partials', 
      context: {
        siteTitle: 'My Static Site',
      },
    }),
  ],
})
