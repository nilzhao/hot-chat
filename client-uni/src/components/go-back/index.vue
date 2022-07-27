<template>
  <view
    :class="classNames('iconfont icon-left', styles.back, props.class)"
    @click="goBack"
  />
</template>
<script lang="ts" setup>
import classNames from 'classnames';
import styles from './index.module.scss';

const props = defineProps<{
  class?: string;
}>();

const goBack = () => {
  uni.navigateBack({
    fail() {
      uni.switchTab({ url: '/pages/index/index' });
    },
  });
};
</script>
