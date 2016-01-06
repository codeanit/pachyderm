/*
Package drive provides the definitions for the low-level pfs storage drivers.
*/
package drive

import (
	"io"

	"go.pedge.io/google-protobuf"

	"github.com/pachyderm/pachyderm/src/pfs"
)

// ReaderAtCloser is an interface that implements both io.ReaderAt and io.Closer.
type ReaderAtCloser interface {
	io.Reader
	io.ReaderAt
	io.Closer
}

// Driver represents a low-level pfs storage driver.
type Driver interface {
	CreateRepo(repo *pfs.Repo, created *google_protobuf.Timestamp, shards map[uint64]bool) error
	InspectRepo(repo *pfs.Repo, shards map[uint64]bool) (*pfs.RepoInfo, error)
	ListRepo(shards map[uint64]bool) ([]*pfs.RepoInfo, error)
	DeleteRepo(repo *pfs.Repo, shards map[uint64]bool) error
	StartCommit(parent *pfs.Commit, commit *pfs.Commit, started *google_protobuf.Timestamp, shards map[uint64]bool) error
	FinishCommit(commit *pfs.Commit, finished *google_protobuf.Timestamp, shards map[uint64]bool) error
	InspectCommit(commit *pfs.Commit, shards map[uint64]bool) (*pfs.CommitInfo, error)
	ListCommit(repo []*pfs.Repo, fromCommit []*pfs.Commit, shards map[uint64]bool) ([]*pfs.CommitInfo, error)
	DeleteCommit(commit *pfs.Commit, shards map[uint64]bool) error
	PutBlock(file *pfs.File, block *Block, shard uint64, reader io.Reader) error
	GetBlock(block *Block, shard uint64) (ReaderAtCloser, error)
	InspectBlock(block *Block, shard uint64) (*BlockInfo, error)
	ListBlock(shard uint64) ([]*BlockInfo, error)
	PutFile(file *pfs.File, shard uint64, offset int64, reader io.Reader) error
	MakeDirectory(file *pfs.File, shards map[uint64]bool) error
	GetFile(file *pfs.File, shard uint64) (ReaderAtCloser, error)
	InspectFile(file *pfs.File, shard uint64) (*pfs.FileInfo, error)
	ListFile(file *pfs.File, shard uint64) ([]*pfs.FileInfo, error)
	ListChange(file *pfs.File, from *pfs.Commit, shard uint64) ([]*pfs.Change, error)
	DeleteFile(file *pfs.File, shard uint64) error
	PullDiff(commit *pfs.Commit, shard uint64, diff io.Writer) error
	PushDiff(commit *pfs.Commit, shard uint64, diff io.Reader) error
}
