package helper

import (
	"errors"
	"fmt"
)

func Org(team, comp int) (string, error) {
	t := Team(team)
	if t == "" {
		return "", errors.New("invalid team")
	}
	c := Components(comp)
	if c == "" {
		return "", errors.New("invalid component")
	}

	return fmt.Sprintf("Welcome to %s of %s", c, t), nil
}

func Team(team int) string {
	switch team {
	case 1:
		return "Schematics"
	case 2:
		return "Event Notification"
	case 3:
		return "AppConfig"
	default:
		return "Paas"
	}
}

func Components(component int) string {
	switch component {
	case 1:
		return "Workspace"
	case 2:
		return "Actions"
	case 3:
		return "Agent"
	case 4:
		return "Inventory"
	case 5:
		return "config"
	case 6:
		return "events"
	case 7:
		return "notifier"
	}

	return ""
}
