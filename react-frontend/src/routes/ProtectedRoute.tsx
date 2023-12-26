import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "../auth/authProvider";


export const ProtectedRoute = () => {
  const auth = useAuth();

  if (!auth?.token) {
    return <Navigate to="/login" />;
  }

  return <Outlet />;
};
