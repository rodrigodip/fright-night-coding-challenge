package main

import (
	"fmt"
	"strconv"
)

func convertPhoneNumber(phoneNumber string) string {
	// Parse the decimal phone number
	num, err := strconv.ParseUint(phoneNumber, 10, 64)
	if err != nil {
		return "0,0" // Handle parsing error
	}

	max32Bit := uint64(4294967295)      // 2^32 - 1
	overflowValue := uint64(4294967296) // 2^32

	var overflows uint64 = 0

	// Count overflows and reduce the number
	for num > max32Bit {
		overflows++
		num -= overflowValue
	}

	// Convert to binary and remove leading zeros
	binaryStr := strconv.FormatUint(num, 2)

	return fmt.Sprintf("%d,%s", overflows, binaryStr)
}

func main() {
	result := convertPhoneNumber("7660716255")
	fmt.Printf(result)
}
