import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import Background from './components/Background';

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <Background/>
      
      <div>
        <button className='w-5 h-5'>
          Easy
        </button>
      </div>
    </>
  )
}

export default App
