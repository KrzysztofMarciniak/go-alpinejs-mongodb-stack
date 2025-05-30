import { defineConfig } from 'vite';

export default defineConfig({
  root: 'src',
  build: {
    outDir: '../dist',
    emptyOutDir: true
  },
  server: {
    host: true,      // listen on 0.0.0.0
    port: 3000
  }
});

