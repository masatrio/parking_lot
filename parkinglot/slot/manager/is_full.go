package manager

func (m *manager) IsFull() bool {
	return m.filled == m.size
}
