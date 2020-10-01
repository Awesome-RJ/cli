package cmd

import (
	"context"
	"fmt"

	"github.com/railwayapp/cli/entity"
)

func (h *Handler) Whoami(ctx context.Context, req *entity.CommandRequest) error {
	user, err := h.ctrl.GetUser(ctx)
	if err != nil {
		return err
	}

	userText := fmt.Sprintf("%s", user.Email)
	if user.Name != "" {
		userText = fmt.Sprintf("%s (%s)", user.Name, user.Email)
	}

	fmt.Println(fmt.Sprintf("Hello 👋 %s", userText))

	// Todo, more info, also more fun
	return nil
}
