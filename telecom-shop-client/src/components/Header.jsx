import React, { useState } from 'react';
import { Link, useNavigate, useLocation } from 'react-router-dom';
import '../../styles/Header.css';

function Header() {
  const navigate = useNavigate();
  const location = useLocation();
  const token = localStorage.getItem('token');
  const [searchQuery, setSearchQuery] = useState('');

  let user = null;
  try {
    const userString = localStorage.getItem('user');
    user = userString ? JSON.parse(userString) : null;
  } catch (error) {
    console.error("Lỗi parse user:", error);
  }

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    navigate('/login');
  };

  const handleSearch = (e) => {
    e.preventDefault();
    if (searchQuery.trim()) {
      navigate(`/search?query=${encodeURIComponent(searchQuery.trim())}`);
    }
  };

  return (
    <header className="main-header">
      <div className="container">
        <div className="logo">
          <Link to="/">🛒 <span>Telecom Shop</span></Link>
        </div>

        <form onSubmit={handleSearch} className="search-bar">
          <input
            type="text"
            placeholder="🔍 Tìm sản phẩm..."
            value={searchQuery}
            onChange={(e) => setSearchQuery(e.target.value)}
          />
          <button type="submit">Tìm</button>
        </form>

        <nav className="nav-links">
          <Link to="/">Trang chủ</Link>
          <Link to="/products">Sản phẩm</Link>
          <Link to="/cart">Giỏ hàng</Link>

          {user?.role === 'admin' && (
            <>
              <Link to="/admin/products">⚙️ Quản lý SP</Link>
              <Link to="/admin">👑 Admin</Link>
            </>
          )}

          {token && user ? (
            <>
              <Link to="/profile" className="username">👤 {user.name}</Link>
              <button onClick={handleLogout} className="logout-btn">Đăng xuất</button>
            </>
          ) : (
            <Link to="/login" className="login-btn">Đăng nhập</Link>
          )}
        </nav>
      </div>
    </header>
  );
}

export default Header;
