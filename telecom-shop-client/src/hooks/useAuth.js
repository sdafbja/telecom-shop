export function useAuth() {
  const token = localStorage.getItem('token');
  const user = JSON.parse(localStorage.getItem('user'));

  return {
    isLoggedIn: !!token,
    user,
    isAdmin: user?.role === 'admin',
  };
}
