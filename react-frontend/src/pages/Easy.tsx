import { makeGuess } from '../scripts/gameLogic.js'
import { isGuessValid } from '../scripts/formLogic.js'
import {useState} from 'react'

function Easy() {
  const [inputGuess, setGuess] = useState(0)
  const [invalidGuess, setInvalidGuess] = useState(false)

  const handleInputChange = event => {
    setGuess(event.target.value)
    if (!isGuessValid(event.target.value, 1, 5)) {
      setInvalidGuess(true)
    } else {
      setInvalidGuess(false)
    }
  }

  return (
    <>
      <div className='bg-sky-200/60 border p-12 rounded-lg text-black'>
        <h2 className='text-3xl'>
          Guess the number!
        </h2>
        <p className='mb-8'>
          The number is between 1 and 5
        </p>

        <form className="mx-auto">
          <div className="mb-16">
            <input 
              onChange={handleInputChange} 
              type="text" 
              id="username-success" 
              className="border border-sky-500 placeholder-sky-700 text-sm text-black rounded-lg block w-full p-2.5 bg-white" 
              placeholder="Enter a guess"/>
            <p className="float-left mt-2 text-sm">Input must be a number between 1 and 5</p>
          </div>
          <button disabled={invalidGuess} onClick={() => makeGuess(inputGuess)} className='disabled:bg-blue-500 enabled:bg-sky-300 enabled:hover:bg-sky-400'>Guess!</button>
        </form>
      </div>

    </>
  )
}

export default Easy
