package gun

type AK47 struct {
	Gun
}

func NewAK47() IGun {
	return &AK47{
		Gun: Gun{
			Name:      "AK47",
			Bullet:    10,
			MaxBullet: 10,
		},
	}
}
