package main

func main() {
	config, err := LoadConfig("config.yaml")
	if err != nil {
		panic(err)
	}
	repo, err := LoadData("data.json")
	if err != nil {
		panic(err)
	}
	RunBot(config, repo)
}
