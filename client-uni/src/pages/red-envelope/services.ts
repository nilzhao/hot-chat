import request from '@/utils/request';
import {
  EnvelopeGoodsWithUser,
  EnvelopeGoodsItem,
  EnvelopeGoodsItemWithUser,
  EnvelopeGoodsTypeEnum,
} from './types';

// 红包详情
export const reqDetailEnvelope = (envelopeNo: string) =>
  request.get<EnvelopeGoodsWithUser>(`/goods/${envelopeNo}`);

export const reqFindEnvelope = () =>
  request.get<EnvelopeGoodsWithUser>('/goods/find');

// 收红包
export const reqReceiveEnvelopeItem = (envelopeNo: string) =>
  request.post<EnvelopeGoodsItem>('/goods/receive', {
    data: { envelopeNo },
  });

// 红包领取详情
export const reqEnvelopeItems = (envelopeNo: string) =>
  request.get<EnvelopeGoodsItemWithUser[]>(`/goods/${envelopeNo}/items`);

// 红包领取详情
export const reqSendOutEnvelope = (data: {
  type?: EnvelopeGoodsTypeEnum;
  amount?: number;
  amountOne?: number;
  quantity?: number;
}) =>
  request.post<EnvelopeGoodsWithUser[]>(`/goods/sendOut`, {
    data,
  });
