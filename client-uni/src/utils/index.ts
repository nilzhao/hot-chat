import { isNil, isEmpty } from 'lodash';
import { compile } from 'path-to-regexp';

export const isNilEmpty = (v: any) => {
  return isNil(v) || isEmpty(v);
};

export const getArrayValue = <T = any>(data: T | T[]) => {
  return Array.isArray(data) ? data[0] : data;
};

export const compileUrl = (
  urlPath: string,
  params: Record<string, string | number>
) => {
  const toPath = compile(urlPath, { encode: encodeURIComponent });

  return toPath(params);
};
