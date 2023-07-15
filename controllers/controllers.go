package controllers

import (
	"net/http"

	"github.com/muhardiansyah15/go-rest-api/helper"
)

func Index(w http.ResponseWriter, r *http.Request) {

	data := []map[string]interface{}{
		{
			"title":  "About Me",
			"designation": "Software Developer",
			"description":  "Hello, my name is Muhardiansyah. I am a Software Engineer and hold a Bachelor of Science degree in Mathematics. With over three years of experience as a software engineer at a software house company, I have developed expertise in ERP software development. My focus has been on backend development, where I have acquired proficiency in Accounting, Manufacturing Resource Planning (MRP), Inventory Management, Sales, and Customer Relationship Management (CRM) functionalities within ERP systems. I possess strong programming skills in Python, PostgreSQL, and MySQL, and I am well-versed in web development technologies such as PHP, HTML, and CSS.",
		},
	}

	helper.ResponseJSON(w, http.StatusOK, data)

}
