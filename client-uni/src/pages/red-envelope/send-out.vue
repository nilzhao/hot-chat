<template>
  <view>
    <uni-card>
      <view class="type" @click="showEnvelopeTypePopup">
        <template v-for="option in envelopeTypeOptions">
          <view :key="option.label" v-if="option.value === envelopeType">
            {{ option.label }}
          </view>
        </template>
        <text class="iconfont icon-down" />
      </view>
      <Input
        type="number"
        label="红包个数"
        v-model="envelope.quantity"
        placeholder="0"
      />
      <Input
        type="number"
        label="总金额"
        v-model="envelope.amount"
        placeholder="¥0.00"
        v-if="envelopeType === EnvelopeGoodsTypeEnum.LUCKY"
      />
      <Input
        type="number"
        label="单个金额"
        v-model="envelope.amountOne"
        placeholder="¥0.00"
        v-if="envelopeType === EnvelopeGoodsTypeEnum.GENERAL"
      />
      <Input
        label="祝福语"
        v-model="envelope.blessing"
        :placeholder="DEFAULT_BLESSING"
      />
      <button class="bg-gradient border-none text-white" @click="send">
        塞进红包
      </button>
    </uni-card>
    <PopupSelect
      :options="envelopeTypeOptions"
      v-model="envelopeType"
      v-model:visible="envelopeTypeSelectVisible"
    />
  </view>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { EnvelopeGoodsTypeEnum, EnvelopeToSendOut } from './types';
import { DEFAULT_BLESSING } from '@/config';
import Input from '@/components/input/index.vue';
import PopupSelect from '@/components/popup-select/index.vue';
import { reqSendOutEnvelope } from './services';

const envelopeTypeOptions = [
  {
    value: EnvelopeGoodsTypeEnum.GENERAL,
    label: '普通红包',
  },
  {
    value: EnvelopeGoodsTypeEnum.LUCKY,
    label: '拼手气红包',
  },
];

const envelopeType = ref<EnvelopeGoodsTypeEnum>(EnvelopeGoodsTypeEnum.LUCKY);
const envelopeTypeSelectVisible = ref(false);
const envelope = ref<Partial<EnvelopeToSendOut>>({
  blessing: '',
});

const showEnvelopeTypePopup = () => {
  envelopeTypeSelectVisible.value = true;
};

const send = async () => {
  uni.showLoading({
    title: '正在发送...',
  });
  const { ok, msg } = await reqSendOutEnvelope({
    ...envelope.value,
    type: envelopeType.value,
  });
  uni.hideLoading();
  if (ok) {
    uni.switchTab({
      url: '/pages/red-envelope/index',
    });
  } else {
    uni.showToast({ title: msg });
  }
};
</script>

<style lang="scss" scoped>
.type {
  color: rgb(243, 195, 132);
  overflow: hidden;
  margin-bottom: 20px;
  display: flex;
  align-items: center;
}
.type-selector {
  border-radius: 10px 10px 0 0;
  overflow: hidden;
}
</style>
