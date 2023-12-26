import { makeGuess } from '../scripts/gameLogic.js'
import { isGuessValid } from '../scripts/formLogic.js'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../auth/authProvider.js'

function Level({ levelName, lower, upper}) {
  const navigate = useNavigate()
  const auth = useAuth();

  const [inputGuess, setGuess] = useState(0)
  const [invalidGuess, setInvalidGuess] = useState(true)
  const [isCorrect, setIsCorrect] = useState(-1)
  const [guessedNumber, setGuessedNumber] = useState(0)

  const handleInputChange = event => {
    const guess = parseInt(event.target.value)
    setGuess(guess)
    if (!isGuessValid(event.target.value, lower, upper)) {
      setInvalidGuess(true)
    } else {
      setInvalidGuess(false)
    }
  }

  const handleGuess = async (guess) => {
    const res = await makeGuess(guess, levelName)
    if (res === 401) {
      auth?.setToken("")
      return
    }
    setIsCorrect(res)
    setGuessedNumber(guess)
  }

  return (
    <>
      <div className='relative flex flex-col items-center'>
        <div className='border-4 text-sky-500 border-sky-500 mb-20 relative mainContainer bg-sky-200/90 h-full p-10 rounded-full'>
          <h2 className='cursor-default text-6xl rounded-full font-bold'>
            Guess The Number!
          </h2>
        </div>

        <div className='relative px-28 py-[69px] flex flex-col items-center gap-10 bg-sky-200/90 text-2xl rounded-lg text-black'>
          <span onClick={() => {navigate('/')}} className="hover:bg-black/0 hover:translate-x-[-0.25em] transition duration-300 cursor-pointer absolute top-5 left-5 text-lg text-black bg-black/20 rounded-lg p-2">
              &#x25c0; Back
          </span>
          <p className=''>
            The number is between {lower} and {upper}<br/>
            <span className='text-lg'>({lower} and {upper} included)</span>
          </p>

          <input 
            onChange={handleInputChange} 
            autoComplete='off'
            type="text" 
            id="username-success" 
            className="border-2 border-sky-500 placeholder-sky-700 text-sm text-black rounded-lg block w-full p-2.5 bg-white" 
            placeholder="Enter a guess"/>

          <button
            disabled={invalidGuess} 
            onClick={() => handleGuess(inputGuess)} 
            className='py-3 px-8 rounded-lg disabled:bg-blue-500 enabled:bg-sky-400 enabled:hover:bg-sky-500'
          >
            Guess!
          </button>
        </div>
        <div className={((isCorrect&&isCorrect!==-1)?'block':'hidden') + ' absolute bottom-[-4em] px-10 py-5 mt-10 border-4 border-green-500 bg-sky-200/80 text-3xl rounded-lg text-green-500'}>
          <span className='relative top-[-10px]'>{guessedNumber} is Correct!<br/></span>
          <span className='absolute bottom-2 left-[7px] text-sm'>(A new number has been generated)</span>
        </div>
        <div className={((!isCorrect&&isCorrect!==-1)?'block':'hidden') + ' absolute bottom-[-4em] px-10 py-5 mt-10 border-4 border-red-500 bg-sky-200/80 text-3xl rounded-lg text-red-500'}>
          {guessedNumber} is Wrong!
        </div>
      </div>

    </>
  )
}

export default Level
