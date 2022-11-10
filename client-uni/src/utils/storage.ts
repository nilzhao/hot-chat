export const getLocalIItem = <T = string>(
  key: string,
  defaultValue?: T
): Promise<T | undefined> => {
  return new Promise((resolve) => {
    uni.getStorage({
      key,
      success(res) {
        try {
          resolve(JSON.parse(res.data));
        } catch (e) {
          resolve(res.data);
        }
      },
      fail(err) {
        console.log(err);
        resolve(defaultValue);
      },
    });
  });
};

export const setLocalItem = (key: string, data: any): Promise<void> => {
  return new Promise((resolve) => {
    if (data == null || data === '') {
      uni.removeStorage({
        key: key,
        complete() {
          resolve();
        },
        fail: console.log,
      });
    } else {
      try {
        const dataStr = JSON.stringify(data);
        uni.setStorage({
          key,
          data: dataStr,
          complete() {
            resolve();
          },
          fail: console.log,
        });
      } catch (e) {
        resolve();
      }
    }
  });
};
