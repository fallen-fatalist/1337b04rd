package domain

type Post struct {
	ID        uint64 `json:"post_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Avatar    string `json:"avatar"`
	Image     string `json:"image,omitempty"`
	CreatedAt string `json:"created_at"`
}
