package translation

import "strings"

func sanitizeInput(w string) string {
  w = strings.TrimSpace(w)
  return strings.ToLower(w)
}

func Translate(word string, language string) string {
  word = sanitizeInput(word)
  language = sanitizeInput(language)

  // handling unknown words
  if word != "hello" {
    return ""
  }

  // handling language
  switch language {
  case "english":
    return "hello"
  case "finnish":
    return "hei"
  case "german":
    return "hallo"
  default:
    return ""
  }
}
