package server

// Init initializes the router
func Init() error {
	r := NewRouter()
	err := r.Run()
	if err != nil {
		return err
	}
	return nil
}
