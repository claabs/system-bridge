//go:build linux

package power

import (
	"os/exec"

	"github.com/timmo001/system-bridge/settings"
)

func shutdown() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Poweroff
	if len(command) == 0 {
		command = []string{"systemctl", "poweroff"}
	}
	cmd := exec.Command(command[0], command[1:]...)
	return cmd.Run()
}

func restart() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Reboot
	if len(command) == 0 {
		command = []string{"systemctl", "reboot"}
	}
	cmd := exec.Command(command[0], command[1:]...)
	return cmd.Run()
}

func sleep() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Suspend
	if len(command) == 0 {
		command = []string{"systemctl", "suspend"}
	}
	cmd := exec.Command(command[0], command[1:]...)
	return cmd.Run()
}

func hibernate() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Hibernate
	if len(command) == 0 {
		command = []string{"systemctl", "hibernate"}
	}
	cmd := exec.Command(command[0], command[1:]...)
	return cmd.Run()
}

func lock() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Lock
	if len(command) == 0 {
		command = []string{"loginctl", "lock-session"}
	}
	cmd := exec.Command(command[0], command[1:]...)
	return cmd.Run()
}

func logout() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Logout
	if len(command) == 0 {
		command = []string{"loginctl", "terminate-session", "$XDG_SESSION_ID"}
	}
	cmd := exec.Command(command[0], command[1:]...)
	return cmd.Run()
}
