import { INIT_USER } from '@/types/user';
import request from '@/utils/request';
import { defineStore } from 'pinia';

const useAuthStore = defineStore('auth', {
  state: () => {
    return {
      currentUser: INIT_USER,
      network: {
        type: '',
        isConnected: true,
      },
    };
  },
  getters: {
    isLogin: (state) => !!(state.currentUser && state.currentUser.id),
  },
  actions: {
    async getCurrentUser() {
      const { ok, data } = await request.get('/user/profile');
      if (ok) {
        this.currentUser = data;
      } else {
        uni.reLaunch({
          url: '/pages/user/login',
        });
      }
    },
    resetCurrentUser() {
      this.currentUser = INIT_USER;
    },
  },
});

export default useAuthStore;
