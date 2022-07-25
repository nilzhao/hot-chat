import { INIT_USER } from '@/types/user';
import request from '@/utils/request';
import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
  state: () => {
    return {
      currentUser: INIT_USER,
      network: {
        type: '',
        isConnected: true,
      },
    };
  },
  actions: {
    async getCurrentUser() {
      const { ok, data } = await request.get('/user/profile');
      if (ok) {
        this.currentUser = data;
      }
    },
    resetCurrentUser() {
      this.currentUser = INIT_USER;
    },
  },
});
