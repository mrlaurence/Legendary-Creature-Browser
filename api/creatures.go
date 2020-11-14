package api

import "sort"

type creaturesModel map[string]creatureModel

type creatureModel struct {
  Desc []string `json:"Description"`
  Img  string   `json:"Img"`
  Link string   `json:"Link"`
}

type creatures []creature

type creature struct {
  name      token
  desc      []token
  img, link string
}

func (csm creaturesModel) parse() creatures {
  cs := make(creatures, len(csm))
  var i int
  for name, cm := range csm {
    cs[i] = creature{
      desc: tokeniseSlice(cm.Desc),
      img:  cm.Img,
      link: cm.Link,
    }
    if nameTok, ok := tokeniseOne(name); ok {
      cs[i].name = nameTok
    }
    i++
  }
  return cs
}

func (cs creatures) toModel() creaturesModel {
  csm := make(creaturesModel)
  for _, c := range cs {
    desc := make([]string, len(c.desc))
    for j, tagTok := range c.desc {
      desc[j] = tagTok.full
    }
    csm[c.name.full] = creatureModel{
      Desc: desc,
      Img:  c.img,
      Link: c.link,
    }
  }
  return csm
}

func readCreatures(path string) (creatures, error) {
  var cs creaturesModel
  if err := readJSONFile(path, &cs); err != nil {
    return nil, err
  }
  return cs.parse(), nil
}

type match struct {
  c     creature
  score float64
}

func (cs creatures) search(query string, n int, s float64) creatures {
  ts := tokeniseQuery(query)

  if len(ts) == 0 {
    return nil
  }

  var m int
  var nameMatches creatures
  var tagMatches []match

CLOOP:
  for _, c := range cs {
    if m >= n {
      break
    }

    for _, t := range ts {
      if t.eq(c.name) {
        nameMatches = append(nameMatches, c)
        m++
        continue CLOOP
      }
    }

    var nt int
    for _, tag := range c.desc {
      for _, t := range ts {
        if t.eq(tag) {
          nt++
        }
      }
    }
    if nt > 0 {
      score := float64(nt) / float64(len(c.desc))
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
