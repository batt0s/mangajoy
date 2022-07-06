package models

func Migrate() {
	CreateUserTable()
	CreateArtistTable()
	CreateMangaTable()
	CreateChapterTable()
	CreateCommentTable()
}
