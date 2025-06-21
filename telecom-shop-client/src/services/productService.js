import API from './api';

export const fetchProducts = async () => {
  const res = await API.get('/products');
  return res.data;
};
