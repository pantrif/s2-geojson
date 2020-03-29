package server

// Init initializes the router
func Init(webPath string, addr ...string) error {
	r := NewRouter(webPath)
	err := r.Run(addr...)
	if err != nil {
		return err
	}
	return nil
}
