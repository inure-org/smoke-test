package main

import (
	"fmt"
	"log"
	"os"

	"github.com/inure-org/smoke-test/internal/config"
	"github.com/inure-org/smoke-test/internal/executor"
)

func init() {
	err := config.InitializeConfig()

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	executor := executor.PlatformExecutor{}

	err := executor.RunSmokeTest()

	if err != nil {
		fmt.Printf("runsmoketest falhou: %v\n", err)

		err := notify(err)

		if err != nil {
			log.Printf("notify falhou com o erro: %v", err)
		}

		os.Exit(1)
	}

	log.Printf("testes smoke executados com sucesso.")

	os.Exit(0)
}

func notify(err error) error {
	log.Println("notificado: ", err)

	return nil
}
