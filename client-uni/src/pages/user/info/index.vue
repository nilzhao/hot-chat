<script lang="ts" setup>
import { INIT_USER, User } from '@/types/user';
import { onLoad } from '@dcloudio/uni-app';
import { isObject } from 'lodash';
import { ref } from 'vue';
import Avatar from '@/components/avatar/index.vue';
import { UserGenderEnum } from '@/types/user';
import { reqAddFriend } from '@/services/user';
import Navbar from '@/components/navbar/index.vue';
import useContactStore from '@/stores/contact';

const userInfo = ref<User>({
  ...INIT_USER,
});

const contactStore = useContactStore();

onLoad((query) => {
  try {
    const info: User = JSON.parse(decodeURIComponent(query.userInfo || ''));
    if (isObject(info)) {
      userInfo.value = info;
    }
  } catch (err) {
    console.log(err);
  }
});

const addFriend = async () => {
  const { ok, msg } = await reqAddFriend(userInfo.value.id);
  if (ok) {
    contactStore.getContacts();
    uni.showToast({ title: '添加成功' });
    uni.switchTab({
      url: '/pages/chat-list/index',
    });
  } else if (msg) {
    uni.showToast({ title: msg, icon: 'error' });
  }
};
</script>
<template>
  <view>
    <Navbar />
    <view class="px-md py-md">
      <view class="flex mb-lg">
        <Avatar :src="userInfo?.avatar" :text="userInfo.name" class="mr-md" />
        <view class="flex flex-column justify-between py-sm">
          <view>
            <text class="f-bold f-18">{{ userInfo.name }}</text>
            <text
              v-if="userInfo.gender === UserGenderEnum.MALE"
              class="iconfont icon-male text-primary"
            />
            <text
              v-else-if="userInfo.gender === UserGenderEnum.FEMALE"
              class="iconfont icon-female text-error"
            />
          </view>
          <view class="text-grey"> 地区：安道尔 </view>
        </view>
      </view>
      <button type="primary" @click="addFriend" size="small">
        添加到通讯录
      </button>
    </view>
  </view>
</template>

<style lang="scss" scoped>
.avatar {
}
</style>
