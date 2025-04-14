package gui

func MakeWindowLoop() (func(), error) {
	aboutWnd, err := makeAboutWindowLoop()
	if err != nil {
		return nil, err
	}

	return func() {
		menuBarLoop()

		aboutWnd()
		convertVorbisWindowLoop()
	}, nil
}
