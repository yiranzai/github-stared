package repoimpl

import (
	"log"

	"github.com/yiranzai/github-starred/domain/model"
	"github.com/yiranzai/github-starred/domain/repository"
)

// MemorySumHistory represents on memory repository of sum history
type MemorySumHistory struct {
	history []*model.SumHistory
}

// NewMemorySumHistory create new MemorySumHistory repository
func NewMemorySumHistory(v []*model.SumHistory) repository.SumHistory {
	if v == nil {
		v = []*model.SumHistory{}
	}
	return &MemorySumHistory{history: v}
}

// Add add new sum history to repository
func (m *MemorySumHistory) Add(sumHistory *model.SumHistory) {
	m.history = append(m.history, sumHistory)
	log.Printf("current history: %v", m.history)
}

// List return list of sum history
func (m *MemorySumHistory) List(limit int) []*model.SumHistory {
	l := len(m.history)
	if l <= limit {
		return m.history
	}
	return m.history[l-limit : l]
}
