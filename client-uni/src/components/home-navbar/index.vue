<script lang="ts" setup>
import { ref } from 'vue';
import Navbar from '@/components/navbar/index.vue';
import classNames from 'classnames';

const moreVisible = ref(false);
const searchFriendPopupRef = ref();
</script>
<template>
  <view>
    <Navbar fixed>
      <view class="input-view">你好</view>
      <block v-slot:right>
        <view class="more" @mouseleave="moreVisible = false">
          <view class="iconfont icon-plus" @mouseenter="moreVisible = true" />
          <view :class="classNames('dropdown', { visible: moreVisible })">
            <navigator
              url="/pages/search/index"
              class="item"
              @click="moreVisible = false"
            >
              <view class="icon iconfont icon-add-friend" />
              <view class="text">添加好友</view>
            </navigator>
            <view class="item">
              <view class="icon iconfont icon-scan" />
              <view class="text">扫一扫</view>
            </view>
          </view>
        </view>
      </block>
    </Navbar>
    <uni-popup ref="searchFriendPopupRef" type="bottom"> </uni-popup>
  </view>
</template>

<style lang="scss" scoped>
@use '~@/styles/_var.scss' as *;

.more {
  position: relative;
  height: 100%;
  display: flex;
  align-items: center;
}
.dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  box-shadow: 0 0 1px rgba(0, 0, 0, 0.5);
  width: 200px;
  background-color: #fff;
  display: none;
  border-radius: $uni-border-radius-sm;
  &.visible {
    display: block;
  }
  .item {
    display: flex;
    align-items: center;
    .icon {
      padding: 0 $space-xs;
    }
    .text {
      flex: 1;
      border-bottom: 1px solid #ddd;
      padding: $space-xs 0;
    }
    &:last-child {
      .text {
        border-bottom: none;
      }
    }
  }
}
</style>
