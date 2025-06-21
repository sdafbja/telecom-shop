// src/pages/MyOrders.jsx
import React, { useEffect, useState } from 'react';
import axios from 'axios';
import '../../styles/orders.css'; // bạn có thể tạo file CSS này

function MyOrders() {
  const [orders, setOrders] = useState([]);
  const token = localStorage.getItem('token');

  useEffect(() => {
    fetchOrders();
  }, []);

  const fetchOrders = async () => {
    try {
      const res = await axios.get('http://localhost:8080/orders', {
        headers: { Authorization: `Bearer ${token}` },
      });
      setOrders(res.data);
    } catch (err) {
      console.error('❌ Lỗi khi lấy đơn hàng:', err);
    }
  };

  return (
    <div className="orders-container">
      <h2>📦 Đơn hàng của bạn</h2>

      {orders.length === 0 ? (
        <p>Bạn chưa có đơn hàng nào.</p>
      ) : (
        orders.map((order) => (
          <div key={order.id} className="order-card">
            <div className="order-info">
              <p><strong>Người nhận:</strong> {order.name}</p>
              <p><strong>Điện thoại:</strong> {order.phone}</p>
              <p><strong>Địa chỉ:</strong> {order.address}</p>
              <p><strong>Ghi chú:</strong> {order.note || "Không có"}</p>
              <p><strong>Trạng thái:</strong> {order.status}</p>
              <p><strong>Thời gian:</strong> {new Date(order.created_at).toLocaleString()}</p>
              <p><strong>Tổng tiền:</strong> {Number(order.total).toLocaleString()} VNĐ</p>
            </div>

            <div className="order-items">
              <h4>Sản phẩm:</h4>
              <ul>
                {order.items.map((item) => (
                  <li key={item.id} className="item">
                    <img
                      src={item.product?.image_url}
                      alt={item.product?.name}
                      onError={(e) => {
                        e.target.src = 'https://via.placeholder.com/50';
                      }}
                    />
                    <div>
                      <p>{item.product?.name}</p>
                      <p>Số lượng: {item.quantity}</p>
                      <p>Giá: {Number(item.price).toLocaleString()} VNĐ</p>
                    </div>
                  </li>
                ))}
              </ul>
            </div>
          </div>
        ))
      )}
    </div>
  );
}

export default MyOrders;
