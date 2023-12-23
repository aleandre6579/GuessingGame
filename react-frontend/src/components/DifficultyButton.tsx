
function DifficultyButton({text, color}) {

  return (
    <button className={`w-64 h-14 hover:scale-110 transition ease-in-out duration-300`} style={{backgroundColor: color}}>
      {text}
    </button>
  )
}

export default DifficultyButton
