package website

import (
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

type Message struct {
	Author  string
	Content string
	Date    time.Time
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/sign/", sign_handler)
}

func messageKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Message", "default_message", 0, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Message").Ancestor(messageKey(c)).Order("-Date").Limit(10)

	messages := make([]Message, 0, 10)
	if _, err := q.GetAll(c, &messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := messageTemplate.Execute(w, messages); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var messageTemplate = template.Must(template.New("log").Parse(`
<html>
	<head>
		<title>Go Guestbook</title>
	</head>
	<body>
	{{range .}}
		<p><b>{{.Author}}</b> wrote:</p>
		<pre>{{.Content}}</pre>
	{{end}}
		<form action="/sign/" method="post">
			<div><input type="text" name="author" placeholder="Author"></div>
			<div><textarea name="content" placeholder="Message" rows="3" cols="60"></textarea></div>
			<div><input type="submit" value="Sign Guestbook"></div>
		</form>
	</body>
</html>
`))

func sign_handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	m := Message{
		Author:  r.FormValue("author"),
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}
	key := datastore.NewIncompleteKey(c, "Message", messageKey(c))
	_, err := datastore.Put(c, key, &m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
