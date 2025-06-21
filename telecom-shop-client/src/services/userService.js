// src/services/userService.js
import API from './api';

export const fetchUsers = async () => {
  const res = await API.get('/admin/users');
  return res.data;
};

export const deleteUser = async (id) => {
  await API.delete(`/admin/users/${id}`);
};

export const updateUser = async (id, data) => {
  const res = await API.put(`/admin/users/${id}`, data);
  return res.data;
};
