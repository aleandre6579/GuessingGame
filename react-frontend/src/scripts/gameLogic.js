import axios from "axios";

export async function makeGuess(guess, level) {
  const res = await axios.post("/api/guess", JSON.stringify({number: guess, level: level}))
  if(res.status === 201) {
    alert("Correct!")
    return true
  } else {
    alert("Wrong!")
    return false
  }
}

