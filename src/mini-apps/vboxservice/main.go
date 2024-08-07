package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/kardianos/service"
)

type program struct{}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
	fmt.Println("I am running...")
	fmt.Scanln()
}
func (p *program) Stop(s service.Service) error {
	fmt.Println("I am stopping...")
	return nil
}

func main() {

	var mode string
	flag.StringVar(&mode, "mode", "", "install/uninstall/run")
	flag.Parse()

	svcConfig := &service.Config{
		Name:        "vboxservice",
		DisplayName: "Vboxservice Bait-Me",
		Description: "This is a fake service for Bait-Me",
	}

	prg := &program{}
	s, err := service.New(prg, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}

	if mode == "install" {
		err = s.Install()
		if err != nil {
			log.Fatal(err)
		}
	}

	if mode == "uninstall" {
		err = s.Uninstall()
		if err != nil {
			log.Fatal(err)
		}
	}

	if mode == "" || mode == "run" {
		err = s.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
