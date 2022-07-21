<template>
  <view class="">
    <uni-popup ref="goodsRef">
      <view class="goods">
        <image
          class="lantern"
          src="/static/image/lantern.png"
          mode="aspectFit"
        />
        <image
          class="lantern"
          src="/static/image/lantern.png"
          mode="aspectFit"
        />
        <view class="from">[{{ goods?.username }}]的红包</view>
        <view class="bless">{{ goods?.blessing || '恭喜发财 大吉大利' }}</view>
        <Avatar
          :src="goods!.user!.avatar"
          :text="goods!.user!.name"
          :width="80"
          class="mb-lg"
        />
        <view v-if="status === StatusEnum.FOUND" class="open-btn">
          <view class="open-text" @click="openEnvelope">开</view>
          <view class="open-border"> </view>
        </view>
        <view v-if="status > StatusEnum.FOUND">
          <view v-if="status === StatusEnum.OPEN_OK">
            <view class="amount">
              <text>88.88</text>
              <text class="unit">元</text>
            </view>
          </view>
          <view v-if="status === StatusEnum.OPEN_FAILED"
            >很遗憾,您没有抢到</view
          >
        </view>
        <navigator
          url="/pages/red-envelope/goods-item-list"
          open-type="navigate"
          hover-class="navigator-hover"
          class="detail-btn"
        >
          <text>查看领取详情</text>
          <uni-icons size="14" color="#fed983" type="forward" />
          <uni-icons size="14" color="#fed983" type="forward" />
        </navigator>
      </view>
    </uni-popup>
    <button type="primary" @click="findEnvelope">捞红包</button>
    <button type="primary" @click="sendEnvelope">发红包</button>
  </view>
</template>
<script lang="ts" setup>
import { ref } from 'vue';
import { reqFindEnvelope, reqReceiveEnvelopeItem } from './services';
import { EnvelopeGoodsWithUser, EnvelopeGoodsItem } from './types';
import Avatar from '@/components/avatar/index.vue';

enum StatusEnum {
  INIT = 0,
  FOUND = 1,
  OPEN_OK = 2,
  OPEN_FAILED = 3,
}

const goods = ref<EnvelopeGoodsWithUser | null>(null);
const goodsItem = ref<EnvelopeGoodsItem | null>(null);
const status = ref<number>(StatusEnum.INIT);
const goodsRef = ref<any>(null);
// 打捞红包
const findEnvelope = async () => {
  uni.showLoading({
    title: '打捞中...',
  });
  const { ok, data } = await reqFindEnvelope();
  if (ok) {
    goods.value = data;
    goodsRef.value!.open();
    status.value = StatusEnum.FOUND;
  }
  uni.hideLoading();
};
// 打开红包
const openEnvelope = async () => {
  if (!goods.value) return;
  uni.showLoading({
    title: '开启中...',
  });
  const { ok, data } = await reqReceiveEnvelopeItem(goods.value?.envelopeNo);
  if (ok) {
    goodsItem.value = data;
    status.value = StatusEnum.OPEN_OK;
  } else {
    status.value = StatusEnum.OPEN_FAILED;
  }
  uni.hideLoading();
};
// 发红包
const sendEnvelope = () => {};
</script>

<style lang="scss" scoped>
@use '~@/styles/_mixin.scss';

$red: #e13130;
$yellow: #fed983;

.goods {
  width: 320px;
  padding: 20px;
  border-radius: 10px;
  display: flex;
  flex-direction: column;
  align-items: center;
  text-align: center;
  background: $red;
  overflow: hidden;
  position: relative;
}

.lantern {
  position: absolute;
  width: 187px * 0.5;
  height: 357px * 0.5;
  left: 0;
  top: 0;
  z-index: 0;
  &:last-of-type {
    left: auto;
    right: 0;
  }
}

.from {
  font-size: 18px;
  color: $yellow;
  margin-bottom: 16px;
  text-align: center;
  position: relative;
  z-index: 1;
}

.bless {
  text-align: center;
  font-size: 14px;
  color: $yellow;
  margin-bottom: 24px;
}

.avatar {
  @include mixin.avatar(80);
}

.open-btn {
  padding: 10px;
  border-radius: 100%;
  background: $yellow;
  margin-bottom: 20px;
  display: inline-block;
}

.open-text {
  $width: 100px;
  font-size: 60px;
  width: $width;
  height: $width;
  line-height: $width;
  color: $red;
  font-weight: bold;
  border-radius: 100%;
  border: 3px solid $red;
  text-align: center;
  position: relative;
  z-index: 1;
}

.open-border {
  position: absolute;
  width: 1000px;
  height: 1000px;
  border-radius: 1000px;
  border: 2px solid $yellow;
  left: 50%;
  bottom: 120px;
  transform: translate(-50%, 0);
  z-index: 0;
}

.amount {
  color: $yellow;
  font-size: 40px;
  font-weight: bold;
  margin-right: -22px;
}

.unit {
  font-size: 16px;
  margin-left: 2px;
}

.detail-btn {
  color: $yellow;
}
</style>
