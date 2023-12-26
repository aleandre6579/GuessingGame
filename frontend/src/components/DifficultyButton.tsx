
function DifficultyButton({text, color}) {

  return (
    <button className={`w-80 h-16 text-xl rounded-lg hover:scale-110 transition ease-in-out duration-300`} style={{backgroundColor: color}}>
      {text}
    </button>
  )
}

export default DifficultyButton
