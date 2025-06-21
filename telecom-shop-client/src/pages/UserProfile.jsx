import React, { useEffect, useState } from 'react';
import axios from 'axios';
import '../../styles/UserProfile.css';
import { FaUserCircle, FaEnvelope, FaIdBadge } from 'react-icons/fa';

function UserProfile() {
  const [user, setUser] = useState(null);
  const token = localStorage.getItem('token');

  useEffect(() => {
    fetchProfile();
  }, []);

  const fetchProfile = async () => {
    try {
      const res = await axios.get('http://localhost:8080/profile', {
        headers: { Authorization: `Bearer ${token}` }
      });
      setUser(res.data);
    } catch (err) {
      console.error('❌ Lỗi khi tải thông tin người dùng:', err);
    }
  };

  if (!user) return <p className="loading">🔄 Đang tải thông tin người dùng...</p>;

  return (
    <div className="profile-wrapper">
      <div className="profile-card">
        <div className="cover"></div>

        <div className="profile-content">
          <img
            className="avatar"
            src={`https://ui-avatars.com/api/?name=${user.name}&size=128&background=0D8ABC&color=fff`}
            alt="avatar"
          />
          <h2 className="name">{user.name}</h2>
          <p className="role">
            <FaUserCircle /> {user.role === 'admin' ? 'Quản trị viên' : 'Người dùng'}
          </p>

          <div className="info-box">
            <h3>📄 Thông tin cá nhân</h3>
            <ul>
              <li><FaEnvelope /> <strong>Email:</strong> {user.email}</li>
              <li><FaIdBadge /> <strong>Vai trò:</strong> {user.role}</li>
            </ul>
          </div>
        </div>
      </div>
    </div>
  );
}

export default UserProfile;
