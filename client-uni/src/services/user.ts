import request from '@/utils/request';

export const reqGetUserInfo = (id: number | string) =>
  request.get(`/users/${id}`);

export const reqSearchUser = (keyword: string) =>
  request.get('/users/search', {
    data: {
      keyword,
    },
  });

export const reqAddFriend = (targetUserId: number) =>
  request.post('/contact/friends/:id', {
    routeParams: {
      id: targetUserId,
    },
  });
