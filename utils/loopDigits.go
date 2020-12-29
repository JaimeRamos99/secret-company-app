package utils

import "strconv"

//creates a map that contains the not splitable runes that are not letter
//(the split rune is consired as a unicode digit...)
func LoopDigits() map[string]int{
  //adding runes that are not digits
  allowed_runes := map[string]int{
    "#": 1,
    ".": 1,
    ",": 1,
    "(": 1,
    ")": 1,
  }

  //adding digits
  for i := 0; i < 10; i++ {
    allowed_runes[strconv.Itoa(i)] = 1
  }
  return allowed_runes
}
