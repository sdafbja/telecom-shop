import React, { useEffect, useState } from 'react';
import axios from 'axios';
import '../../styles/cart.css';
import '../../styles/global.css';

function CartPage() {
  const [cartItems, setCartItems] = useState([]);
  const [showForm, setShowForm] = useState(false);
  const [orderInfo, setOrderInfo] = useState({
    name: '',
    phone: '',
    address: '',
    note: '',
  });
  const token = localStorage.getItem('token');

  useEffect(() => {
    fetchCart();
  }, []);

  const fetchCart = async () => {
    try {
      const res = await axios.get('http://localhost:8080/cart', {
        headers: { Authorization: `Bearer ${token}` },
      });
      setCartItems(res.data);
    } catch (err) {
      console.error('❌ Lỗi khi lấy giỏ hàng:', err);
    }
  };

  const updateQuantity = async (itemId, newQuantity) => {
    if (newQuantity < 1) return;
    try {
      await axios.put(
        `http://localhost:8080/cart/${itemId}`,
        { quantity: newQuantity },
        { headers: { Authorization: `Bearer ${token}` } }
      );
      fetchCart();
    } catch (err) {
      console.error('❌ Lỗi khi cập nhật số lượng:', err);
    }
  };

  const removeItem = async (id) => {
    if (!window.confirm('Xóa sản phẩm này khỏi giỏ?')) return;
    try {
      await axios.delete(`http://localhost:8080/cart/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      fetchCart();
    } catch (err) {
      console.error('❌ Lỗi khi xoá sản phẩm:', err);
    }
  };

  const calculateTotal = () => {
    return cartItems.reduce((total, item) => {
      return total + item.quantity * (item.product?.price || 0);
    }, 0);
  };

  const handleSubmitOrder = async () => {
    if (!orderInfo.name || !orderInfo.phone || !orderInfo.address) {
      alert('Vui lòng điền đầy đủ thông tin.');
      return;
    }

    try {
      const total = calculateTotal(); // ✅ Thêm dòng này để tính tổng
      const payload = { ...orderInfo, total }; // ✅ Gộp orderInfo với tổng tiền

      await axios.post('http://localhost:8080/orders', payload, {
        headers: { Authorization: `Bearer ${token}` },
      });

      alert('✅ Đặt hàng thành công!');
      setShowForm(false);
      setOrderInfo({ name: '', phone: '', address: '', note: '' });
      fetchCart(); // làm trống lại giỏ hàng
    } catch (err) {
      console.error('❌ Lỗi khi đặt hàng:', err);
      alert('❌ Đặt hàng thất bại!');
    }
  };


  return (
    <div className="cart-container">
      <h2 className="cart-title">🛒 Giỏ hàng của bạn</h2>
      {cartItems.length === 0 ? (
        <p>Không có sản phẩm nào trong giỏ hàng.</p>
      ) : (
        <>
          <table className="cart-table">
            <thead>
              <tr>
                <th>Sản phẩm</th>
                <th>Số lượng</th>
                <th>Giá</th>
                <th>Thao tác</th>
              </tr>
            </thead>
            <tbody>
              {cartItems.map((item) => (
                <tr key={item.id}>
                  <td className="product-info">
                    <img
                      src={item.product?.image_url}
                      alt={item.product?.name}
                      onError={(e) => {
                        e.target.src = 'https://via.placeholder.com/60x60';
                      }}
                    />
                    <span>{item.product?.name}</span>
                  </td>
                  <td>
                    <div className="qty-control">
                      <button onClick={() => updateQuantity(item.id, item.quantity - 1)} disabled={item.quantity <= 1}>−</button>
                      <span>{item.quantity}</span>
                      <button onClick={() => updateQuantity(item.id, item.quantity + 1)}>＋</button>
                    </div>
                  </td>
                  <td>{Number(item.product?.price).toLocaleString()} VNĐ</td>
                  <td>
                    <button onClick={() => removeItem(item.id)}>🗑️ Xoá</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>

          <div className="cart-summary">
            <h3>Tổng cộng: {calculateTotal().toLocaleString()} VNĐ</h3>
            <button className="checkout-btn" onClick={() => setShowForm(true)}>🚚 Đặt hàng</button>
          </div>
        </>
      )}

      {showForm && (
        <div className="order-form">
          <h3>📦 Thông tin nhận hàng</h3>
          <input
            type="text"
            placeholder="Họ tên"
            value={orderInfo.name}
            onChange={(e) => setOrderInfo({ ...orderInfo, name: e.target.value })}
            required
          />
          <input
            type="tel"
            placeholder="Số điện thoại"
            value={orderInfo.phone}
            onChange={(e) => setOrderInfo({ ...orderInfo, phone: e.target.value })}
            required
          />
          <input
            type="text"
            placeholder="Địa chỉ"
            value={orderInfo.address}
            onChange={(e) => setOrderInfo({ ...orderInfo, address: e.target.value })}
            required
          />
          <textarea
            placeholder="Ghi chú (tuỳ chọn)"
            value={orderInfo.note}
            onChange={(e) => setOrderInfo({ ...orderInfo, note: e.target.value })}
          ></textarea>
          <button className="confirm-btn" onClick={handleSubmitOrder}>✅ Xác nhận đặt hàng</button>
          <button className="cancel-btn" onClick={() => setShowForm(false)}>❌ Hủy</button>
        </div>
      )}
    </div>
  );
}

export default CartPage;
