// src/components/ConfirmOrder.jsx
import React, { useState } from 'react';
import axios from 'axios';

function ConfirmOrder({ total, onSuccess }) {
  const [info, setInfo] = useState({
    name: '',
    phone: '',
    address: '',
    note: '',
  });

  const handleChange = (e) => {
    setInfo({ ...info, [e.target.name]: e.target.value });
  };

  const handleSubmit = async () => {
    try {
      const token = localStorage.getItem('token');
      await axios.post(
        'http://localhost:8080/orders',
        { ...info, total },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      alert('✅ Đặt hàng thành công!');
      onSuccess();
    } catch (err) {
      alert('❌ Đặt hàng thất bại!');
      console.error(err);
    }
  };

  return (
    <div className="confirm-order">
      <h2>🔒 Xác nhận đơn hàng</h2>
      <input name="name" placeholder="Tên người nhận" onChange={handleChange} />
      <input name="phone" placeholder="Số điện thoại" onChange={handleChange} />
      <input name="address" placeholder="Địa chỉ giao hàng" onChange={handleChange} />
      <textarea name="note" placeholder="Ghi chú thêm" onChange={handleChange}></textarea>
      <button onClick={handleSubmit}>🛒 Xác nhận & Thanh toán</button>
    </div>
  );
}

export default ConfirmOrder;
