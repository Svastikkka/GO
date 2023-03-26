package handlers

import (
	"net/http"

	"github.com/svastikkka/GO/BuildingModernWebApplicationsWithGo/33SettingApplicationWideConfiguration/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")

}
