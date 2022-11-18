<template>
  <view class="">
    <uni-popup ref="goodsPopupRef">
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
        <Avatar
          round
          :src="goods!.user!.avatar"
          :text="goods!.user!.name"
          :width="80"
          class="mb-sm"
        />
        <view class="from">[{{ goods?.username }}]的红包</view>
        <view class="bless">{{ goods?.blessing || DEFAULT_BLESSING }}</view>
        <view class="info">
          <template v-if="status === StatusEnum.FOUND">
            <view @click="openEnvelope" class="open-btn">
              <view class="open-text">開</view>
            </view>
            <view class="open-border"> </view>
          </template>

          <view v-if="status > StatusEnum.FOUND">
            <view v-if="status === StatusEnum.OPEN_OK">
              <view class="amount">
                <text>{{ formatAmount(goodsItem?.amount) }}</text>
                <text class="unit">元</text>
              </view>
            </view>
            <view v-if="status === StatusEnum.OPEN_FAILED">
              <text>很遗憾,您没有抢到</text>
            </view>
          </view>
        </view>
        <navigator
          :url="`/pages/red-envelope/goods-item-list?envelopeNo=${goods?.envelopeNo}`"
          open-type="navigate"
          hover-class="navigator-hover"
          class="detail-btn"
        >
          <text>查看领取详情</text>
          <text class="iconfont icon-doubleright" />
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
import { DEFAULT_BLESSING } from '@/config';
import { formatAmount } from '@/utils/num';

enum StatusEnum {
  INIT = 0,
  FOUND = 1,
  OPEN_OK = 2,
  OPEN_FAILED = 3,
}

const goods = ref<EnvelopeGoodsWithUser | null>(null);
const goodsItem = ref<EnvelopeGoodsItem | null>(null);
const status = ref<number>(StatusEnum.INIT);
const goodsPopupRef = ref<any>(null);
const sendPopupRef = ref<any>(null);
// 打捞红包
const findEnvelope = async () => {
  uni.showLoading({
    title: '打捞中...',
  });
  const { ok, data } = await reqFindEnvelope();
  uni.hideLoading();
  if (ok) {
    goods.value = data;
    goodsPopupRef.value!.open();
    status.value = StatusEnum.FOUND;
  } else {
    uni.showModal({
      title: '空空如也~',
    });
  }
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
const sendEnvelope = async () => {
  uni.navigateTo({ url: '/pages/red-envelope/send-out' });
};
</script>

<style lang="scss" scoped>
@use '~@/styles/_mixin.scss';

$red: #e13130;
$yellow: #fed983;

$info-height: 120px;
$info-padding: 10px;

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
  color: $yellow;
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
  margin-bottom: 16px;
  text-align: center;
  position: relative;
  z-index: 1;
}

.bless {
  text-align: center;
  font-size: 16px;
  margin-bottom: 100px;
}

.avatar {
  @include mixin.avatar(80);
}

.info {
  height: $info-height;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
}

.open-btn {
  padding: $info-padding;
  border-radius: 100%;
  background: $yellow;
  margin-bottom: 20px;
}

.open-text {
  $width: $info-height - $info-padding * 2;
  font-size: 60px;
  color: $red;
  width: $width;
  height: $width;
  line-height: $width;
  font-weight: bold;
  border-radius: 100%;
  border: 3px solid $red;
  text-align: center;
  position: relative;
  z-index: 1;
}

.open-border {
  $width: 1000px;
  position: absolute;
  width: $width;
  height: $width;
  border-radius: $width;
  border: 2px solid $yellow;
  left: 50%;
  bottom: $info-height * 0.5;
  transform: translate(-50%, 0);
  z-index: 0;
}

.amount {
  font-size: 40px;
  font-weight: bold;
  margin-right: -20px;
}

.unit {
  font-size: 16px;
  margin-left: 2px;
}

.detail-btn {
  display: flex;
  align-items: center;
}
</style>
