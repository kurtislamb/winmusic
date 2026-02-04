package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/urfave/cli/v3"
)

func sendMediaKey(ctx context.Context, cmd *cli.Command, mediaCmd int) error {
	ps := exec.Command("powershell.exe", fmt.Sprintf("(New-Object -ComObject WScript.Shell).SendKeys([char]%d)", mediaCmd))
	ps.Stdout = os.Stdout
	ps.Stderr = os.Stderr
	return ps.Run()
}

func sendVolKey(ctx context.Context, cmd *cli.Command, mediaCmd int) error {
	increments := cmd.Int("increments")
	for i := 0; i < increments; i++ {
		ps := exec.Command("powershell.exe", fmt.Sprintf("(New-Object -ComObject WScript.Shell).SendKeys([char]%d)", mediaCmd))
		ps.Stdout = os.Stdout
		ps.Stderr = os.Stderr
		if err := ps.Run(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	cmd := &cli.Command{
		Commands: []*cli.Command{
			{
				Name:    "play",
				Usage:   "Play the current loaded track",
				Action:  CommandPlay,
				Aliases: []string{"p", "start", "resume"},
			},
			{
				Name:    "stop",
				Usage:   "Stop the current loaded track",
				Action:  CommandStop,
				Aliases: []string{"s", "pause", "halt"},
			},
			{
				Name:    "next",
				Usage:   "Next track",
				Action:  CommandNext,
				Aliases: []string{"n", "forward"},
			},
			{
				Name:    "back",
				Usage:   "Previous track",
				Action:  CommandBack,
				Aliases: []string{"b", "rewind", "previous"},
			},
			{
				Name:    "volumeup",
				Usage:   "Increase system volume by 2: winmusic up 10 (increases volume by 20)",
				Action:  CommandVolumeUp,
				Aliases: []string{"volup", "up"},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "increments",
						Aliases: []string{"p"},
						Usage:   "number of increments to increase volume by",
						Value:   10,
					},
				},
			},
			{
				Name:    "volumedown",
				Usage:   "Decrease system volume by 2: winmusic down 10 (decreases volume by 20)",
				Action:  CommandVolumeDown,
				Aliases: []string{"voldown", "down"},
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "increments",
						Aliases: []string{"p"},
						Usage:   "number of increments to decrease volume by",
						Value:   10,
					},
				},
			},
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}

func CommandPlay(ctx context.Context, cmd *cli.Command) error {
	return sendMediaKey(ctx, cmd, 179) // Play Command
}

func CommandStop(ctx context.Context, cmd *cli.Command) error {
	return sendMediaKey(ctx, cmd, 178) // Stop Command
}

func CommandNext(ctx context.Context, cmd *cli.Command) error {
	return sendMediaKey(ctx, cmd, 176) // Next Command
}

func CommandBack(ctx context.Context, cmd *cli.Command) error {
	return sendMediaKey(ctx, cmd, 177) // Back Command
}

func CommandVolumeUp(ctx context.Context, cmd *cli.Command) error {
	return sendVolKey(ctx, cmd, 175) // Volume Up
}

func CommandVolumeDown(ctx context.Context, cmd *cli.Command) error {
	return sendVolKey(ctx, cmd, 174) // Volume Down
}
