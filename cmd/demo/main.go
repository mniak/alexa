package main

import (
	"context"
	"net/http"

	"github.com/mniak/alexa"
)

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	skillName := "The Simple Example"
	appID := "amzn1.ask.skill.<SKILL_ID>"

	handler := alexa.NewSkillBuilder(appID).
		SetIgnoreApplicationID(true).
		SetIgnoreTimestamp(true).
		WithOnSessionStarted(func(ctx context.Context, req *alexa.Request, session *alexa.Session, context *alexa.Context, resp *alexa.Response) error {
			message := "Skill started"
			resp.SetOutputText(message)
			resp.SetSimpleCard(skillName, message)
			resp.ShouldSessionEnd = false
			return nil
		}).
		WithOnIntent(func(ctx context.Context, req *alexa.Request, session *alexa.Session, context *alexa.Context, resp *alexa.Response) error {
			message := "Doing skill stuff"
			resp.SetOutputText(message)
			resp.SetSimpleCard(skillName, message)
			resp.ShouldSessionEnd = true
			return nil
		}).
		BuildHttpHandler()

	handler(w, r)
}
