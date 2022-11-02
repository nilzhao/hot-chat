import { STORAGE_KEYS } from '@/config';
import { Message, MessageMediaEnum } from '@/types/chat';
import { defineStore } from 'pinia';
import { computed, ref } from 'vue';

export enum WsStatusEnum {
  WAITING = 'waiting',
  CONNECTED = 'connected',
  ERROR = 'error',
  CLOSED = 'closed',
}

const useWsStore = defineStore('ws', () => {
  const status = ref<WsStatusEnum>(WsStatusEnum.WAITING);
  const newMsg = ref<Message | null>(null);
  let ws: WebSocket;

  const init = (url?: string) => {
    if (ws) return;
    ws = new WebSocket(
      url ||
        `ws://127.0.0.1:9000${__API_PREFIX__}/ws?X-Token=${uni.getStorageSync(
          STORAGE_KEYS.token
        )}`
    );

    ws.addEventListener('open', () => {
      status.value = WsStatusEnum.CONNECTED;
    });

    ws.addEventListener('error', (e: WebSocketEventMap['error']) => {
      status.value = WsStatusEnum.ERROR;
      console.dir(e);
    });

    ws.addEventListener('message', (e: WebSocketEventMap['message']) => {
      newMsg.value = JSON.parse(e.data);
    });

    ws.addEventListener('close', () => {
      status.value = WsStatusEnum.CLOSED;
    });
  };

  const isConnected = computed(() => {
    return status.value === WsStatusEnum.CONNECTED;
  });

  const send = (data: Partial<Message>) => {
    if (!isConnected.value) return;
    ws.send(JSON.stringify(data));
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
