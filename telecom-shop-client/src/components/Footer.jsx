import React from 'react';
import '../../styles/Footer.css';

function Footer() {
  const currentYear = new Date().getFullYear();

  return (
    <footer className="footer">
      <div className="footer-container">
        {/* Cột 1: Giới thiệu */}
        <div className="footer-col">
          <h2 className="footer-logo">Telecom<span>Shop</span></h2>
          <p className="footer-desc">
            Cửa hàng công nghệ chuyên cung cấp thiết bị chính hãng, chất lượng và dịch vụ tận tâm.
          </p>
        </div>

        {/* Cột 2: Liên kết */}
        <div className="footer-col">
          <h4>Dịch vụ</h4>
          <ul>
            <li><a href="#">Sản phẩm mới</a></li>
            <li><a href="#">Hướng dẫn mua hàng</a></li>
            <li><a href="#">Thanh toán</a></li>
            <li><a href="#">Đơn hàng của tôi</a></li>
          </ul>
        </div>

        {/* Cột 3: Đăng ký và mạng xã hội */}
        <div className="footer-col">
          <h4>Nhận tin mới</h4>
          <div className="newsletter">
            <input type="email" placeholder="Email của bạn" />
            <button>Gửi</button>
          </div>

          <h4>Kết nối với chúng tôi</h4>
          <div className="social-icons">
            <a href="#"><img src="/images/fb.jpg" alt="Facebook" /></a>
            <a href="#"><img src="/images/ins.jpg" alt="Instagram" /></a>
            <a href="#"><img src="/images/tw.jpg" alt="Twitter" /></a>
          </div>
        </div>
      </div>

      <div className="footer-bottom">
        © {currentYear} Telecom Shop. Thiết kế bởi bạn 💙
      </div>
    </footer>
  );
}

export default Footer;
