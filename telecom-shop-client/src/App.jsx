import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Product from './pages/Product';
import Login from './pages/Login';
import Header from './components/Header';
import Footer from './components/Footer';
import Register from './pages/Register';
import AdminPage from './pages/AdminPage';
import AdminProducts from './pages/AdminProducts';
import CartPage from './pages/CartPage';
import SearchView from './pages/SearchView';
import HomePage from './pages/HomePage';
import ProductDetail from './pages/ProductDetail';
import UserProfile from './pages/UserProfile';
import HeroSlider from './components/HeroSlider';
import AdminDashboard from './pages/AdminDashboard';
import ConfirmOrder from './pages/ConfirmOrder';
import OrderList from './pages/OrderList';
import MyOrders from './pages/MyOrders';

function App() {
  return (
    <Router>
      <div className="layout">
        <Header />
        <main className="main-content">
          <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/products" element={<Product />} />
            <Route path="/products/:id" element={<ProductDetail />} />
            <Route path="/login" element={<Login />} />
            <Route path="/register" element={<Register />} />
            <Route path="/admin" element={<AdminPage />} />
            <Route path="/admin/products" element={<AdminProducts />} />
            <Route path="/admin/dashboard" element={<AdminDashboard />} />
            <Route path="/cart" element={<CartPage />} />
            <Route path="/search" element={<SearchView />} />
            <Route path="/profile" element={<UserProfile />} />
            <Route path="/heroslider" element={<HeroSlider />} />
            <Route path="/confirm-order" element={<ConfirmOrder />} />
            <Route path="/orders" element={<OrderList />} />
            <Route path="/my-orders" element={<MyOrders />} />
            {/* Redirect to home if no match */}
          </Routes>
        </main>
        <Footer />
      </div>
    </Router>
  );
}

export default App;
