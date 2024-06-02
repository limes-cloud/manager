package menu

type ListMenuRequest struct {
	Order   *string `json:"order"`
	OrderBy *string `json:"orderBy"`
	Title   *string `json:"title"`
}
