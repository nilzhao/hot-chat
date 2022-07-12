import { defineConfig } from 'vite'
import uni from '@dcloudio/vite-plugin-uni'
import { join } from 'path'

const pathJoin = (dir: string) => {
  return join(__dirname, dir)
}

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [uni()],
  resolve: {
    alias: [
      {
        find: /@\//,
        replacement: pathJoin('src/'),
      },
    ],
  },
})
