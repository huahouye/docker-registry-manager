package manager

import "time"

type Repository struct {
	Name string
	Tags map[string]*Tag
}

func (r *Repository) LastModified() time.Time {
	var lastModified time.Time
	for _, tag := range r.Tags {
		for _, history := range tag.History {
			if history.Created.After(lastModified) {
				lastModified = history.Created
			}
		}
	}
	return lastModified
}

func (r *Repository) Size() (size int64) {
	dedup := make(map[string]struct{})
	for _, tag := range r.Tags {
		for _, layer := range tag.Layers {
			if _, ok := dedup[layer.Digest.String()]; !ok {
				dedup[layer.Digest.String()] = struct{}{}
				size += layer.Size
			}
		}
	}
	return size
}
