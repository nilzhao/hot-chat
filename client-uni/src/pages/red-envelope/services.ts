import request from '@/utils/request';
import { EnvelopeGoodsWithUser, EnvelopeGoodsItem } from './types';

export const reqFindEnvelope = () =>
  request.get<EnvelopeGoodsWithUser>('/goods/find');

// 收红包
export const reqReceiveEnvelopeItem = (envelopeNo: string) =>
  request.post<EnvelopeGoodsItem>('/goods/receive', {
    data: { envelopeNo },
  });
