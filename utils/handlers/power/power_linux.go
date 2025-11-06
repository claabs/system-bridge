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
	if command == "" {
		command = "systemctl poweroff"
	}
	cmd := exec.Command(command)
	return cmd.Run()
}

func restart() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Reboot
	if command == "" {
		command = "systemctl reboot"
	}
	cmd := exec.Command(command)
	return cmd.Run()
}

func sleep() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Suspend
	if command == "" {
		command = "systemctl suspend"
	}
	cmd := exec.Command(command)
	return cmd.Run()
}

func hibernate() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Hibernate
	if command == "" {
		command = "systemctl hibernate"
	}
	cmd := exec.Command(command)
	return cmd.Run()
}

func lock() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Lock
	if command == "" {
		command = "loginctl lock-session"
	}
	cmd := exec.Command(command)
	return cmd.Run()
}

func logout() error {
	s, err := settings.Load()
	if err != nil {
		return err
	}

	command := s.PowerCommands.Logout
	if command == "" {
		command = "loginctl terminate-session $XDG_SESSION_ID"
	}
	cmd := exec.Command(command)
	return cmd.Run()
}
