export enum UserStatusEnum {
  NORMAL = 0,
  DISABLED,
}

export interface User {
  id: number;
  status: UserStatusEnum;
  email: string;
  name: string;
  avatar: string;
}

export const INIT_USER: User = {
  id: 0,
  status: UserStatusEnum.NORMAL,
  email: '',
  name: '',
  avatar: '',
};
