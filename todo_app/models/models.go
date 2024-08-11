package models

type Todo struct {
      Id      int   `json:"Id"`
    Title   string `json:"Title"`
    Content string `json:"Content"`
    Status  string `json:"Status"`
}

