import { isNil, isEmpty } from 'lodash';

export const isNilEmpty = (v: any) => {
  return isNil(v) || isEmpty(v);
};

export const getArrayValue = <T = any>(data: T | T[]) => {
  return Array.isArray(data) ? data[0] : data;
};
