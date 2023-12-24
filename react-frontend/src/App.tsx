import AuthProvider from "./auth/authProvider";
import Routes from "./routes";
import Background from "./components/Background";

function App() {
  return (
    <>
      <Background/>
      <AuthProvider>
        <Routes />
      </AuthProvider>
    </>
  );
}

export default App;
