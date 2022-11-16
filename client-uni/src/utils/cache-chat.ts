import { CHAT_MESSAGE_LIMIT } from '@/config';
import { Chat, Message, MessageCmdEnum, MessageMediaEnum } from '@/types/chat';
import { ContactDetail } from '@/types/contact';
import { User } from '@/types/user';
import { getLocalIItem, setLocalItem } from './storage';

class CacheChat {
  chatListKey = `chat-list`;
  messageListKeyPrefix = `message-list`;

  currentUser: User;
  contactsMap: Map<number, ContactDetail>;

  constructor(user: User, contactsMap: Map<number, ContactDetail>) {
    this.currentUser = user;
    this.contactsMap = contactsMap;
  }

  getTargetId = (...ids: number[]) => {
    return ids.find((id) => this.currentUser.id !== id) || 0;
  };

  private generateMessageListKey = (
    cmd: MessageCmdEnum,
    userId: number,
    targetId: number
  ) => {
    return `message-list-${cmd}-${[userId, targetId].sort().join('-')}`;
  };

  cacheOnReceivedMessage = async (message: Message) => {
    await this.cacheMessage(message);
    await this.cacheChatList(message);
  };

  cacheMessage = async (message: Message) => {
    const key = this.generateMessageListKey(
      message.cmd,
      message.userId,
      message.targetId
    );
    const cachedList = await getLocalIItem<Message[]>(key);
    let newList: Message[] = [...(cachedList || []), message];
    await setLocalItem(key, newList.slice(-CHAT_MESSAGE_LIMIT));
  };

  getChatNote = (message: Message) => {
    switch (message.media) {
      case MessageMediaEnum.IMG:
        return '[图片]';
      case MessageMediaEnum.VIDEO:
        return '[视频]';
      default:
        return message.content;
    }
  };

  refreshChatList = async (message: Message) => {
    const cachedList = await this.getChatList();
    const targetId = this.getTargetId(message.userId, message.targetId);
    const targetUser = this.contactsMap.get(targetId)?.targetUser;
    if (!targetUser) return [];

    const newChat: Chat = {
      note: this.getChatNote(message),
      updatedAt: message.createdAt,
      avatar: targetUser.avatar,
      name: targetUser.name,
      targetId: targetUser.id,
    };
    return [
      newChat,
      ...cachedList.filter((chat) => {
        return chat.targetId !== targetId;
      }),
    ];
  };

  cacheChatList = async (message: Message) => {
    const newList = await this.refreshChatList(message);
    await setLocalItem(this.chatListKey, newList);
  };

  removeChat = async (targetId: number) => {
    const cachedList = await this.getChatList();
    await setLocalItem(
      this.chatListKey,
      cachedList.filter((chat) => {
        return chat.targetId !== targetId;
      })
    );
  };

  getChatList = async () => {
    return (await getLocalIItem<Chat[]>(this.chatListKey)) || [];
  };

  getTargetMessageList = async (
    cmd: MessageCmdEnum,
    targetId: number | undefined
  ): Promise<Message[]> => {
    if (!targetId) return [];
    const messageList = await getLocalIItem<Message[]>(
      this.generateMessageListKey(cmd, this.currentUser.id, targetId)
    );
    return messageList || [];
  };
}

export default CacheChat;
