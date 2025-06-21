import React, { useEffect, useState } from 'react';
import axios from 'axios';
import '../../styles/product.css';
import { useNavigate } from 'react-router-dom';

function Product() {
  const [products, setProducts] = useState([]);
  const [categories, setCategories] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    fetchCategories();
    fetchProducts();
  }, []);

  const fetchCategories = async () => {
    try {
      const res = await axios.get('http://localhost:8080/categories');
      setCategories(res.data);
    } catch (err) {
      console.error('❌ Lỗi khi lấy danh mục:', err);
    }
  };

  const fetchProducts = async (categoryId = '') => {
    try {
      const url = categoryId
        ? `http://localhost:8080/products?category_id=${categoryId}`
        : 'http://localhost:8080/products';
      const res = await axios.get(url);
      setProducts(res.data);
    } catch (err) {
      console.error('❌ Lỗi khi lấy sản phẩm:', err);
    }
  };

  const handleCategoryChange = (e) => {
    const categoryId = e.target.value;
    setSelectedCategory(categoryId);
    fetchProducts(categoryId);
  };

  return (
    <div className="product-page container">
      <div className="header-row">
        <h2 className="product-title">🛒 Sản phẩm nổi bật</h2>
        <select
          value={selectedCategory}
          onChange={handleCategoryChange}
          className="category-select"
        >
          <option value="">Tất cả danh mục</option>
          {categories.map((cat) => (
            <option key={cat.id} value={cat.id}>
              {cat.name}
            </option>
          ))}
        </select>
      </div>

      <div className="product-grid">
        {products.map((p) => (
          <div
            key={p.id}
            className="product-card"
            onClick={() => navigate(`/products/${p.id}`)}
          >
            <div className="product-img-wrapper">
              <img
                src={p.image_url}
                alt={p.name}
                className="product-img"
                onError={(e) => {
                  e.target.src = 'https://via.placeholder.com/240x240?text=No+Image';
                }}
              />
            </div>
            <div className="product-info">
              <h3 className="product-name">{p.name}</h3>
              <p className="product-price">
                {Number(p.price).toLocaleString()} <span>VNĐ</span>
              </p>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}

export default Product;
