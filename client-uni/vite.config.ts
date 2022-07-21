import { defineConfig } from 'vite';
import uni from '@dcloudio/vite-plugin-uni';
import { join } from 'path';
import { networkInterfaces } from 'os';
import vueTypeImports from 'vite-plugin-vue-type-imports';

const API_PORT = 9000;
const API_PREFIX = '/api/v1';

const getPublicIp = () => {
  const network = Object.values(networkInterfaces())
    .flat()
    .find((item) => item && item.family === 'IPv4' && !item.internal);

  if (!network) {
    throw Error('获取公网 IP 失败');
  }
  return network.address;
};

const publicIp = getPublicIp();

const pathJoin = (dir: string) => {
  return join(__dirname, dir);
};

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [uni(), vueTypeImports()],
  envPrefix: 'GROW_',
  define: {
    __API_PREFIX__: JSON.stringify(API_PREFIX),
    __API_PORT__: JSON.stringify(API_PORT),
    __PUBLIC_IP__: JSON.stringify(publicIp),
  },
  resolve: {
    alias: [
      {
        find: /@\//,
        replacement: pathJoin('src/'),
      },
    ],
  },
  server: {
    port: 9001,
    proxy: {
      [API_PREFIX]: `http://${publicIp}:${API_PORT}`,
    },
  },
});
