import DifficultyButton from '../components/DifficultyButton';
import { Link } from "react-router-dom";
import '../App.css'

function App() {

  return (
    <div className='flex flex-col items-center'>
      <div className='border-4 text-sky-500 border-sky-500 mb-20 relative mainContainer bg-sky-200/90 h-full p-10 rounded-full'>
        <h2 className='cursor-default text-6xl rounded-full font-bold'>
          Guess The Number!
        </h2>
      </div>

      <div className='relative w-100 bg-sky-200/90 h-full p-14 rounded-lg'>
        <div className='flex flex-col gap-10'>
          <Link to='/easy'><DifficultyButton text='Easy' color='cornflowerblue'></DifficultyButton></Link>
          <Link to='/medium'><DifficultyButton text='Medium' color='darksalmon'></DifficultyButton></Link>
          <Link to='/hard'><DifficultyButton text='Hard' color='indianred'></DifficultyButton></Link>
        </div>
      </div>
    </div>
  )
}

export default App
