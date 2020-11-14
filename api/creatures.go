package api

import "sort"

type creatures []creature

type creature struct {
  Name   string      `json:"Name"`
  Desc   description `json:"Description"`
  Img    string      `json:"Img"`
  Source string      `json:"Source"`
}

type description []string

func readCreatures(path string) (creatures, error) {
  var cs creatures
  if err := readJSONFile(path, &cs); err != nil {
    return creatures{}, err
  }
  return cs, nil
}

type match struct {
  c     creature
  score float64
}

func (cs creatures) search(query string, n int, s float64) creatures {
  ts := tokeniseQuery(query)

  var m int
  var nameMatches creatures
  var tagMatches []match

CLOOP:
  for _, c := range cs {
    if m >= n {
      break
    }

    nameTok, ok := tokeniseOne(c.Name)
    if ok {
      for _, t := range ts {
        if t.eq(nameTok) {
          nameMatches = append(nameMatches, c)
          m++
          continue CLOOP
        }
      }
    }

    var nt int
    for _, tag := range c.Desc {
      tagTok, ok := tokeniseOne(tag)
      if ok {
        for _, t := range ts {
          if t.eq(tagTok) {
            nt++
          }
        }
      }
    }
    if nt > 0 {
      score := float64(nt) / float64(len(c.Desc))
      if score > s {
        tagMatches = append(tagMatches, match{
          c:     c,
          score: score,
        })
        m++
      }
    }
  }

  if m == 0 {
    return nil
  }

  sort.SliceStable(tagMatches, func(i, j int) bool {
    return tagMatches[i].score > tagMatches[j].score
  })

  var allMatches = make(creatures, m)
  var i int

  for _, c := range nameMatches {
    allMatches[i] = c
    i++
  }
  for _, m := range tagMatches {
    allMatches[i] = m.c
    i++
  }

  return allMatches
}
