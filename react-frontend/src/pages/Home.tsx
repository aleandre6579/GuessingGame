import Background from '../components/Background';
import DifficultyButton from '../components/DifficultyButton';
import { Link } from "react-router-dom";
import '../App.css'

function App() {

  return (
    <>
      <Background/>
      <div className='flex flex-col gap-10'>
        <Link to='/easy'><DifficultyButton text='Easy' color='cornflowerblue'></DifficultyButton></Link>
        <Link to='/medium'><DifficultyButton text='Medium' color='darksalmon'></DifficultyButton></Link>
        <Link to='/hard'><DifficultyButton text='Hard' color='indianred'></DifficultyButton></Link>
      </div>

    </>
  )
}

export default App
