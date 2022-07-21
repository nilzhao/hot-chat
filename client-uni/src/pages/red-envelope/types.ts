import { User } from '@/types/user';

export interface EnvelopeGoods {
  accountNo: string;
  amount: string;
  amountOne: string;
  blessing: '';
  createdAt: string;
  deletedAt: string | null;
  envelopeNo: string;
  expired: string;
  id: number;
  orderType: number;
  originEnvelopeNo: string;
  payStatus: number;
  quantity: number;
  remainAmount: string;
  remainQuantity: number;
  status: number;
  type: number;
  updatedAt: string;
  userId: number;
  username: string;
}

export interface EnvelopeGoodsWithUser extends EnvelopeGoods {
  user: User;
}

export interface EnvelopeGoodsItem {
  accountNo: string;
  amount: string;
  createdAt: string;
  deletedAt: string | null;
  desc: string;
  envelopNo: string;
  id: number;
  itemNo: string;
  payStatus: number;
  recvUserId: number;
  recvUsername: string;
  remainAmount: string;
  updatedAt: string;
}
