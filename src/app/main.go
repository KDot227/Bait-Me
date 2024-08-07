package main

import (
	"bait-me/console"
	"bait-me/icons"
	"bait-me/privileges"
	"bait-me/regkeys"
	"bait-me/services"
	"fmt"

	"github.com/getlantern/systray"
	"github.com/getlantern/systray/example/icon"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	privileges.EnsureAdmin()

	if !regkeys.BackupExists() {
		fmt.Println("Backing up registry keys...")
		regkeys.BackupRegKeys()
	}

	console.HideConsole()

	systray.SetTemplateIcon(icon.Data, icons.Icon)
	systray.SetTitle("Bait Me")
	systray.SetTooltip("Bait Me")

	mShowConsole := systray.AddMenuItem("Show Console", "Show the console")
	mHideConsole := systray.AddMenuItem("Hide Console", "Hide the console")
	mStartServices := systray.AddMenuItem("Start Services", "Start the services")
	mStopServices := systray.AddMenuItem("Stop Services", "Stop the services")
	mRemoveServices := systray.AddMenuItem("Remove Services", "Remove the services")
	mRestore := systray.AddMenuItem("Restore Registry Keys", "Restore Registry Keys")
	mChangeRegKeys := systray.AddMenuItem("Change Registry Keys", "Change the registry keys")
	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-mShowConsole.ClickedCh
		fmt.Println("Showing the console...")
		console.ShowConsole()
	}()

	go func() {
		<-mHideConsole.ClickedCh
		fmt.Println("Hiding the console...")
		console.HideConsole()
	}()

	go func() {
		<-mStartServices.ClickedCh
		fmt.Println("Starting the services...")
		services.StartServices()
	}()

	go func() {
		<-mStopServices.ClickedCh
		fmt.Println("Stopping the services...")
		services.StopServices()
	}()

	go func() {
		<-mRemoveServices.ClickedCh
		fmt.Println("Removing the services...")
		services.RemoveServices()
	}()

	go func() {
		<-mRestore.ClickedCh
		fmt.Println("Restoring the registry keys...")
		regkeys.RestoreRegKeys()
	}()

	go func() {
		<-mChangeRegKeys.ClickedCh
		fmt.Println("Changing the registry keys...")
		regkeys.ChangeRegKeys()
	}()

	go func() {
		<-mQuitOrig.ClickedCh
		fmt.Println("Requesting quit")
		systray.Quit()
		fmt.Println("Finished quitting")
	}()
}

func onExit() {
	fmt.Println("Exiting...")
}
