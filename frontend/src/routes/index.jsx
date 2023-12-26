import { RouterProvider, createBrowserRouter } from "react-router-dom";
import { useAuth } from "../auth/authProvider";
import { ProtectedRoute } from "./ProtectedRoute";
import Home from "../pages/Home"
import Level from "../pages/Level"
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
          element: <Level levelName="easy" lower={1} upper={5}/>,
        },
        {
          path: "/medium",
          element: <Level levelName="medium" lower={1} upper={10}/>,
        },
        {
          path: "/hard",
          element: <Level levelName="hard" lower={1} upper={20}/>,
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
