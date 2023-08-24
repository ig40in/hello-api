package translation_test

import (
  "testing"
  "github.com/isorokin-info/hello-api/translation"
)

func TestTranslate(t *testing.T) {
  // Arrange
  word := "hello"
  language := "english"

  // Act
  res := translation.Translate(word, language)

  // Assert
  if res != "hello" {
    t.Errorf(`expected "%s" but recieved "%s"`, word, res)
  }
}
