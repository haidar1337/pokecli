package api

type Config struct {
	Next *string
	Previous *string
}

var apiUrl string = baseUrl + "/location-area" 
var config Config = Config{
	Next: &apiUrl,
	Previous: nil,
}

func GetConfig() Config {
	return config
}

func setNext(n *string) {
	config.Next = n
}

func setPrev(p *string) {
	config.Previous = p
}


func setConfig(next *string, prev *string) {
	setNext(next)
	setPrev(prev)
}