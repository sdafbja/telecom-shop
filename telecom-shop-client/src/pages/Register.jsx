import React, { useState } from 'react';
import { useNavigate, Link } from 'react-router-dom';
import axios from 'axios';
import "../../styles/login.css";

function Register() {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await axios.post('http://localhost:8080/register', {
        name,
        email,
        password,
      });
      alert('Đăng ký thành công. Mời bạn đăng nhập!');
      navigate('/login');
    } catch (err) {
      alert('Lỗi đăng ký. Email có thể đã tồn tại.');
    }
  };

  return (
    <div className="login-container">
      <h2 className="login-title">📝 Đăng ký</h2>
      <form className="login-form" onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Họ tên:</label>
          <input type="text" value={name} onChange={(e) => setName(e.target.value)} required />
        </div>
        <div className="form-group">
          <label>Email:</label>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>
        <div className="form-group">
          <label>Mật khẩu:</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        <button type="submit">Đăng ký</button>
      </form>
      <p className="login-links">
        <Link to="/login">🔙 Đã có tài khoản? Đăng nhập</Link>
      </p>
    </div>
  );
}

export default Register;
