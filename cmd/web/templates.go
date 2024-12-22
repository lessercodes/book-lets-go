package main

import "snippetbox.lessercodes.com/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
