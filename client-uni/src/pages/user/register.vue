<template>
  <uni-card title="注册">
    <uni-forms :modelValue="formData">
      <uni-forms-item>
        <uni-easyinput
          type="text"
          v-model="formData.name"
          placeholder="用户名"
        />
      </uni-forms-item>
      <uni-forms-item>
        <uni-easyinput
          type="text"
          v-model="formData.email"
          placeholder="邮箱"
        />
      </uni-forms-item>
      <uni-forms-item>
        <uni-easyinput
          type="password"
          v-model="formData.password"
          placeholder="密码"
        />
      </uni-forms-item>
      <button @click="submit" type="primary" class="mb-md">注册</button>
      <view class="flex justify-between">
        <navigator
          url="/pages/user/login"
          open-type="navigate"
          hover-class="navigator-hover"
        >
          已有账号?登录
        </navigator>
      </view>
    </uni-forms>
  </uni-card>
</template>

<script lang="ts" setup>
import { ref } from 'vue';

const formData = ref({
  email: '',
  password: '',
  name: '',
});
const submit = () => {
  console.log(formData.value);
};
</script>
