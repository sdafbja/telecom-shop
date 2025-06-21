import React from 'react';
import { Swiper, SwiperSlide } from 'swiper/react';
import { Autoplay } from 'swiper/modules';
import 'swiper/css';
import 'swiper/css/autoplay';
import "../../styles/HeroSlider.css";

function HeroSlider() {
  return (
    <div className="hero-slider">
      <Swiper

      
        modules={[Autoplay]}  // 👈 đúng cách mới để dùng module
        autoplay={{ delay: 3000 }}
        loop={true}
      >
        {slides.map((slide) => (
          <SwiperSlide key={slide.id}>
            <img src={slide.img} alt={slide.caption} className="slider-img" />
            <div className="caption">{slide.caption}</div>
          </SwiperSlide>
        ))}
      </Swiper>
    </div>
  );
}

const slides = [
  {
    id: 1,
    img: '/images/sale.jpg',
    caption: '🔥 Giảm sốc cuối tuần',
  },
  {
    id: 2,
    img: '/images/giaohang.png',
    caption: '⚡ Giao hàng hoả tốc 2H',
  },
  {
    id: 3,
    img: '/images/combo.jpg',
    caption: '🎧 Combo phụ kiện giá hời',
  },
];



export default HeroSlider;
