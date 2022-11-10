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
  cmd: number;
  targetId: number;
  media: MessageMediaEnum;
  content: string;
  pic: string;
  url: string;
  memo: string;
  amount: number;
  createdAt: string;
}

export interface Chat {
  note: string;
  updatedAt: string;
  avatar: string;
  name: string;
  targetId: number;
}
