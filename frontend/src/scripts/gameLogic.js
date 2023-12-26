import axios from "axios";

export async function makeGuess(guess, level) {
  let res = null
  try {
    res = await axios.post("/api/guess", JSON.stringify({number: guess, level: level}))
  } catch(e) {
    if (e.response.status === 401) {
      return 401
    } else {
      return -1
    }
  }

  if(res.status === 201) {
    return 1
  } else {
    return 0
  }
}

