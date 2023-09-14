package gun

type Musket struct {
	Gun
}

func NewMusket() IGun {
	return &Musket{
		Gun: Gun{
			Name:      "musket",
			Bullet:    25,
			MaxBullet: 25,
		},
	}
}
