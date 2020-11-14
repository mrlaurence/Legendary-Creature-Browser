package api

import (
  "regexp"
  "strings"
  "unicode"
)

var tokenSplitRegexp = regexp.MustCompile(`[\s,;]+`)

type token struct {
  full, clean string
}

func (t token) eq(t2 token) bool {
  return strings.EqualFold(t.clean, t2.clean)
}

func tokeniseQuery(q string) []token {
  return tokeniseSlice(tokenSplitRegexp.Split(q, -1))
}

func tokeniseSlice(sl []string) []token {
  var tokens []token
  for _, s := range sl {
    if tok, ok := tokeniseOne(s); ok {
      tokens = append(tokens, tok)
    }
  }
  return tokens
}

func tokeniseOne(s string) (token, bool) {
  var letters strings.Builder
  for _, r := range s {
    if unicode.IsLetter(r) {
      letters.WriteRune(r)
    }
  }
  if letters.Len() > 0 {
    return token{
      full:  s,
      clean: strings.ToLower(letters.String()),
    }, true
  }
  return token{}, false
}
