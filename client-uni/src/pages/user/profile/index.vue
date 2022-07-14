<template>
  <view class="container">
    <view class="flex justify-center mb-md">
      <view class="round">
        <image
          :src="currentUser.avatar"
          style="width: 200px; height: 200px; background-color: #eeeeee"
        />
      </view>
    </view>
    <button v-if="currentUser.id" type="primary" @click="logout">退出</button>
    <navigator v-else url="/pages/user/login">
      <button type="primary">登录/注册</button>
    </navigator>
  </view>
</template>
<script lang="ts">
import { STORAGE_KEYS } from '@/config';
import { useAuthStore } from '@/stores/auth';
import { storeToRefs } from 'pinia';
import { defineComponent } from 'vue';

export default defineComponent({
  setup() {
    const authStore = useAuthStore();
    const { currentUser } = storeToRefs(authStore);

    // 退出登录
    const logout = () => {
      uni.removeStorageSync(STORAGE_KEYS.token);
      authStore.resetCurrentUser();
      // TODO: 请求退出的接口
    };

    return {
      currentUser,
      logout,
    };
  },
});
</script>
