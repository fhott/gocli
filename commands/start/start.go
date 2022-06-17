package start

import (
	"flag"
	"fmt"
	"net/http"
)

type Start struct {
	Port    int
	Version string
	HelpF   bool
}

func (s *Start) Name() string {
	return "start"
}

func (s *Start) Example() string {
	return `
	gocli start --port 3000 --version 1.0.0
	gocli start --version 1.0.0
	`
}

func (s *Start) Help() string {
	return `Responsable for initialize the application`
}

func (s *Start) LongHelp() string {
	return `Responsable for initialize the application but longer`
}

func (s *Start) Register(fs *flag.FlagSet) {
	fs.IntVar(&s.Port, "port", 8080, "Server's port")
	fs.StringVar(&s.Version, "version", "", "Server's version")
	fs.BoolVar(&s.HelpF, "help", false, "Help command")
}

func (s *Start) Run() {
	if s.HelpF {
		fmt.Println(s.LongHelp())
		return
	}

	if s.Version == "" {
		fmt.Println("[--version] é obrigatório")
		return
	}

	fmt.Printf("Server v%s running on port %v", s.Version, s.Port)
	http.ListenAndServe(fmt.Sprintf(":%v", s.Port), nil)
}
