package main

import (
	"fmt"
	"os"
	"os/exec"
)

type runnable interface {
	run() error
	handleErr(error) error
}

type step struct {
	name     string
	startMsg string
	cmd      *exec.Cmd
}

type outputStep struct {
	step
}

func newStep(name, startMsg string, cmd *exec.Cmd) step {
	return step{
		name:     name,
		startMsg: startMsg,
		cmd:      cmd,
	}
}

func newOutputStep(name, startMsg string, cmd *exec.Cmd) outputStep {
	return outputStep{newStep(name, startMsg, cmd)}
}

func (s step) run() error {
	fmt.Print(s.startMsg)
	if err := s.cmd.Run(); err != nil {
		return err
	}

	fmt.Println("\tDone!")
	return nil
}
func (o outputStep) run() error {
	fmt.Println(o.startMsg)
	out, err := o.cmd.CombinedOutput()
	if err != nil && o.step.name != "NPM Outdated" && err.Error() != "exit status 1" {
		return err
	}

	fmt.Println(string(out))
	return nil
}

func (s step) handleErr(err error) error {
	return fmt.Errorf("Error during %v: %w\n", s.name, err)
}

func main() {
	composerInstall := newStep("Composer Install", "Installing composer packages...", exec.Command("composer", "install"))
	composerUpdate := newStep("Composer Update", "Updating composer packages...", exec.Command("composer", "update"))
	npmInstall := newStep("NPM Install", "Installing node packages...", exec.Command("npm", "install"))
	npmUpdate := newStep("NPM Update", "Updating node packages...", exec.Command("npm", "update"))
	migrate := newStep("Artisan Migrate", "Running a fresh migration...", exec.Command("php", "artisan", "migrate:fresh", "--env=testing"))

	steps := []step{
		composerInstall,
		composerUpdate,
		npmInstall,
		npmUpdate,
		migrate,
	}

	composerOutdated := newOutputStep("Composer Outdated",
		"\nThis is the state of the composer packages...\n",
		exec.Command("composer", "outdated", "-D"),
	)

	npmOutdated := newOutputStep("NPM Outdated",
		"\nThis is the state of the npm packages...\n",
		exec.Command("npm", "outdated"),
	)

	runTests := newOutputStep("PhpUnit Tests",
		"\nRunning all tests...\n",
		exec.Command("./vendor/bin/phpunit"),
	)

	outputSteps := []outputStep{
		composerOutdated,
		npmOutdated,
		runTests,
	}

	for _, step := range steps {
		if err := step.run(); err != nil {
			if err := step.handleErr(err); err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		}
	}

	for _, oStep := range outputSteps {
		if err := oStep.run(); err != nil {
			if err := oStep.step.handleErr(err); err != nil {
				fmt.Print(err)
				os.Exit(1)
			}
		}
	}
}
