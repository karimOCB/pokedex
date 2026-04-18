package main



func commandMapf(cfg *config) error {
	locations, err := cfg.Client.ListLocations(cfg.Next)
	if err != nil {
		
	}
	cfg.Next = locations.Next
}

func commandMapb(cfg *config) error {
	
}