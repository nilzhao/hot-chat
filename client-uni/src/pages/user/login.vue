<script lang="ts" setup>
import { computed, ref, watch } from 'vue';
import request from '@/utils/request';
import { STORAGE_KEYS } from '@/config';
import useAuthStore from '@/stores/auth';
import Input from '@/components/input/index.vue';
import GoBack from '@/components/go-back/index.vue';

const getInitialValues = () => {
  return {
    email: '',
    password: '',
    name: '',
  };
};

const { getCurrentUser } = useAuthStore();
const formData = ref(getInitialValues());
const failMessageRef = ref<any>(null);
const info = ref<{
  type: 'success' | 'warning' | 'info' | 'error';
  msg: string;
}>({
  type: 'error',
  msg: '',
});
const mode = ref<'login' | 'register'>('login');
const isLoginMode = computed(() => {
  return mode.value === 'login';
});
const isRegisterMode = computed(() => {
  return mode.value === 'register';
});

const reset = () => {
  formData.value = getInitialValues();
};

watch(mode, reset);

const login = async () => {
  const { ok, data, msg } = await request.post('/login', {
    data: formData.value,
  });
  if (ok) {
    uni.setStorageSync(STORAGE_KEYS.token, data.token);
    await getCurrentUser();
    uni.reLaunch({ url: '/pages/index/index' });
  } else {
    info.value = {
      msg,
      type: 'error',
    };
    failMessageRef.value!.open();
  }
};

const register = async () => {
  const { ok, data, msg } = await request.post('/register', {
    data: formData.value,
  });
  if (ok) {
    mode.value = 'login';
    info.value = {
      msg: '注册成功,请登录',
      type: 'success',
    };
    failMessageRef.value!.open();
  } else {
    info.value = {
      msg,
      type: 'error',
    };
    failMessageRef.value!.open();
  }
};
</script>

<template>
  <view class="wrapper text-white">
    <GoBack />
    <view class="inner w-85">
      <view class="f-34 mb-xxl">
        <text v-if="isLoginMode">登录</text>
        <text v-else>注册</text>
      </view>
      <Input
        v-if="isRegisterMode"
        label="用户名"
        v-model="formData.name"
        layout="internal"
        class="mb-lg"
      />
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
      <button
        @click="login"
        v-if="isLoginMode"
        class="border-none text-white bg-gradient"
      >
        登录
      </button>
      <button
        @click="register"
        v-else
        class="border-none text-white bg-gradient"
      >
        注册
      </button>
    </view>
    <view class="tip flex justify-between">
      <view v-if="isLoginMode" @click="mode = 'register'">还没有账号?注册</view>
      <view v-else @click="mode = 'login'">已有账号?登录</view>
    </view>
    <uni-popup ref="failMessageRef" :type="info.type" :mask-click="false">
      <uni-popup-message type="error" :message="info.msg" :duration="2000" />
    </uni-popup>
  </view>
</template>

<style lang="scss" scoped>
.wrapper {
  height: 100vh;
  background: url('/static/image/login-bg.jpg') center/cover;
  display: flex;
  justify-content: center;
  position: relative;
  align-items: center;
}
.inner {
  max-width: 400px;
}
.tip {
  position: absolute;
  bottom: 66px;
}
</style>
