export enum UserStatusEnum {
  NORMAL = 0,
  DISABLED,
}

export enum UserGenderEnum {
  UNKNOWN = 0,
  MALE,
  FEMALE,
}

export interface User {
  id: number;
  status: UserStatusEnum;
  email: string;
  name: string;
  avatar: string;
  gender: UserGenderEnum;
}

export const INIT_USER: User = {
  id: 0,
  status: UserStatusEnum.NORMAL,
  email: '',
  name: '',
  avatar: '',
  gender: UserGenderEnum.UNKNOWN,
};
