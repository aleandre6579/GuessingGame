import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { useAuth } from "../auth/authProvider";
import { ProtectedRoute } from "./ProtectedRoute";
import Home from "../pages/Home"
import Easy from "../pages/Easy"
import Medium from "../pages/Medium"
import Hard from "../pages/Hard"
import Login from "../pages/Login"
import Register from "../pages/Register"

const Routes = () => {
  const { token } = useAuth();

  const routesForPublic = [
    {
      path: "/",
      element: <Home/>,
    },
  ];

  const routesForAuthenticatedOnly = [
    {
      path: "/",
      element: <ProtectedRoute />,
      children: [
        {
          path: "/easy",
          element: <Easy/>,
        },
        {
          path: "/medium",
          element: <Medium/>,
        },
        {
          path: "/hard",
          element: <Hard/>,
        },
      ],
    },
  ];

  const routesForNotAuthenticatedOnly = [
    {
      path: "/login",
      element: <Login/>,
    },
    {
      path: "/register",
      element: <Register/>,
    },
  ];

  const router = createBrowserRouter([
    ...routesForPublic,
    ...(!token ? routesForNotAuthenticatedOnly : []),
    ...routesForAuthenticatedOnly,
  ]);
  
  return <RouterProvider router={router} />;
};

export default Routes;
