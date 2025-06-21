import React, { useEffect, useState } from 'react';
import axios from 'axios';

function OrderList() {
  const [orders, setOrders] = useState([]);

  useEffect(() => {
    fetchOrders();
  }, []);

  const fetchOrders = async () => {
    const token = localStorage.getItem('token');
    const res = await axios.get('http://localhost:8080/orders', {
      headers: { Authorization: `Bearer ${token}` },
    });
    setOrders(res.data);
  };

  return (
    <div className="order-list">
      <h2>📦 Đơn hàng của bạn</h2>
      {orders.map(order => (
        <div key={order.id} className="order">
          <p><strong>{order.name}</strong> - {order.phone}</p>
          <p>Địa chỉ: {order.address}</p>
          <p>Ghi chú: {order.note}</p>
          <p>Tổng tiền: {order.total.toLocaleString()} VNĐ</p>
          <p>Trạng thái: {order.status}</p>
        </div>
      ))}
    </div>
  );
}

export default OrderList;
