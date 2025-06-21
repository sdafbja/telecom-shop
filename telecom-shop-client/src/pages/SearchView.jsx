import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useLocation } from 'react-router-dom';
import "../../styles/home.css";
import '../../styles/global.css';

function SearchView() {
  const location = useLocation();
  const [products, setProducts] = useState([]);
  const query = new URLSearchParams(location.search).get("query");

  useEffect(() => {
    if (query && query.trim() !== '') {
      fetchProducts(query);
    }
  }, [query]);

  const fetchProducts = async (keyword) => {
    try {
      const res = await axios.get(`http://localhost:8080/products?search=${encodeURIComponent(keyword)}`);
      setProducts(res.data);
    } catch (err) {
      console.error('❌ Lỗi khi tìm kiếm:', err);
    }
  };

  return (
    <div className="home">
      <h2>🔍 Kết quả tìm kiếm cho: <em>{query}</em></h2>
      <div className="products">
        {products.map((p) => (
          <div className="product-card" key={p.id}>
            <img src={p.image_url} alt={p.name} />
            <h3>{p.name}</h3>
            <p>{Number(p.price).toLocaleString()} VNĐ</p>
          </div>
        ))}
      </div>
    </div>
  );
}

export default SearchView;
