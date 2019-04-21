package main

type BuildProcess interface {
	SetStructure() BuildProcess
	SetCamera() BuildProcess
	SetMonitor() BuildProcess
	GetGadget() ElectronicProduct
}
