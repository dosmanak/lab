package cmd

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_mrListAssignedTo(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "--assignee=zaquestion")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	t.Log(mrs)
	require.Contains(t, mrs, "!1 Test MR for lab list")
	require.NotContains(t, mrs, "!3")
	require.NotContains(t, mrs, "filtering with labels and lists")
}

func Test_mrList(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	t.Log(mrs)
	require.Contains(t, mrs, "!1 Test MR for lab list")
}

func Test_mrListFlagLabel(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "-l", "confirmed")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	t.Log(mrs)
	require.Contains(t, mrs, "!3 for testings filtering with labels and lists")
}

func Test_mrListStateMerged(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "-s", "merged")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	t.Log(mrs)
	require.Contains(t, mrs, "!4 merged merge request")
}

func Test_mrListStateClosed(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "-n", "1", "-s", "closed")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	require.Regexp(t, `!\d+`, string(b))
}

func Test_mrListFivePerPage(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "-n", "5")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	mrs := getAppOutput(b)
	require.Len(t, mrs, 5)
}

func Test_mrFilterByTargetBranch(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "-t", "non-existing")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := getAppOutput(b)
	assert.Empty(t, mrs, "Expected to find no MRs for non-existent branch")
}

var (
	latestCreatedTestMR = "!968 README: dummy commit for CI tests"
	latestUpdatedTestMR = "!329 MR for assign and review commands"
)

func Test_mrListByTargetBranch(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "-t", "master")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	require.Equal(t, latestUpdatedTestMR, mrs[0])
}

// updated,asc
// !1
func Test_mrListUpdatedAscending(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "--number=1", "--order=updated_at", "--sort=asc")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	t.Log(mrs)
	require.Contains(t, mrs, "!3 for testings filtering with labels and lists")
}

// updatead,desc
// !18
func Test_mrListUpdatedDescending(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "--number=1", "--order=updated_at", "--sort=desc")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	t.Log(mrs)
	require.Equal(t, latestUpdatedTestMR, mrs[0])
}

// created,asc
// !1
func Test_mrListCreatedAscending(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "--number=1", "--order=created_at", "--sort=asc")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	t.Log(mrs)
	require.Contains(t, mrs, "!1 Test MR for lab list")
}

// created,desc
// !18
func Test_mrListCreatedDescending(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "--number=1", "--order=created_at", "--sort=desc")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	t.Log(mrs)
	require.Equal(t, latestCreatedTestMR, mrs[0])
}

func Test_mrListSearch(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	cmd := exec.Command(labBinaryPath, "mr", "list", "emoji")
	cmd.Dir = repo

	b, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}

	mrs := strings.Split(string(b), "\n")
	t.Log(mrs)
	require.Contains(t, mrs, "!6 test award emoji")
}
