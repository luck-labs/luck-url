import { defineConfig } from 'umi';

export default defineConfig({
  nodeModulesTransform: {
    type: 'none',
  },
  routes: [
    { path: '/', component: '@/pages/index' },
  ],
  fastRefresh: {},
  hash: true,
  history: {type: 'hash'},
  theme: {
    'primary-color': "#6CC449"
  }
});
