// src/pages/AdminPage.jsx
import React, { useEffect, useState } from 'react';
import { fetchUsers, deleteUser, updateUser } from '../services/userService';

function AdminPage() {
  const [users, setUsers] = useState([]);
  const [editingUser, setEditingUser] = useState(null);

  useEffect(() => {
    loadUsers();
  }, []);

  const loadUsers = async () => {
    const data = await fetchUsers();
    setUsers(data);
  };

  const handleDelete = async (id) => {
    if (window.confirm('Bạn có chắc chắn muốn xóa người dùng này?')) {
      await deleteUser(id);
      loadUsers();
    }
  };

  const handleEdit = (user) => {
    setEditingUser({ ...user });
  };

  const handleSave = async () => {
    await updateUser(editingUser.id, editingUser);
    setEditingUser(null);
    loadUsers();
  };

  return (
    <div style={{ maxWidth: 800, margin: 'auto' }}>
      <h2>👑 Quản lý người dùng</h2>
      {users.map((u) => (
        <div key={u.id} style={{ marginBottom: 10, borderBottom: '1px solid #ccc', paddingBottom: 10 }}>
          {editingUser?.id === u.id ? (
            <>
              <input
                value={editingUser.name}
                onChange={(e) => setEditingUser({ ...editingUser, name: e.target.value })}
              />
              <input
                value={editingUser.email}
                onChange={(e) => setEditingUser({ ...editingUser, email: e.target.value })}
              />
              <select
                value={editingUser.role}
                onChange={(e) => setEditingUser({ ...editingUser, role: e.target.value })}
              >
                <option value="user">user</option>
                <option value="admin">admin</option>
              </select>
              <button onClick={handleSave}>💾 Lưu</button>
              <button onClick={() => setEditingUser(null)}>❌ Hủy</button>
            </>
          ) : (
            <>
              <span>👤 {u.name} ({u.email}) - {u.role}</span>
              <button onClick={() => handleEdit(u)}>✏️ Sửa</button>
              <button onClick={() => handleDelete(u.id)}>🗑️ Xóa</button>
            </>
          )}
        </div>
      ))}
    </div>
  );
}

export default AdminPage;
