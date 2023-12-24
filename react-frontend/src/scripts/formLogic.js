

export function isGuessValid(guess, lower, upper) {
  try {
    guess = parseInt(guess)
  } catch(e) {
    return false
  }
  if (!Number.isInteger(guess)) {
    return false
  }
  if (guess > upper || guess < lower) {
    return false
  }

  return true
}
