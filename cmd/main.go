package cmd

import "log"

func main() {
	srv := new(cmd.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("Error")
	}
}
