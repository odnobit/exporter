package storage

import "sync/atomic"

var Memory = &MemoryMetricsStorage{}

type MemoryMetricsStorage struct {
	totalMessages   atomic.Uint64
	failedMessages  atomic.Uint64
	successMessages atomic.Uint64
}

func (m *MemoryMetricsStorage) AddTotalMessages(value uint64) {
	m.totalMessages.Add(value)
}

func (m *MemoryMetricsStorage) AddFailedMessages(value uint64) {
	m.failedMessages.Add(value)
}

func (m *MemoryMetricsStorage) AddSuccessMessages(value uint64) {
	m.successMessages.Add(value)
}

func (m *MemoryMetricsStorage) IncTotalMessages() {
	m.totalMessages.Add(1)
}

func (m *MemoryMetricsStorage) IncFailedMessages() {
	m.failedMessages.Add(1)
}

func (m *MemoryMetricsStorage) IncSuccessMessages() {
	m.successMessages.Add(1)
}

func (m *MemoryMetricsStorage) GetTotalMessages() float64 {
	return float64(m.totalMessages.Load())
}

func (m *MemoryMetricsStorage) GetFailedMessages() float64 {
	return float64(m.failedMessages.Load())
}

func (m *MemoryMetricsStorage) GetSuccessMessages() float64 {
	return float64(m.successMessages.Load())
}
