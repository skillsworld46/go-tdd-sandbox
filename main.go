package main

import (
	"fmt"
	"os"
	"time"

	"github.com/skillsworld46/go-tdd-sandbox/mocking"
	"github.com/skillsworld46/go-tdd-sandbox/services"
)

func main() {
	fmt.Println(services.Hello("Jessica",""))

	// mocking test with DI, SleepFn, sleeper, writer dependency are passed in later
	sleeper := &mocking.ConfigurableSleeper{Duration: 1*time.Second, SleepFn: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
