package interfaces

type Post struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Content string
	Author  Author `toml:"author"`
}

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}
