package main

func main() {
	var _ = ReadFileIntoString("resources/dec_03/test_ex_02.txt")
	var testInputString = ReadFileIntoString("resources/dec_03/input_1.txt")

	//solveExercise1Day03(testInputString)
	solveExercise2Day03(testInputString)
}
