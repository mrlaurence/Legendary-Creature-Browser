package api

import (
  "regexp"
  "strings"
  "unicode"
)

var tokenSplitRegexp = regexp.MustCompile(`[\s,;]+`)

type token struct {
  full, clean string
  subtokens []token
}

func (t token) eq(t2 token) bool {
  if strings.EqualFold(t.clean, t2.clean) {
    return true
  }
  for _, st := range t.subtokens {
    if st.eq(t2) {
      return true
    }
  }
  for _, st := range t2.subtokens {
    if t.eq(st) {
      return true
    }
  }
  return false
}

func tokeniseQuery(q string) []token {
  return tokeniseSlice(tokenSplitRegexp.Split(q, -1), false)
}

func tokeniseSlice(sl []string, subtokens bool) []token {
  var tokens []token
  for _, s := range sl {
    if tok, ok := tokeniseOne(s, subtokens); ok {
      tokens = append(tokens, tok)
    }
  }
  return tokens
}

func tokeniseOne(s string, subtokens bool) (token, bool) {
  var letters strings.Builder
  for _, r := range s {
    if unicode.IsLetter(r) {
      letters.WriteRune(r)
    }
  }
  if letters.Len() > 0 {
    tok := token{
      full:  s,
      clean: strings.ToLower(letters.String()),
    }
    if subtokens {
      split := tokenSplitRegexp.Split(s, -1)
      if len(split) > 1 {
        tok.subtokens = tokeniseSlice(split, false)
      }
    }
    return tok, true
  }
  return token{}, false
}
