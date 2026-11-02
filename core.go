package core

import (
	"errors"
	"sort"
)

type Hit struct {
	ID    string
	Score float64
}

func Fuse(rankings [][]string, k int) ([]Hit, error) {
	if k <= 0 {
		return nil, errors.New("k must be positive")
	}
	scores := map[string]float64{}
	for _, list := range rankings {
		seen := map[string]bool{}
		for rank, id := range list {
			if id == "" || seen[id] {
				continue
			}
			seen[id] = true
			scores[id] += 1 / float64(k+rank+1)
		}
	}
	out := make([]Hit, 0, len(scores))
	for id, score := range scores {
		out = append(out, Hit{id, score})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Score == out[j].Score {
			return out[i].ID < out[j].ID
		}
		return out[i].Score > out[j].Score
	})
	return out, nil
}
