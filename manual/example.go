package manual

func examples() string {
	return `	# Here are a few example code snippets

	x := 6 * 7 # The answer to ultimate question
	@Println("6 * 7 = ", x)

	# Print the parity of all numbers up to and including 'max'
	i, max := 1, 12
	loop [i <= max] {
  
  	[i % 2 == 0] {
    	@Println(i, " is even")
  	}
  
  	[i % 2 == 1] {
    	@Println(i, " is odd")
  	}

  	i := i + 1
	}

	# 'e' represents an error string, if it's not empty then there is an error
	e := "Panic!"; [e != ""] {
		@Println("ERROR: ", e)
		@Exit(1)
	}

	@Exit(0)`
}
