// vite.config.js
import { defineConfig } from 'vite'

import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
  root: 'src',              
  build: {
    outDir: '../dist',      
    emptyOutDir: true
  },
  server: {
    port: 80,
    host: true,
  },
  plugins: [tailwindcss()]
})
