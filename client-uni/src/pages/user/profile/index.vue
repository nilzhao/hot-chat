<template>
  <view class="container">
    <view class="flex justify-center mb-md">
      <view class="avatar">
        <image
          v-if="currentUser.id"
          :src="currentUser.avatar"
          class="w-full h-full"
        />
        <navigator v-else url="/pages/user/login">登录 / 注册</navigator>
      </view>
    </view>
    <button v-if="currentUser.id" type="primary" @click="logout">退出</button>
  </view>
</template>
<script lang="ts" setup>
import { STORAGE_KEYS } from '@/config';
import { useAuthStore } from '@/stores/auth';
import { storeToRefs } from 'pinia';

const authStore = useAuthStore();
const { currentUser } = storeToRefs(authStore);

// 退出登录
const logout = () => {
  uni.removeStorageSync(STORAGE_KEYS.token);
  authStore.resetCurrentUser();
  // TODO: 请求退出的接口
};
</script>

<style lang="scss" scoped>
@use '~@/styles/_mixin.scss';

.avatar {
  @include mixin.avatar(100);
  background: $uni-color-subtitle;
  color: white;
}
</style>
