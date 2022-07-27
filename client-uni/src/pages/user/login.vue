<template>
  <view class="wrapper text-white">
    <view class="w-85">
      <view class="f-34 mb-xxl">登录</view>
      <Input
        label="邮箱"
        v-model="formData.email"
        layout="internal"
        class="mb-lg"
      />
      <Input
        label="密码"
        v-model="formData.password"
        layout="internal"
        class="mb-xxl"
        type="password"
      />
      <button @click="submit" class="border-none text-white bg-gradient">
        登录
      </button>
    </view>
    <view class="bottom flex justify-between">
      <navigator
        url="/pages/user/register"
        open-type="navigate"
        hover-class="navigator-hover"
      >
        还没有账号?注册
      </navigator>
    </view>

    <uni-popup ref="failMessageRef" type="message" :mask-click="false">
      <uni-popup-message type="error" :message="errMsg" :duration="2000" />
    </uni-popup>
  </view>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import request from '@/utils/request';
import { STORAGE_KEYS } from '@/config';
import { useAuthStore } from '@/stores/auth';
import Input from '@/components/input/index.vue';

const { getCurrentUser } = useAuthStore();
const formData = ref({
  email: '',
  password: '',
});
const errMsg = ref<string>('');
const failMessageRef = ref<any>(null);
const submit = async () => {
  const { ok, data, msg } = await request.post('/login', {
    data: formData.value,
  });
  if (ok) {
    uni.setStorageSync(STORAGE_KEYS.token, data.token);
    await getCurrentUser();
    uni.reLaunch({ url: '/pages/index/index' });
  } else {
    errMsg.value = msg;
    failMessageRef.value!.open();
  }
};
</script>

<style lang="scss" scoped>
.wrapper {
  height: calc(100vh - 44px);
  background: url('/static/image/login-bg.jpg') center/cover;
  display: flex;
  justify-content: center;
  position: relative;
  align-items: center;
}

.bottom {
  position: absolute;
  bottom: 66px;
}
</style>
