<template>
  <view>
    <view>
      <view>的红包</view>
      <view class="blessing">恭喜发财,大吉大利</view>
    </view>

    <uni-list>
      <uni-list-item
        title="10个红包,1 分钟被抢光"
        note="列表描述信息"
      ></uni-list-item>
      <uni-list-item
        :disabled="true"
        title="列表文字"
        note="列表禁用状态"
      ></uni-list-item>
    </uni-list>
  </view>
</template>
<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  setup() {},
});
</script>
