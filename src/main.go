package main

func main() {
	memory, startPCH, startPCL := loadFile("1_first.prg")
	cpu := newCPU(memory)
	cpu.Start(startPCH, startPCL)

	/*cpu.setStatusCarry(true)
	fmt.Printf("%08b", cpu.S)
	fmt.Println(cpu.getStatusCarry())

	cpu.setStatusCarry(false)
	fmt.Printf("%08b", cpu.S)
	fmt.Println(cpu.getStatusCarry())*/

}
