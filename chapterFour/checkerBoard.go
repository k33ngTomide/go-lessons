package main

func main() {

	for counter := 0; counter < 8; counter++ {
		for count := 0; count < 8; count++ {
			if counter%2 == 0 {
				print("*  ")

			}
			if counter%2 != 0 {
				print("  *")
			}
		}
		println()
	}

}
