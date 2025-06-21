import React from 'react';
import '../../styles/home.css';
import HeroSlider from '../components/HeroSlider';
 

function HomePage() {
  return (
    <div className="homepage">
       {/* 👈 thêm slider ở đầu trang */}

      <section className="hero">
        <h1>📱 Chào mừng đến với Telecom Shop</h1>
        <p>Khám phá các sản phẩm công nghệ tiên tiến, giá cực tốt!</p>
        <a href="/products" className="btn-shop">🛒 Xem sản phẩm</a>
      </section>
      <HeroSlider />

      <section className="ads-section">
        <div className="ad-card">
          <img src="/images/iphone15.jpg" alt="iPhone Promo" />
          <h3>🔥 Siêu phẩm iPhone 15</h3>
          <p>Giảm đến 2.000.000 VNĐ – số lượng có hạn!</p>
        </div>

        <div className="ad-card">
          <img src="/images/router.jpg" alt="Router Promo" />
          <h3>📡 Router tốc độ cao</h3>
          <p>Chỉ từ 899.000 VNĐ – Wifi khắp ngôi nhà bạn</p>
        </div>

        <div className="ad-card">
          <img src="/images/router.jpg" alt="Accessory Promo" />
          <h3>🎧 Phụ kiện công nghệ</h3>
          <p>Tai nghe, sạc nhanh, ốp lưng – giá siêu tốt</p>
        </div>
      </section>
    </div>
  );
}

export default HomePage;
