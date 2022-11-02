import request from '@/utils/request';

export const reqGetUserInfo = (id: number | string) =>
  request.get(`/users/${id}`);

export const reqSearchUser = (keyword: string) =>
  request.get('/users', {
    data: {
      keyword,
    },
  });
