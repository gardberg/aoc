package main

import (
    "strings"
	"strconv"
)

func areNumbersAscending(s string) bool {
    
    tokens := strings.Fields(s) // split string on whitespace
    
    last := -1
    
    for _, c := range tokens {
        if i, err := strconv.Atoi(c); err == nil {
            if last >= i {
                return false;
            } else {
                last = i
            }
        }
    }
    
    return true;
}