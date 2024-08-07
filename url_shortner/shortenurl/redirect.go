package shortenurl

import (
	"net/http"
	"gorm.io/gorm"
)

type URL struct {
    ID uint `gorm:"primary_key"`
    Original string `gorm:"not_null"`
    Shortened string `gorm:"not_null"`
}

func RedirectUrl (db *gorm.DB, w http.ResponseWriter, r *http.Request) {
    id := r.URL.Path[1:]
    var url URL
    shortend := "http://localhost:8080/" + id
    db.First(&url, "shortend = ?", shortend)
    http.Redirect(w, r, url.Original, http.StatusFound)
}
