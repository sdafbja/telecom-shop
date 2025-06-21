import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { Bar } from 'react-chartjs-2';
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  BarElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js';
import '../../styles/admin-dashboard.css';

ChartJS.register(CategoryScale, LinearScale, BarElement, Title, Tooltip, Legend);

function AdminDashboard() {
  const [stats, setStats] = useState({ users: 0, orders: 0, revenue: 0 });

  useEffect(() => {
    fetchStats();
  }, []);

  const fetchStats = async () => {
    try {
      const token = localStorage.getItem('token');
      const res = await axios.get('http://localhost:8080/admin/stats', {
        headers: { Authorization: `Bearer ${token}` },
      });
      setStats(res.data);
    } catch (err) {
      console.error('❌ Lỗi khi lấy thống kê:', err);
    }
  };

  const data = {
    labels: ['Người dùng', 'Đơn hàng', 'Doanh thu'],
    datasets: [
      {
        label: 'Thống kê',
        data: [stats.users, stats.orders, stats.revenue],
        backgroundColor: ['#4caf50', '#2196f3', '#ff9800'],
      },
    ],
  };

  return (
    <div className="dashboard">
      <h2>📊 Thống kê Quản trị</h2>
      <div className="summary">
        <div className="card">👥 Người dùng: <strong>{stats.users}</strong></div>
        <div className="card">📦 Đơn hàng: <strong>{stats.orders}</strong></div>
        <div className="card">💰 Doanh thu: <strong>{stats.revenue.toLocaleString()} VNĐ</strong></div>
      </div>
      <div className="chart">
        <Bar data={data} />
      </div>
    </div>
  );
}

export default AdminDashboard;
