import { STORAGE_KEYS } from '@/config';
import { Message, MessageMediaEnum } from '@/types/chat';
import { defineStore } from 'pinia';
import { computed, ref } from 'vue';
import useAuthStore from './auth';
import CacheChat from '@/utils/cache-chat';
import { User } from '@/types/user';
import { ContactDetail } from '@/types/contact';

export enum WsStatusEnum {
  WAITING = 'waiting',
  CONNECTED = 'connected',
  ERROR = 'error',
  CLOSED = 'closed',
}

const useWsStore = defineStore('ws', () => {
  const status = ref<WsStatusEnum>(WsStatusEnum.WAITING);
  const newMsg = ref<Message | null>(null);
  let ws: UniApp.SocketTask;

  const init = ({
    url,
    currentUser,
    contactsMap,
  }: {
    url?: string;
    currentUser: User;
    contactsMap: Map<number, ContactDetail>;
  }) => {
    if (ws) return;
    const cacheChat = new CacheChat(currentUser, contactsMap);
    ws = uni.connectSocket({
      url:
        url ||
        `ws://127.0.0.1:9000${__API_PREFIX__}/ws?X-Token=${uni.getStorageSync(
          STORAGE_KEYS.token
        )}`,
      complete: () => {},
    });

    ws.onOpen(() => {
      status.value = WsStatusEnum.CONNECTED;
    });

    ws.onError((e: UniApp.GeneralCallbackResult) => {
      status.value = WsStatusEnum.ERROR;
      console.dir(e);
    });

    ws.onMessage(async (e: UniApp.OnSocketMessageCallbackResult) => {
      if (typeof e.data !== 'string') return;
      try {
        const data: Message = JSON.parse(e.data);
        console.log('receive message', data);

        newMsg.value = data;
        cacheChat.cacheOnReceivedMessage(data);
      } catch (e) {
        console.log('消息解析失败:', e);
      }
    });

    ws.onClose(() => {
      status.value = WsStatusEnum.CLOSED;
    });
  };

  const isConnected = computed(() => {
    return status.value === WsStatusEnum.CONNECTED;
  });

  const send = async (data: Partial<Message>) => {
    if (!isConnected.value) return;
    const { currentUser } = useAuthStore();
    const message = {
      ...data,
      userId: currentUser.id,
    } as Message;
    ws.send({
      data: JSON.stringify(message),
      success() {
        console.log('send success');
      },
    });
  };

  const sendText = (data: Partial<Message>) => {
    send({
      ...data,
      media: MessageMediaEnum.TEXT,
    });
  };

  return { init, newMsg, sendText, isConnected, status };
});

export default useWsStore;
