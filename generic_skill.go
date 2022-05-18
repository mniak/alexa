package alexa

import (
	"context"

	alexakit "github.com/ericdaugherty/alexa-skills-kit-golang"
)

type alexaKitSkillEventFunc func(context.Context, *alexakit.Request, *alexakit.Session, *alexakit.Context, *alexakit.Response) error

type genericSkill struct {
	onSessionStarted alexaKitSkillEventFunc
	onLaunch         alexaKitSkillEventFunc
	onIntent         alexaKitSkillEventFunc
	onSessionEnded   alexaKitSkillEventFunc
}

func (gs genericSkill) OnSessionStarted(ctx context.Context, req *alexakit.Request, sess *alexakit.Session, actx *alexakit.Context, resp *alexakit.Response) error {
	return gs.onSessionStarted(ctx, req, sess, actx, resp)
}

func (gs genericSkill) OnLaunch(ctx context.Context, req *alexakit.Request, sess *alexakit.Session, actx *alexakit.Context, resp *alexakit.Response) error {
	return gs.onLaunch(ctx, req, sess, actx, resp)
}

func (gs genericSkill) OnIntent(ctx context.Context, req *alexakit.Request, sess *alexakit.Session, actx *alexakit.Context, resp *alexakit.Response) error {
	return gs.onIntent(ctx, req, sess, actx, resp)
}

func (gs genericSkill) OnSessionEnded(ctx context.Context, req *alexakit.Request, sess *alexakit.Session, actx *alexakit.Context, resp *alexakit.Response) error {
	return gs.onSessionEnded(ctx, req, sess, actx, resp)
}
