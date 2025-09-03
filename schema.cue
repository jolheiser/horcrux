import "time"

#Repo: {
	name: string
	source: string
	dest: [...string]
}

#Schema: {
	key: string
	interval: time.Duration | *"1h"
	storage: string | *".horcrux"
	repos: [...#Repo]
}

#Schema
