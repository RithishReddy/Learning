package erratum

func Use(opener ResourceOpener, input string) (err error) {

LOOP:
	resource, err := opener()

	if err != nil {
		if _, ok := err.(TransientError); ok {
			goto LOOP
		} else {
			return err
		}

	}
	

	defer resource.Close()

	defer func() {
		if r := recover(); r != nil {

			if _, ok := r.(FrobError); ok {
				fr := r.(FrobError)
				resource.Defrob(fr.defrobTag)

			}
			err = r.(error)

		}

	}()

	resource.Frob(input)
	return nil
}
