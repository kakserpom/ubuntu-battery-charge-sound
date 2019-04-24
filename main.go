package main

import (
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if (e != nil) {
		panic(e);
	}
}

func playSound() {
	f, err := os.Open("carrier-has-arrived.mp3");
	check(err);
	streamer, format, err := mp3.Decode(f);
	check(err);
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer);
	//defer streamer.Close();
}
func main() {
	var soundPlayed bool;

	for true {
		data, err := ioutil.ReadFile("/sys/class/power_supply/BAT1/capacity");
		check(err);
		percentage, err := strconv.Atoi(strings.Trim(string(data), "\n"));
		check(err);
		data, err = ioutil.ReadFile("/sys/class/power_supply/BAT1/status");
		check(err);
		status := strings.Trim(string(data), "\n");
		if status == "Charging" {
			if (!soundPlayed && percentage > 95) {
				playSound();
				soundPlayed = true;
			}
		} else {
			soundPlayed = false;
		}
		//fmt.Println(percentage);
		//fmt.Println(status);
		time.Sleep(time.Second);
	}
}
