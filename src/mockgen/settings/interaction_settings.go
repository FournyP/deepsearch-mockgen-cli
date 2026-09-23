package settings

type InteractionSettings struct {
	AcceptAll      bool
	SkipPathPrompt bool
}

func NewInteractionSettings(acceptAll, skipPathPrompt bool) *InteractionSettings {
	return &InteractionSettings{AcceptAll: acceptAll, SkipPathPrompt: skipPathPrompt}
}
