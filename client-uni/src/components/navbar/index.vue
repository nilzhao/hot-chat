<script lang="ts" setup>
import GoBack from '@/components/go-back/index.vue';
import { navbarProps } from './types';

const props = defineProps(navbarProps);
</script>

<template>
  <view class="navbar">
    <view class="content">
      <view class="side">
        <slot name="left" v-if="props.hideLeft !== true">
          <GoBack class="text-grey" />
        </slot>
        <view></view>
        <slot name="right"> </slot>
      </view>
      <view class="center">
        <slot>
          {{ props.title }}
        </slot>
      </view>
    </view>
  </view>
</template>

<style lang="scss" scoped>
@use '~@/styles/_var.scss' as *;

$navbar-height: 44px;

.navbar {
  height: $navbar-height;
  position: relative;
}
.content {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: $navbar-height;
  z-index: 999;
  padding: 0 $space-sm;
  box-shadow: 0 0 2px rgba(0, 0, 0, 0.5);
}
.side {
  position: relative;
  display: flex;
  justify-content: space-between;
  align-items: center;
  z-index: 2;
  height: $navbar-height;
}
.center {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1;
}
</style>
