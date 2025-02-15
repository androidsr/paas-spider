import vue from '@vitejs/plugin-vue';
import { resolve } from 'path';
import Components from 'unplugin-vue-components/vite';
import { defineConfig } from 'vite';
import Pages from "vite-plugin-pages";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    Components({}),
    Pages({
      dirs: [{ dir: "src/pages", baseRoute: "" }],
      importMode: "async",
      extendRoute(route) {
        let meta = {};
        return {
          ...route,
          meta
        }
      }
    })
  ],
  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'), // 设置 @ 别名为 src 目录
      '*': resolve('')
    },
    transpileDependencies: true,
    lintOnSave: false,
  },
  css: {
    preprocessorOptions: {
      less: {
        math: "always",
        additionalData: '@import "@/assets/style/global.less";',
        javascriptEnabled: true,
      }
    }
  },
})
