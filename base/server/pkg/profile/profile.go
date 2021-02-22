package profile

import (
	"github.com/pkg/profile"
	"os"
)

func Start(mode string) {
	switch mode {
	case "cpu":
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(os.Getenv("HOME"))).Stop()
	case "mem":
		defer profile.Start(profile.MemProfile, profile.ProfilePath(os.Getenv("HOME"))).Stop()
	case "mutex":
		defer profile.Start(profile.MutexProfile, profile.ProfilePath(os.Getenv("HOME"))).Stop()
	case "block":
		defer profile.Start(profile.BlockProfile, profile.ProfilePath(os.Getenv("HOME"))).Stop()
	default:
		defer profile.Start(profile.CPUProfile, profile.MemProfile, profile.MutexProfile, profile.BlockProfile, profile.ProfilePath(os.Getenv("HOME"))).Stop()
	}
}
