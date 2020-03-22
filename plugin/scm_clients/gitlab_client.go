package scm_clients

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/drone/drone-go/drone"
	"github.com/google/uuid"
	"github.com/xanzy/go-gitlab"
)

type GitlabClient struct {
	delegate *gitlab.Client
	repo     drone.Repo
}

func NewGitLabClient(ctx context.Context, uuid uuid.UUID, server string, token string, repo drone.Repo) (ScmClient, error) {
	client := gitlab.NewClient(nil, token)
	if server != "" {
		client.SetBaseURL(server)
	}
	return GitlabClient{
		delegate: client,
		repo:     repo,
	}, nil
}

func (s GitlabClient) ChangedFilesInPullRequest(ctx context.Context, pullRequestID int) ([]string, error) {
	var changedFiles []string
	mr, _, err := s.listFiles(pullRequestID)
	if err != nil {
		return nil, err
	}
	changes := mr.Changes
	for _, file := range changes {
		if file.DeletedFile || file.RenamedFile {
			changedFiles = append(changedFiles, file.OldPath)
		}
		if !file.DeletedFile {
			changedFiles = append(changedFiles, file.NewPath)
		}
	}
	return changedFiles, nil
}

func (s GitlabClient) ChangedFilesInDiff(ctx context.Context, base string, head string) ([]string, error) {
	var changedFiles []string
	changes, _, err := s.compareCommits(base, head)
	if err != nil {
		return nil, err
	}
	for _, file := range changes.Diffs {
		if file.DeletedFile || file.RenamedFile {
			changedFiles = append(changedFiles, file.OldPath)
		}
		if !file.DeletedFile {
			changedFiles = append(changedFiles, file.NewPath)
		}
	}
	return changedFiles, nil
}

func (s GitlabClient) GetFileContents(ctx context.Context, path string, commitRef string) (content string, err error) {
	data, _, err := s.getContents(ctx, path, commitRef)
	if data == nil {
		err = fmt.Errorf("failed to get %s: is not a file", path)
	}
	fmt.Println(data)
	if err != nil {
		return "", err
	}
	return s.decode(data)
}

func (s GitlabClient) GetFileListing(ctx context.Context, path string, commitRef string) (
	fileListing []FileListingEntry, err error) {
	ls, _, err := s.getTree(ctx, path, commitRef)
	var result []FileListingEntry
	if err != nil {
		return result, err
	}
	for _, f := range ls {
		var fileType string
		if f.Type == "blob" {
			fileType = "file"
		} else if f.Type == "tree" {
			fileType = "dir"
		} else {
			continue
		}
		fileListingEntry := FileListingEntry{
			Path: f.Path,
			Name: f.Name,
			Type: fileType,
		}
		result = append(result, fileListingEntry)
	}
	return result, err
}

func (s GitlabClient) listFiles(id int) (*gitlab.MergeRequest, *gitlab.Response, error) {
	return s.delegate.MergeRequests.GetMergeRequestChanges(int(s.repo.ID), id)
}

func (s GitlabClient) compareCommits(base, head string) (
	*gitlab.Compare, *gitlab.Response, error) {
	opts := &gitlab.CompareOptions{
		From: &base,
		To:   &head,
	}
	return s.delegate.Repositories.Compare(int(s.repo.ID), opts)
}

func (s GitlabClient) getTree(ctx context.Context, path string, commitRef string) (
	fileContent []*gitlab.TreeNode, resp *gitlab.Response, err error) {
	opts := &gitlab.ListTreeOptions{
		Path: &path,
		Ref:  &commitRef,
	}
	t, _, e := s.delegate.Repositories.ListTree(int(s.repo.ID), opts)

	fmt.Println(t)
	return t, nil, e
}

func (s GitlabClient) getContents(ctx context.Context, path string, commitRef string) (
	fileContent *gitlab.File, resp *gitlab.Response, err error) {
	opts := &gitlab.GetFileOptions{
		Ref: &commitRef,
	}
	return s.delegate.RepositoryFiles.GetFile(int(s.repo.ID), path, opts)
}

func (s GitlabClient) decode(file *gitlab.File) (string, error) {
	var encoding string
	if file.Encoding != "" {
		encoding = file.Encoding
	}

	switch encoding {
	case "base64":
		c, err := base64.StdEncoding.DecodeString(*&file.Content)
		return string(c), err
	case "":
		if file.Content == "" {
			return "", nil
		}
		return file.Content, nil
	default:
		return "", fmt.Errorf("Unsupported content encoding: %v", encoding)
	}
}
