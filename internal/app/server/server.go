package server

// Init initializes the router
func Init(webPath string) error {
	r := NewRouter(webPath)
	err := r.Run()
	if err != nil {
		return err
	}
	return nil
}
