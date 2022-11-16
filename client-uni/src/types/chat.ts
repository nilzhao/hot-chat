export enum MessageCmdEnum {
  HEART = 0,
  SINGLE,
  ROOM,
}
export enum MessageMediaEnum {
  TEXT = 0,
  NEWS,
  VOICE,
  IMG,
  VIDEO,
  MUSIC,
}

export interface Message {
  id: string;
  userId: number;
  cmd: MessageCmdEnum;
  targetId: number;
  media: MessageMediaEnum;
  content: string;
  pic: string;
  memo: string;
  amount: number;
  createdAt: string;
  width?: number;
  height?: number;
}

export interface Chat {
  note: string;
  updatedAt: string;
  avatar: string;
  name: string;
  targetId: number;
}
