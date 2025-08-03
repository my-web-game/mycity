import tailwindcss from '@tailwindcss/vite';
import { defineConfig } from 'vite';

export default defineConfig({
  base: "/mycity/",
  plugins: [tailwindcss()],
  server: {
    port: 3000,
  },
  optimizeDeps: {
    exclude: ["@babylonjs/havok"],
  },
  build: {
    target: 'esnext',
  },
}); 
