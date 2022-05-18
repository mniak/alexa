package alexa

import (
	"context"

	alexakit "github.com/ericdaugherty/alexa-skills-kit-golang"
)

func emptyAlexaKitSkillEvent(context.Context, *alexakit.Request, *alexakit.Session, *alexakit.Context, *alexakit.Response) error {
	return nil
}

func (fn SkillEventFunc) toAlexaKitFunc() alexaKitSkillEventFunc {
	return func(ctx context.Context, akreq *alexakit.Request, aksess *alexakit.Session, akactx *alexakit.Context, akresp *alexakit.Response) error {
		var req *Request
		if akreq != nil {
			req = &Request{*akreq}
		}
		var sess *Session
		if aksess != nil {
			sess = &Session{*aksess}
		}
		var actx *Context
		if akactx != nil {
			actx = &Context{*akactx}
		}
		var resp *Response
		if akresp != nil {
			resp = &Response{*akresp}
		}
		return fn(ctx, req, sess, actx, resp)
	}
}

func coalesceAlexaFunc(fn SkillEventFunc) alexaKitSkillEventFunc {
	if fn != nil {
		return fn.toAlexaKitFunc()
	}
	return emptyAlexaKitSkillEvent
}
