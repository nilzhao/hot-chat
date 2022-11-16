import { STORAGE_KEYS } from '@/config';

export const updateFile = (
  filePath: string
): Promise<{
  data: {
    path: string;
    width: number;
    height: number;
  } | null;
  msg: string;
}> => {
  return new Promise((resolve) => {
    uni.uploadFile({
      filePath,
      url: `${__API_PREFIX__}/attach`,
      name: 'file',
      header: {
        'X-Token': uni.getStorageSync(STORAGE_KEYS.token),
        Accept: 'application/json',
      },
      success: ({ data }) => {
        resolve(JSON.parse(data));
      },
      fail: ({ errMsg }) => {
        resolve({
          data: null,
          msg: errMsg,
        });
      },
    });
  });
};
