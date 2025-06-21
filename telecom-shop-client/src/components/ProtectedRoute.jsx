import React from 'react';
import { Navigate } from 'react-router-dom';
import { useAuth } from '../hooks/useAuth';

export function ProtectedRoute({ children, adminOnly = false }) {
  const { isLoggedIn, isAdmin } = useAuth();

  if (!isLoggedIn) return <Navigate to="/login" />;
  if (adminOnly && !isAdmin) return <Navigate to="/" />;

  return children;
}
