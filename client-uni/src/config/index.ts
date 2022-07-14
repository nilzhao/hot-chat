export const IS_DEV = import.meta.env.DEV;
export const IS_PROD = import.meta.env.PROD;
const uniPlatform = process.env.UNI_PLATFORM;
// 生产环境域名
const PROD_URL = 'http://www.grow.com';

export const API_BASE_URL = (() => {
  if (uniPlatform === 'h5') {
    return __API_PREFIX__;
  }
  if (IS_DEV) {
    return `http://${__PUBLIC_IP__}:${__API_PORT__}${__API_PREFIX__}`;
  }
  return `${PROD_URL}${__API_PREFIX__}`;
})();

export const STORAGE_KEYS = {
  token: 'token',
};
