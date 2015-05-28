package asciicast

import (
	"time"
	"os/exec"
	"github.com/asciinema/asciinema/terminal"
)

type Player interface {
	Play(string, uint) error
}

type AsciicastPlayer struct {
	Terminal terminal.Terminal
}

func NewPlayer() Player {
	return &AsciicastPlayer{Terminal: terminal.NewTerminal()}
}

func Say(text string) {
     cmd := exec.Command("say", text)
     _ = cmd.Run()
}

func (r *AsciicastPlayer) Play(path string, maxWait uint) error {
	asciicast, err := Load(path)
	if err != nil {
		return err
	}

	for _, frame := range asciicast.Stdout {
		delay := frame.Delay
		if delay < 0 {
		        go Say(string(frame.Data));
			continue
		}
		
		if maxWait > 0 && delay > float64(maxWait) {
			delay = float64(maxWait)
		}
		time.Sleep(time.Duration(float64(time.Second) * delay))
		r.Terminal.Write(frame.Data)
	}

	return nil
}
