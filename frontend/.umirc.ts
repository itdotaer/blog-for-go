import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  locale: { antd: true },
  routes: [
    { path: '/', component: '@/pages/index' },
    { path: '/posts', component: '@/pages/posts' },
    { path: '/posts/:id', component: '@/pages/post' },
  ],
});
