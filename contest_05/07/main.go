type Teapot struct {
	temp int
}

type Water struct {
	temp int
}

func (t Teapot) is_boiling() bool {
	return t.temp >= 100
}

func NewWater(aboba int) Water {
	return Water{aboba}
}

func (t *Teapot) heat_up(add_temp int) {
	t.temp += add_temp
}

func NewTeapot(water Water) Teapot {
	return Teapot(water)
}
