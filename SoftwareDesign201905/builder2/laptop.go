package main

type LapTop struct {
	electronicProduct ElectronicProduct
}

func (l *LapTop) SetStructure() BuildProcess {
	l.electronicProduct.Structure = "LapTop"
	return l
}

func (l *LapTop) SetMonitor() BuildProcess {
	l.electronicProduct.Monitor = 1
	return l
}

func (l *LapTop) SetCamera() BuildProcess {
	l.electronicProduct.Camera = 1
	return l
}

func (l *LapTop) GetGadget() ElectronicProduct {
	return l.electronicProduct
}
