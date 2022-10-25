import { isNil, isEmpty } from 'lodash';

export const isNilEmpty = (v: any) => {
  return isNil(v) || isEmpty(v);
};
