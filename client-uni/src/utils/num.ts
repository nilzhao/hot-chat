export const formatAmount = (amount?: number | string): string => {
  if (!amount) return '0.00';
  let amountStr = String(amount);
  const [int, decemal = ''] = amountStr.split('.');

  return `${int}.${decemal.padEnd(2, '0')}`;
};
