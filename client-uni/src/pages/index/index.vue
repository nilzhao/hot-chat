<template>
  <view class="container">
    <view>欢迎来到 GROW !</view>
    <text class="iconfont icon-user f-12" />
    <view>网络状态: {{ authStore.network.isConnected ? '连接' : '断开' }}</view>
    <view>网络类型: {{ authStore.network.type }}</view>
    <button @click="chooseImg">选择图片</button>
  </view>
</template>

<script lang="ts" setup>
import { useAuthStore } from '@/stores/auth';
const authStore = useAuthStore();

const chooseImg = () => {
  uni.chooseImage({
    count: 6, //默认9
    sizeType: ['original', 'compressed'], //可以指定是原图还是压缩图，默认二者都有
    sourceType: ['album'], //从相册选择
    success: function (res) {
      console.log(JSON.stringify(res.tempFilePaths));
    },
  });
};
</script>
