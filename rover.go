package main

type rover struct {
	X         uint
	Y         uint
	Direction rune
	Commands  []rune
	Plateau   *plateau
}

func NewRover(pos, commands string) (*rover, error) {
	x, y, direction, err := convertPosition(p)
	if err != nil {
		return nil, err
	}

	commands, err := convertCommands(c)
	if err != nil {
		return nil, err
	}

	return &rover{x, y, direction, commands, nil}, nil
}

func (r *rover) Deploy(p *plateau) error {
	err := p.put(r.x, r.y, r)
	if err != nil {
		return err
	}

	r.Plateau = p

	return nil
}

func (r *rover) Run() error {
	for _, c := range r.Commands {
		err := r.applyCommand(c)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *rover) applyCommand(c rune) error {
	switch c {
	case 'R':
		switch r.Direction {
		case 'N':
			r.Direction = 'E'
		case 'E':
			r.Direction = 'S'
		case 'S':
			r.Direction = 'W'
		case 'W':
			r.Direction = 'N'
		}
	case 'L':
		switch r.Direction {
		case 'N':
			r.Direction = 'W'
		case 'E':
			r.Direction = 'N'
		case 'S':
			r.Direction = 'E'
		case 'W':
			r.Direction = 'S'
		}
	case 'M':
		newX := r.X
		newY := r.Y
		switch r.Direction {
		case 'N':
			newY = r.Y + 1
		case 'E':
			newX = r.X + 1
		case 'S':
			newY = r.Y - 1
		case 'W':
			newX = r.X - 1
		}

		err := r.Plateau.Update(r.X, r.Y, newX, newY, r)
		if err != nil {
			return err
		}

		r.X = newX
		r.Y = newY
	}

	return nil
}
