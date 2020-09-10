package main

import (
	"flag"
	"fmt"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func playSound(path string) {
	f, err := os.Open(path)
	check(err)
	streamer, format, err := mp3.Decode(f)
	check(err)
	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	check(err)
	speaker.Play(streamer)
	//defer streamer.Close()
}
func main() {
	var soundPlayed bool

	targetPercentage := flag.Int("p", 95, "target percentage")
	verbosity := flag.Int("v", 0, "verbosity level")
	soundFile := flag.String("f", "carrier-has-arrived.mp3", "mp3 file")
	flag.Parse()

	for {
		data, err := ioutil.ReadFile("/sys/class/power_supply/BAT1/capacity")
		check(err)
		percentage, err := strconv.Atoi(strings.Trim(string(data), "\n"))
		check(err)
		data, err = ioutil.ReadFile("/sys/class/power_supply/BAT1/status")
		check(err)
		status := strings.Trim(string(data), "\n")

		if *verbosity > 1 {
			fmt.Println(percentage)
			fmt.Println(status)
		}

		if status == "Charging" {
			if !soundPlayed && percentage >= *targetPercentage {
				if *verbosity > 0 {
					fmt.Println("Playing the sound.")
				}
				playSound(*soundFile)
				soundPlayed = true
			}
		} else {
			soundPlayed = false
		}
		time.Sleep(time.Second)
	}
}
