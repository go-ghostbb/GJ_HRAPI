package identity

func Init() error {
	if err := employeeCodeInit(); err != nil {
		return err
	}

	if err := formOrderInit(); err != nil {
		return err
	}

	return nil
}
