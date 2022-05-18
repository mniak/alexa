package alexa

import (
	"context"
	"net/http"

	alexakit "github.com/ericdaugherty/alexa-skills-kit-golang"
)

type (
	// Request contains the data in the request within the main request.
	Request struct{ *alexakit.Request }
	// Response contains the body of the response.
	Response struct{ *alexakit.Response }
	// Session contains the session data from the Alexa request.
	Session struct{ *alexakit.Session }
	// Context contains the context data from the Alexa Request.
	Context struct{ *alexakit.Context }
	// SkillEventFunc is an event handler funcion type
	SkillEventFunc func(context.Context, *Request, *Session, *Context, *Response) error
)

// NewSkillBuilder creates a skill builder
func NewSkillBuilder(appID string) *SkillBuilder {
	return &SkillBuilder{
		ApplicationID: appID,
	}
}

// SkillBuilder is used to fluently create Alexa skill handlers
type SkillBuilder struct {
	ApplicationID       string
	IgnoreApplicationID bool
	IgnoreTimestamp     bool

	onSessionStarted SkillEventFunc
	onLaunch         SkillEventFunc
	onIntent         SkillEventFunc
	onSessionEnded   SkillEventFunc
}

// SetIgnoreApplicationID sets the field IgnoreApplicationID
func (sb *SkillBuilder) SetIgnoreApplicationID(ignore bool) *SkillBuilder {
	sb.IgnoreApplicationID = ignore
	return sb
}

// SetIgnoreTimestamp sets the field IgnoreTimestamp
func (sb *SkillBuilder) SetIgnoreTimestamp(ignore bool) *SkillBuilder {
	sb.IgnoreTimestamp = ignore
	return sb
}

// WithOnSessionStarted sets the OnSessionStarted event handler
func (sb *SkillBuilder) WithOnSessionStarted(fn SkillEventFunc) *SkillBuilder {
	sb.onSessionStarted = fn
	return sb
}

// WithOnLaunch sets the OnLaunch event handler
func (sb *SkillBuilder) WithOnLaunch(fn SkillEventFunc) *SkillBuilder {
	sb.onLaunch = fn
	return sb
}

// WithOnIntent sets the OnIntent event handler
func (sb *SkillBuilder) WithOnIntent(fn SkillEventFunc) *SkillBuilder {
	sb.onIntent = fn
	return sb
}

// WithOnSessionEnded sets the OnSessionEnded event handler
func (sb *SkillBuilder) WithOnSessionEnded(fn SkillEventFunc) *SkillBuilder {
	sb.onSessionEnded = fn
	return sb
}

type skillFunc func(ctx context.Context, requestEnv *alexakit.RequestEnvelope) (*alexakit.ResponseEnvelope, error)

func (sb SkillBuilder) buildSkillHandler() skillFunc {
	handler := genericSkill{
		onSessionStarted: coalesceAlexaFunc(sb.onSessionStarted),
		onLaunch:         coalesceAlexaFunc(sb.onLaunch),
		onIntent:         coalesceAlexaFunc(sb.onIntent),
		onSessionEnded:   coalesceAlexaFunc(sb.onSessionEnded),
	}

	skill := alexakit.Alexa{
		ApplicationID:       sb.ApplicationID,
		IgnoreApplicationID: sb.IgnoreApplicationID,
		IgnoreTimestamp:     sb.IgnoreTimestamp,
		RequestHandler:      handler,
	}
	return skill.ProcessRequest
}

// BuildHTTPHandler returns a HandlerFunc ready to process Skill requests
func (sb SkillBuilder) BuildHTTPHandler() http.HandlerFunc {
	skillHandler := sb.buildSkillHandler()
	return skillHandler.MakeHTTPHandler()
}
