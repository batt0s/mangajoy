# MangaJoy

MangaJoy websitesi

FullStack MVC yapısı kullanılacak. Go (muhtemelen Gin framework) ile kodlanacak. Database olarak PostgreSQL (prod) ve SQLite (dev, test) kullanılacak ve orm olarak bun kullanılacak. Template engine html/template. Session managment için redis ve gin-contrib/sessions.

Kullanıcılar hesap oluşturup daha sonra giriş yapabilecek (google ile hesap açma ve giriş yapma seçenekleri de olacak). Mangalara yorum yazabilecek ve puan verebilecekler. Staff kullanıcılar (IsStaff : true) mangalara chapter ekleyebilecek. Mangaların tagları olacak. Kullanıcılar öneri ekleyebilecek ve bu öneriler manga sayfasında görünecek. Admin kullanıcılar (IsAdmin : true) için admin paneli olacak ve database de olan her şey oradan düzenlenebilecek. Anasayfa da son eklenen mangalar olacak ve başka bir sekmede en yüksek puanlı mangalar şeklinde bir sıralama olacak. Anasayfanın arama çubuğu sadece kelimeyi arayacak (belki tag filtreleme de), search sayfasında gelişmiş arama filtreleri (puan, tag vs.) olacak.

## TODO

- [x] Create github repo 
- [x] Make initial commit 
- [x] package database 
- [x] model : User 
    - [x] func : Save()
    - [x] func : Authenticate()

    | name | type |
    |:----:|:----:|
    | ID | uint | 
    | CreatedAt | time.Time | 
    | UpdatedAt | time.Time | 
    | DeletedAt | time.Time | 
    | LastLogin | time.Time | 
    | IsAdmin | bool | 
    | IsStaff | bool | 
    | Username | string |
    | Email | string | 
    | Password | string |

- [ ] model : Manga 
- [ ] controllers : User
    - [ ] Register
    - [ ] Login
    - [ ] Login, register with Google