package main

func main() {
	manufacturingDirector := ManufacturingDirector{}

	laptop := &LapTop{}
	manufacturingDirector.SetBuilder(laptop)
	manufacturingDirector.Construct()
	manufacturingDirector.PrintProduct()
}
