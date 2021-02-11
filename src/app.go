package main

import (
	"flag"
	"fmt"
	"os"

	"./sezar"
)

func main() {

	encrypt := flag.NewFlagSet("encrypt", flag.ExitOnError)
	enText_ptr := encrypt.String("txt", "", "string that you want encrypt or decrypt")
	enShift_ptr := encrypt.Int("shift", 26, "shift cipher")

	decrypt := flag.NewFlagSet("decrypt", flag.ExitOnError)
	deText_ptr := decrypt.String("txt", "", "string that you want encrypt or decrypt")
	deShift_ptr := decrypt.Int("shift", 26, "shift cipher")

	crack := flag.NewFlagSet("crack", flag.ExitOnError)
	crackText_ptr := crack.String("txt", "", "find possible decryption with brute force")

	if len(os.Args) < 2 {
		fmt.Println("expected 'encrypt' or 'decrypt' subcommands ")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "encrypt":
		encrypt.Parse(os.Args[2:])
		svar := sezar.Seasar{Text: *enText_ptr, Shift: int32(*enShift_ptr)}
		fmt.Println(svar.EncodeText())

	case "decrypt":
		decrypt.Parse(os.Args[2:])
		svar := sezar.Seasar{Text: *deText_ptr, Shift: int32(*deShift_ptr)}
		fmt.Println(svar.DecodeText())

	case "crack":
		crack.Parse(os.Args[2:])
		bruteForce(*crackText_ptr)
	default:
		fmt.Println("expected 'encrypt' or 'decrypt' subcommands")
		os.Exit(1)
	}
}

func bruteForce(decrypted string) {
	for i := 1; i < 26; i++ {
		sezar := sezar.Seasar{Text: decrypted, Shift: int32(i)}
		fmt.Println(i, sezar.DecodeText())
	}
}
