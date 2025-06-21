import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import axios from 'axios';
import "../../styles/login.css";

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const res = await axios.post('http://localhost:8080/login', { email, password });

      console.log('Login response:', res.data); // Debug response

      if (res.data.token && res.data.user) {
        localStorage.setItem('token', res.data.token);
        localStorage.setItem('user', JSON.stringify(res.data.user));
        alert('Đăng nhập thành công');

        // Chuyển hướng theo role
        if (res.data.user.role === 'admin') {
          navigate('/admin');
        } else {
          navigate('/');
        }
      } else {
        alert('Phản hồi từ máy chủ không hợp lệ!');
      }
    } catch (err) {
      console.error('Login error:', err);
      alert('Sai thông tin đăng nhập');
    }
  };

  return (
    <div className="login-container">
      <h2 className="login-title">🔐 Đăng nhập</h2>
      <form className="login-form" onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Email:</label>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>
        <div className="form-group">
          <label>Mật khẩu:</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        <button type="submit">Đăng nhập</button>
      </form>
      <p className="login-links">
        <Link to="/register">👉 Chưa có tài khoản? Đăng ký</Link>
      </p>
      <p className="login-links">
        <a href="#">❓ Quên mật khẩu?</a>
      </p>
    </div>
  );
}

export default Login;
