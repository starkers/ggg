package git_test

import (
	"github.com/starkers/ggg/pkg/git"
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		name     string
		baseDir  string
		inputURL string
		want     string
	}{
		{name: "githubHttpSimple",
			baseDir:  "/home/user/src",
			inputURL: "https://github.com/acmecorp/codebase",
			want:     "/home/user/src/github.com/acmecorp/codebase"},
		{name: "githubHttpWithTrailing.git",
			baseDir:  "/home/user/src",
			inputURL: "https://github.com/acmecorp/codebase.git",
			want:     "/home/user/src/github.com/acmecorp/codebase"},
		{name: "githubGitWithTrailing.git",
			baseDir:  "/home/user/src",
			inputURL: "git@github.com:foo/bar.git",
			want:     "/home/user/src/github.com/foo/bar"},
		{name: "githubGitWithTrailing.git",
			baseDir:  "/home/user/src",
			inputURL: "git@github.com:foo/bar.git",
			want:     "/home/user/src/github.com/foo/bar"},
		{name: "gitLabDeep",
			baseDir:  "/home/user/src",
			inputURL: "git@gitlab.com:librewolf-community/browser/linux",
			want:     "/home/user/src/gitlab.com/librewolf-community/browser/linux"},
		{name: "gitLabDeepWithTrailingGit",
			baseDir:  "/home/user/src",
			inputURL: "git@gitlab.com:librewolf-community/browser/linux.git",
			want:     "/home/user/src/gitlab.com/librewolf-community/browser/linux"},
		{name: "gitLabDeepHttpSimple",
			baseDir:  "/home/user/src",
			inputURL: "https://gitlab.com/librewolf-community/website",
			want:     "/home/user/src/gitlab.com/librewolf-community/website"},
		{name: "gitLabDeepHttpWithTrailingGit",
			baseDir:  "/home/user/src",
			inputURL: "https://gitlab.com/librewolf-community/website.git",
			want:     "/home/user/src/gitlab.com/librewolf-community/website"},
	}

	for _, tc := range tests {
		got, _ := git.FigureUnixDiskPath(tc.baseDir, tc.inputURL)
		// TODO: test the errors
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("%s: expected: %v, got: %v", tc.name, tc.want, got)
		}
	}
}
