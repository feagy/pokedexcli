package main

import (
    "testing"
)


func TestCleanInput(t *testing.T) {
    cases := []struct {
	input    string
	expected []string
    }{
        {
	        input:    "  hello  world  ",
	        expected: []string{"hello", "world"},
        },
        {
	        input:    "Charmander Bulbasaur PIKACHU",
	        expected: []string{"charmander", "bulbasaur", "pikachu"},
        },
    }

    for j, c := range cases {
	    actual := cleanInput(c.input)
	    if len(actual) != len(c.expected) {
		    t.Errorf("Clean Input Test Case %d: Lenghts of slices doesn't match %v != %v", j, len(actual), len(c.expected))
            continue
	    }
	    for i := range actual {
		    word := actual[i]
		    expectedWord := c.expected[i]
            if word != expectedWord {
    		    t.Errorf("Clean Input Test Case %d: index %d: expected \"%v\", got \"%v\"", j, i, expectedWord, word)
                break
            }
	    }
    }    
}
