import AuthProvider from "./auth/authProvider";
import Routes from "./routes";

function App() {
  return (
    <AuthProvider>
      <Routes />
    </AuthProvider>
  );
}

export default App;
