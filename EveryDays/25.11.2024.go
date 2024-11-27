package EveryDays

import "os/exec"

func beep() {
	var cmd *exec.Cmd
	for i := 0; i < 100; i++ {
		cmd = exec.Command("powershell", "[console]::beep(800,200)")
		cmd.Run()
	}
}
