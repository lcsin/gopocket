package cmd

import (
	"context"
	"fmt"
	"os/exec"
)

const (
	git   = "git"
	clone = "clone"
	pull  = "pull"
	push  = "push"
)

func GitClone(ctx context.Context, url, branch string, path string) error {
	var cmd *exec.Cmd
	if branch == "" {
		cmd = exec.CommandContext(ctx, git, clone, url, path)
	} else {
		cmd = exec.CommandContext(ctx, git, clone, "-b", branch, url, path)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))

	return nil
}

func GitPull(ctx context.Context, path string) error {
	cmd := exec.CommandContext(ctx, git, pull)
	if path != "" {
		cmd.Dir = path
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(string(output))

	return nil
}
