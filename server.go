package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/billzajac/rltao"

	"time"
)

func GenerateHTML() string {
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	id := rltao.GenerateId()
	tetragram := rltao.GetTetragram(id)
	passage := rltao.GetPassage(id)
	html := fmt.Sprintf(`<html>
<head><style>
/* === BASE HEADING === */
h1 { position: relative; padding: 0; margin: 0; font-family: "Raleway", sans-serif; font-weight: 300; font-size: 40px; color: #080808; -webkit-transition: all 0.4s ease 0s; -o-transition: all 0.4s ease 0s; transition: all 0.4s ease 0s; }
h1 span { display: block; font-size: 0.5em; line-height: 1.3; }
h1 em { font-style: normal; font-weight: 600; }

/* === HEADING STYLE === */
.title h1 { text-transform: capitalize; }
.title h1:before { position: absolute; left: 0; bottom: 0; width: 60px; height: 2px; content: ""; background-color: #c50000; }
.title h1 span { font-size: 13px; font-weight: 500; text-transform: uppercase; letter-spacing: 4px; line-height: 3em; padding-left: 0.25em; color: rgba(0, 0, 0, 0.4); padding-bottom: 10px; }
.footer { font-size: 11px; font-weight: 500; text-transform: uppercase; line-height: 3em; padding-left: 0.25em; color: rgba(0, 0, 0, 0.4); padding-bottom: 10px; }
</style></head>
<body>
<div class="title"><h1>%02d %s
<span>%s</span>
</h1></div>
<br>
%s
<br><br>
<span class="footer">%s</span>
</body>
</html>`, id, tetragram, passage.Title, passage.Body, now)
	return html
}

func generateTetragram(w http.ResponseWriter, r *http.Request) {
	body := GenerateHTML()
	fmt.Fprintf(w, "%s", body)
}

func main() {
	http.HandleFunc("/", generateTetragram)  // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
