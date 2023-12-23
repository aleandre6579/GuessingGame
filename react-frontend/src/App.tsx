import './App.css'
import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from './pages/Home';
import Easy from './pages/Easy';
import Medium from './pages/Medium';
import Hard from './pages/Hard';

function App() {

  return (
    <>
      <BrowserRouter>
        <Routes>
            <Route index element={<Home />} />
            <Route path="easy" element={<Easy />} />
            <Route path="medium" element={<Medium />} />
            <Route path="hard" element={<Hard />} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
