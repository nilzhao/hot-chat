<script lang="ts" setup>
import { reqSearchUser } from '@/services/user';
import { trim } from 'lodash';
import { ref } from 'vue';

const keyword = ref('');

const handleSearch = async () => {
  const v = trim(keyword.value);
  if (!v) return;
  const { data } = await reqSearchUser(keyword.value);
  if (Array.isArray(data) && data.length) {
    uni.navigateTo({
      url: `/pages/user/info/index?userInfo=${encodeURIComponent(
        JSON.stringify(data[0])
      )}`,
    });
  } else {
    uni.showToast({
      title: '没有结果~',
      icon: 'error',
    });
  }
};

const handleCancel = () => {
  keyword.value = '';
  uni.navigateBack();
};
</script>
<template>
  <view>
    <uni-search-bar
      @confirm="handleSearch"
      :focus="true"
      v-model="keyword"
      @cancel="handleCancel"
      @clear="keyword = ''"
      placeholder="输入账号/名称搜索"
    />
  </view>
</template>
