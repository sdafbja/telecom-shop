import React, { useEffect, useState } from 'react';
import axios from 'axios';

function AdminProducts() {
  const [products, setProducts] = useState([]);
  const [editingProduct, setEditingProduct] = useState(null);
  const [form, setForm] = useState({
    name: '',
    description: '',
    price: '',
    imageFile: null,
    category_id: '',
  });

  const token = localStorage.getItem('token');

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = async () => {
    try {
      const res = await axios.get('http://localhost:8080/products');
      setProducts(res.data);
    } catch (err) {
      console.error('❌ Lỗi khi lấy sản phẩm:', err);
    }
  };

  const handleChange = (e) => {
    const { name, value, files } = e.target;
    if (name === 'imageFile') {
      setForm({ ...form, imageFile: files[0] });
    } else {
      setForm({ ...form, [name]: value });
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const headers = { Authorization: `Bearer ${token}` };

    const price = parseFloat(form.price);
    const categoryId = parseInt(form.category_id);

    if (isNaN(price) || isNaN(categoryId) || categoryId <= 0) {
      alert('Giá và Danh mục phải là số hợp lệ và lớn hơn 0');
      return;
    }

    const data = new FormData();
    data.append('name', form.name);
    data.append('description', form.description);
    data.append('price', price);
    data.append('category_id', categoryId);
    if (form.imageFile) {
      data.append('image', form.imageFile);
    }

    try {
      if (editingProduct) {
        await axios.put(`http://localhost:8080/products/${editingProduct.id}`, data, {
          headers: { ...headers, 'Content-Type': 'multipart/form-data' },
        });
      } else {
        await axios.post('http://localhost:8080/products', data, {
          headers: { ...headers, 'Content-Type': 'multipart/form-data' },
        });
      }

      setForm({ name: '', description: '', price: '', imageFile: null, category_id: '' });
      setEditingProduct(null);
      fetchProducts();
    } catch (err) {
      console.error('❌ Lỗi khi lưu sản phẩm:', err);
      alert('Không thể lưu sản phẩm. Kiểm tra lại dữ liệu.');
    }
  };

  const handleEdit = (product) => {
    setEditingProduct(product);
    setForm({
      name: product.name,
      description: product.description,
      price: String(product.price),
      imageFile: null,
      category_id: product.category_id ? String(product.category_id) : '1',
    });
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Bạn có chắc muốn xoá sản phẩm này?')) return;

    try {
      await axios.delete(`http://localhost:8080/products/${id}`, {
        headers: { Authorization: `Bearer ${token}` },
      });
      fetchProducts();
    } catch (err) {
      console.error('❌ Lỗi khi xoá sản phẩm:', err);
      alert('Không thể xoá sản phẩm');
    }
  };

  return (
    <div style={{ maxWidth: '800px', margin: 'auto' }}>
      <h2>📦 Quản lý Sản phẩm</h2>
      <form onSubmit={handleSubmit} style={{ display: 'grid', gap: '0.5rem' }} encType="multipart/form-data">
        <input name="name" placeholder="Tên" value={form.name} onChange={handleChange} required />
        <input name="description" placeholder="Mô tả" value={form.description} onChange={handleChange} required />
        <input name="price" type="number" placeholder="Giá (VNĐ)" value={form.price} onChange={handleChange} required />
        <input name="imageFile" type="file" accept="image/*" onChange={handleChange} />
        <input name="category_id" type="number" placeholder="ID Danh mục" value={form.category_id} onChange={handleChange} required />
        <button type="submit">{editingProduct ? '💾 Cập nhật' : '➕ Thêm mới'}</button>
      </form>

      <ul style={{ marginTop: '2rem', listStyle: 'none', padding: 0 }}>
        {products.map((p) => (
          <li key={p.id} style={{ marginBottom: '2rem', border: '1px solid #ccc', padding: '1rem', borderRadius: '10px' }}>
            <img
              src={p.image_url.startsWith('/images') ? p.image_url : '/images/no-image.png'}
              alt={p.name}
              width="120"
              height="120"
              style={{ objectFit: 'cover', borderRadius: '8px' }}
              onError={(e) => {
                e.target.onerror = null;
                e.target.src = '/images/no-image.png';
              }}
            />
            <div style={{ marginTop: '0.5rem', fontWeight: 'bold' }}>{p.name}</div>
            <div style={{ color: '#e53935', fontSize: '1rem' }}>{Number(p.price).toLocaleString()} VNĐ</div>
            <div style={{ marginTop: '0.5rem' }}>
              <button onClick={() => handleEdit(p)} style={{ marginRight: '0.5rem' }}>✏️ Sửa</button>
              <button onClick={() => handleDelete(p.id)}>🗑️ Xoá</button>
            </div>
          </li>
        ))}
      </ul>
    </div>
  );
}

export default AdminProducts;
