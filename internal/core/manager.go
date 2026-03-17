package core

import "sync"

type BuildManager struct {
	mu     sync.RWMutex
	builds map[string]*Build
}

func NewBuildManager() *BuildManager {
	return &BuildManager{builds: make(map[string]*Build)}
}

func (bm *BuildManager) Add(b *Build) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	bm.builds[b.ID] = b
}

func (bm *BuildManager) Get(id string) (*Build, bool) {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	b, ok := bm.builds[id]
	return b, ok
}

func (bm *BuildManager) All() []*Build {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	out := make([]*Build, 0, len(bm.builds))
	for _, b := range bm.builds {
		out = append(out, b)
	}
	return out
}

func (bm *BuildManager) Remove(id string) {
	bm.mu.Lock()
	defer bm.mu.Unlock()
	delete(bm.builds, id)
}

func (bm *BuildManager) Count() int {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	return len(bm.builds)
}

func (bm *BuildManager) Active() []*Build {
	bm.mu.RLock()
	defer bm.mu.RUnlock()
	var out []*Build
	for _, b := range bm.builds {
		if b.State == StateBuilding || b.State == StateQueued {
			out = append(out, b)
		}
	}
	return out
}
