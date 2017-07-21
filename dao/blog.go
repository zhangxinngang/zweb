package dao

type Blog struct {
	Id         string `json:"id"`
	Editor     string `json:"editor"`
	Blog       string `json:"blog"`
	CreateTime int64  `json:"create_time"` // second
}

func GetBlog(id string) Blog {
	return GlogMap[id]
}

func AddBlog(blog Blog) {
	GlogMap[blog.Id] = blog
}

func DelBlog(id string) {
	delete(GlogMap, id)
}
