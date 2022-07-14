import { API_BASE_URL, STORAGE_KEYS } from '@/config';

type BaseRequestOptions = Omit<
  UniApp.RequestOptions,
  'method' | 'url' | 'success' | 'complete' | 'fail'
>;

const SUCCESS_CODE = 0;

interface RequestOptions extends BaseRequestOptions {
  routeParams?: Record<string, string | number>;
}

interface RequestResult<T = any> {
  ok: boolean;
  data: T;
  msg: string;
  code: number;
}

interface Request {
  <T = any>(options: RequestOptions): Promise<RequestResult<T>>;
  get<T = any>(
    url: string,
    options?: BaseRequestOptions
  ): Promise<RequestResult<T>>;
  post<T = any>(
    url: string,
    options?: BaseRequestOptions
  ): Promise<RequestResult<T>>;
  put<T = any>(
    url: string,
    options?: BaseRequestOptions
  ): Promise<RequestResult<T>>;
  delete<T = any>(
    url: string,
    options?: BaseRequestOptions
  ): Promise<RequestResult<T>>;
}

async function request(options: UniApp.RequestOptions) {
  const method =
    options.method?.toLocaleUpperCase() as UniApp.RequestOptions['method'];
  const result: RequestResult<any> = {
    ok: true,
    code: 0,
    msg: '',
    data: null,
  };
  const url = `${API_BASE_URL}${options.url}`;

  const header = {
    ...options.header,
    'X-Token': uni.getStorageSync(STORAGE_KEYS.token),
  };

  try {
    // @ts-ignore
    const { data } = await uni.request({
      ...options,
      method,
      url,
      header,
    });
    if (data.code !== SUCCESS_CODE) {
      throw new Error(data.msg);
    }
    // @ts-ignore
    Object.assign(result, data);
  } catch (err) {
    const error = err as any;
    Object.assign(result, {
      ok: false,
      code: 400,
      msg: error.errMsg || error.message,
    });
  }
  return result;
}

type Method = 'get' | 'post' | 'put' | 'delete';
const methods: Method[] = ['get', 'post', 'put', 'delete'];

methods.forEach((method) => {
  // @ts-ignore
  request[method] = (url: string, options?: BaseRequestOptions) => {
    return request({
      ...options,
      method: method as UniApp.RequestOptions['method'],
      url,
    });
  };
});

export default request as Request;
