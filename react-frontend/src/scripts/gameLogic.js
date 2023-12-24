import axios from "axios";

export async function makeGuess(guess) {
  const res = await axios.post("/api/guess", guess)
  if(res.status === 201) {
    alert("Correct!")
    return true
  } else {
    alert("Wrong!")
    return false
  }
}

