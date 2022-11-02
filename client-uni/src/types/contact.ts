import { User } from './user';

export enum ContactTypeEnum {
  SINGLE = 0,
  COMMUNITY,
}

export interface Contact {
  id: number;
  ownerId: number;
  targetId: number;
  type: ContactTypeEnum;
  memo: string;
}

export interface ContactDetail extends Contact {
  targetUser: User;
}
