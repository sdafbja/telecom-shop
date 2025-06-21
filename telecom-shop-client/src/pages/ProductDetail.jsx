import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router-dom';
import axios from 'axios';
import '../../styles/product-detail.css';
import '../../styles/global.css';

function ProductDetail() {
  const { id } = useParams();
  const [product, setProduct] = useState(null);
  const [quantity, setQuantity] = useState(1);
  const [rating, setRating] = useState(5);
  const [comment, setComment] = useState('');
  const [reviews, setReviews] = useState([]);

  const token = localStorage.getItem('token');

  useEffect(() => {
    fetchProduct();
    fetchReviews();
  }, []);

  const fetchProduct = async () => {
    try {
      const res = await axios.get(`http://localhost:8080/products/${id}`);
      setProduct(res.data);
    } catch (err) {
      console.error('Lỗi khi lấy chi tiết sản phẩm:', err);
    }
  };

  const fetchReviews = async () => {
    try {
      const res = await axios.get(`http://localhost:8080/reviews/product/${id}`);
      setReviews(res.data);
    } catch (err) {
      console.error('Lỗi khi lấy đánh giá:', err);
    }
  };

  const handleAddToCart = async () => {
    try {
      if (!token) {
        alert("Bạn cần đăng nhập trước khi mua hàng!");
        return;
      }

      await axios.post(
        'http://localhost:8080/cart',
        {
          product_id: product.id,
          quantity: quantity,
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
            'Content-Type': 'application/json',
          },
        }
      );

      alert('✅ Đã thêm vào giỏ hàng!');
    } catch (err) {
      console.error('Lỗi thêm vào giỏ:', err);
      alert('❌ Thêm vào giỏ thất bại!');
    }
  };

  const submitReview = async () => {
    try {
      if (!token) {
        alert("Bạn cần đăng nhập để đánh giá!");
        return;
      }

      await axios.post(
        'http://localhost:8080/reviews',
        {
          product_id: product.id,
          rating,
          comment,
        },
        {
          headers: { Authorization: `Bearer ${token}` },
        }
      );

      alert("✅ Gửi đánh giá thành công!");
      setComment('');
      setRating(5);
      fetchReviews();
    } catch (err) {
      console.error("Lỗi khi gửi đánh giá:", err);
      alert("❌ Gửi đánh giá thất bại!");
    }
  };

  if (!product) return <div>Đang tải...</div>;

    return (
    <div className="product-detail-container">
      <div className="product-detail">
        <div className="product-image">
          <img
            src={product.image_url}
            alt={product.name}
            onError={(e) => {
              e.target.src = 'https://via.placeholder.com/300x300?text=No+Image';
            }}
          />
        </div>

        <div className="product-info">
          <h2>{product.name}</h2>
          <p className="price">{Number(product.price).toLocaleString()} VNĐ</p>
          <p className="description">{product.description || 'Chưa có mô tả.'}</p>

          <div className="quantity-control">
            <label>Số lượng:</label>
            <input
              type="number"
              min="1"
              value={quantity}
              onChange={(e) => setQuantity(parseInt(e.target.value))}
            />
          </div>

          <button className="btn-add" onClick={handleAddToCart}>
            🛒 Thêm vào giỏ hàng
          </button>
        </div>
      </div>

      {/* Đánh giá nằm dưới sản phẩm */}
      <div className="review-section">
        <h3>⭐ Đánh giá sản phẩm</h3>

        <div className="review-form">
          <label>Số sao:</label>
          <select value={rating} onChange={(e) => setRating(parseInt(e.target.value))}>
            {[1, 2, 3, 4, 5].map((r) => (
              <option key={r} value={r}>{r} sao</option>
            ))}
          </select>

          <textarea
            placeholder="Viết bình luận..."
            value={comment}
            onChange={(e) => setComment(e.target.value)}
          />
          <button onClick={submitReview}>Gửi đánh giá</button>
        </div>

        <ul className="review-list">
          {reviews.length === 0 && <p>Chưa có đánh giá nào.</p>}
          {reviews.map((r) => (
            <li key={r.id}>
              <strong>• {r.user?.name || 'Ẩn danh'}: {r.rating} ⭐</strong>
              <p>{r.comment}</p>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default ProductDetail;
