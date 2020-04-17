package commands

type CreateFight struct {
	Heroes []string
}

func NewCreateFight(heroes []string) *CreateFight {
	return &CreateFight{heroes}
}

func (c CreateFight) GetHeroes() []string {
	return c.Heroes
}

type UpdateFight struct {
	Heroes []string
}

func NewUpdateFight(heroes []string) *UpdateFight {
	return &UpdateFight{heroes}
}

func (u UpdateFight) GetHeroes() []string {
	return u.Heroes
}
