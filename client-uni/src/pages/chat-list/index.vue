<script lang="ts" setup>
import useAuthStore from '@/stores/auth';
import useContactStore from '@/stores/contact';
import { Chat } from '@/types/chat';
import CacheChat from '@/utils/cache-chat';
import { onMounted, ref, watch } from 'vue';
import { formatDateTime } from '@/utils/time';
import useWsStore from '@/stores/ws';

const chatList = ref<Chat[]>([]);
const wsStore = useWsStore();

const authStore = useAuthStore();
const contactStore = useContactStore();

const cacheChat = new CacheChat(
  authStore.currentUser,
  contactStore.contactsMap
);

onMounted(async () => {
  chatList.value = await cacheChat.getChatList();
});

watch(
  () => wsStore.newMsg,
  async (v) => {
    if (!v) return;
    chatList.value = await cacheChat.refreshChatList(v);
  }
);
</script>
<template>
  <uni-list>
    <template v-for="chat in chatList" :key="chat.targetId">
      <uni-list-chat
        link
        clickable
        :to="`/pages/chat/index?targetInfo=${JSON.stringify({
          name: chat.name,
          avatar: chat.avatar,
          id: chat.targetId,
        })}`"
        :avatar-circle="true"
        :title="chat.name"
        :avatar="chat.avatar"
        :note="chat.note"
        :time="formatDateTime(chat.updatedAt)"
      />
    </template>
  </uni-list>
</template>
