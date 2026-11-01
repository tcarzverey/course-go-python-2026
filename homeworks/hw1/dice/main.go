package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	os.Exit(run())
}

func run() int {
	fs := flag.NewFlagSet("dice", flag.ContinueOnError)
	var (
		count   = 1
		sum     = false
		verbose = false
		seed    = int64(0)
		help    = false
	)

	fs.IntVar(&count, "n", count, "количество итераций бросков")
	fs.IntVar(&count, "count", count, "количество итераций бросков (alias -n)")
	fs.BoolVar(&sum, "s", false, "выводить только сумму костей")
	fs.BoolVar(&sum, "sum", false, "выводить только сумму костей (alias -s)")
	fs.BoolVar(&verbose, "v", false, "подробный вывод по каждому броску")
	fs.BoolVar(&verbose, "verbose", false, "подробный вывод по каждому броску (alias -v)")
	fs.Int64Var(&seed, "seed", 0, "зерно генератора случайных чисел")
	fs.BoolVar(&help, "h", false, "показать справку")
	fs.BoolVar(&help, "help", false, "показать справку (alias -h)")
	fs.SetOutput(os.Stdout)
	fs.Usage = func() {
		fmt.Fprintln(os.Stdout, "dice — инструмент для броска игральных костей в формате NdM.")
		fmt.Fprintln(os.Stdout, "Использование: dice [flags] NdM")
		fmt.Fprintln(os.Stdout, "где NdM — N костей (1..100) по M граней (1..100).")
		fmt.Fprintln(os.Stdout, "Флаги:")
		fs.PrintDefaults()
	}

	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)

		return 1
	}

	if help {
		fs.Usage()

		return 0
	}

	if fs.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "expected exactly one dice argument in the format NdM, got %d\n", fs.NArg())
		fs.Usage()

		return 1
	}

	spec, err := parseSpec(fs.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)

		return 1
	}

	if sum && verbose {
		fmt.Fprintln(os.Stderr, "flags -s/--sum and -v/--verbose are mutually exclusive")

		return 1
	}
	if count < 1 {
		fmt.Fprintln(os.Stderr, "number of iterations (-n/--count) must be positive")

		return 1
	}

	if !flagWasSet(fs, "seed") {
		seed = time.Now().UnixNano()
	}
	r := rand.New(rand.NewSource(seed))

	if verbose {
		if count > 1 {
			fmt.Printf("Rolling %dd%d (%d iterations):\n", spec.count, spec.sides, count)
		} else {
			fmt.Printf("Rolling %dd%d:\n", spec.count, spec.sides)
		}
	}

	for iter := 1; iter <= count; iter++ {
		values := roll(r, spec)
		pad := ""
		if count > 1 {
			fmt.Printf("Iteration %d:\n", iter)
			pad = "  "
		}
		switch {
		case verbose:
			printDetailed(pad, values)
		case sum:
			fmt.Printf("%sSum: %d\n", pad, sumValues(values))
		default:
			fmt.Printf("%s%s\n", pad, formatRolls(values))
		}
	}

	return 0
}

// flagWasSet сообщает, был ли флаг задан в командной строке.
func flagWasSet(fs *flag.FlagSet, name string) bool {
	set := false
	fs.Visit(func(f *flag.Flag) {
		if f.Name == name {
			set = true
		}
	})

	return set
}
