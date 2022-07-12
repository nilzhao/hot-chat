<template>
  <view>
    <uni-card title="登录">
      <uni-forms :modelValue="formData">
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
        <button @click="submit" type="primary" class="mb-md">登录</button>
        <view class="flex justify-between">
          <navigator
            url="/pages/user/register"
            open-type="navigate"
            hover-class="navigator-hover"
          >
            还没有账号?注册
          </navigator>
        </view>
      </uni-forms>
    </uni-card>
    <uni-popup ref="failMessageRef" type="message" :mask-click="false">
      <uni-popup-message
        type="error"
        message="登录失败,请重试"
        :duration="2000"
      />
    </uni-popup>
  </view>
</template>

<script lang="ts">
import { defineComponent, ref, getCurrentInstance } from "vue";
import request from "@/utils/request";
import { STORAGE_KEYS } from "@/config";

const Login = defineComponent({
  setup() {
    const formData = ref({
      email: "",
      password: "",
    });
    const failMessageRef = ref(null);
    const submit = async () => {
      const { ok, data } = await request.post("/login", {
        data: formData.value,
      });
      if (ok) {
        uni.setStorageSync(STORAGE_KEYS.token, data.token);
        uni.reLaunch({ url: "/pages/index/index" });
      } else {
        failMessageRef.value?.open();
      }
    };
    return {
      formData,
      submit,
      failMessageRef,
    };
  },
});

export default Login;
</script>
