package cmd

func (h *handle) isParkingInitiated() bool {
	return h.parkingmanager != nil
}
