package handlers

import (
	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/29WorkingWithLayouts/pkg/render"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

}
