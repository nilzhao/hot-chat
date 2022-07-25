<template>
  <view>
    <view v-if="goods" class="text-center mb-md">
      <view>{{ goods.username }}的红包</view>
      <view class="blessing">{{ goods.blessing || DEFAULT_BLESSING }}</view>
    </view>

    <uni-list>
      <!-- 显示圆形头像 -->
      <uni-list-chat
        v-for="item in list"
        :avatar-circle="true"
        :title="item.recvUser.name"
        :avatar="item.recvUser.avatar"
        :note="formatDateTime(item.createdAt)"
        :time="formatAmount(item.amount)"
      />
    </uni-list>
  </view>
</template>
<script lang="ts" setup async>
import { onLoad } from '@dcloudio/uni-app';
import { ref } from 'vue';
import { reqDetailEnvelope, reqEnvelopeItems } from './services';
import { EnvelopeGoodsItemWithUser, EnvelopeGoodsWithUser } from './types';
import { DEFAULT_BLESSING } from '@/config';
import { formatDateTime } from '@/utils/time';
import { formatAmount } from '@/utils/num';

const list = ref<EnvelopeGoodsItemWithUser[]>([]);
const goods = ref<EnvelopeGoodsWithUser | null>(null);

onLoad(async (option) => {
  if (!option.envelopeNo) return;
  const [goodsResult, listResult] = await Promise.all([
    reqDetailEnvelope(option.envelopeNo),
    reqEnvelopeItems(option.envelopeNo),
  ]);

  if (goodsResult.ok) {
    goods.value = goodsResult.data;
  }
  if (listResult.ok) {
    list.value = listResult.data;
  }
});
</script>
