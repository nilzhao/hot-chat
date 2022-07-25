const parseDate = (time?: string | null) => {
  let date = new Date();
  if (time) {
    date = new Date(time);
  }
  const y = date.getFullYear();
  const M = date.getMonth() + 1;
  const d = date.getDate();
  const h = date.getHours();
  const m = date.getMinutes();
  const s = date.getSeconds();
  return [y, M, d, h, m, s];
};

export const formatDate = (time?: string | null) => {
  const times = parseDate(time);
  return times.slice(0, 3).join('-');
};

export const formatDateTime = (time?: string | null) => {
  const times = parseDate(time);
  return times.slice(0, 3).join('-') + ' ' + times.slice(3).join(':');
};
