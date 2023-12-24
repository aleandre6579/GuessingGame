import Background from '../components/Background';

function Easy() {

  return (
    <>
      <Background/>
      <div className='bg-sky-200/60 border p-12 rounded-lg text-black'>
        <h2 className='text-3xl'>
          Guess the number!
        </h2>
        <p className='mb-8'>
          The number is between 1 and 5
        </p>

        <form className="mx-auto">
          <div className="mb-12">
            <input type="text" id="username-success" className="border border-sky-500 placeholder-sky-700 text-sm text-black rounded-lg block w-full p-2.5 bg-white" placeholder="Guess here..."/>
            <p className="float-left mt-2 text-sm">Must be a number between 1 and 5!</p>
          </div>
          <button className='bg-sky-300 hover:bg-sky-400'>Guess!</button>
        </form>
      </div>

    </>
  )
}

export default Easy
