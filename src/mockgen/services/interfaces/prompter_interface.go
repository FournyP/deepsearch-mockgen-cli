package interfaces

type PrompterInterface interface {
	Prompt(question, defaultValue string) (string, error)
}
