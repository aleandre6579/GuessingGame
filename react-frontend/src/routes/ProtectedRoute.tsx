import { Navigate, Outlet, useLocation } from "react-router-dom";
import { useAuth } from "../auth/authProvider";

export const ProtectedRoute = () => {
  const auth = useAuth();
  const location = useLocation();

  if (!auth?.token) {
    return <Navigate to="/login" state={{path: location.pathname}}/>;
  }

  return <Outlet />;
};
