package core

func (s *parkslot) IsEmpty() bool {
	return s.vhc == nil
}
