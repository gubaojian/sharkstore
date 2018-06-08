package server

import (
	"fmt"
	"model/pkg/taskpb"
	"sort"
	"time"
	"util/log"
)

const (
	defaultChangeLeaderTaskTimeout   = time.Second * time.Duration(30)
	defaultMaxChangeLeaderRetryTimes = 3
)

// ChangeLeaderTask  change leader task
type ChangeLeaderTask struct {
	*BaseTask
	fromNodeID uint64
	toNodeID   uint64 // zero means someone else exclude the fromNodeID

	retries           int
	roundRobinCounter int
}

// NewChangeLeaderTask new change leader task
func NewChangeLeaderTask(id uint64, rangeID uint64, from, to uint64) *ChangeLeaderTask {
	return &ChangeLeaderTask{
		BaseTask:   newBaseTask(id, rangeID, TaskTypeChangeLeader, defaultChangeLeaderTaskTimeout),
		fromNodeID: from,
		toNodeID:   to,
	}
}

func (t *ChangeLeaderTask) String() string {
	return fmt.Sprintf("{%s, \"from\": %d, \"to\": %d}", t.BaseTask.String(), t.fromNodeID, t.toNodeID)
}

// Step step
func (t *ChangeLeaderTask) Step(cluster *Cluster, r *Range) (over bool, task *taskpb.Task, err error) {
	// task is over
	if t.CheckOver() {
		return true, nil, nil
	}

	if r == nil {
		log.Warn("% invalid input range: <nil>", t.loggingID)
		return false, nil, fmt.Errorf("invalid step input: range is nil")
	}

	switch t.GetState() {
	case TaskStateStart:
		t.tryToChangeLeader(cluster, r)
		t.state = WaitLeaderChanged
		return false, nil, nil

	case WaitLeaderChanged:
		if t.isChanged(r) {
			log.Info("%s change finished. final leader: %d", t.loggingID, r.Leader.GetNodeId())
			t.state = TaskStateFinished
			return true, nil, nil
		}

		if t.retries >= defaultMaxChangeLeaderRetryTimes {
			log.Info("%s change canceled(max retry reached). final leader: %d", t.loggingID, r.Leader.GetNodeId())
			t.state = TaskStateCanceled
			return true, nil, nil
		}

		t.tryToChangeLeader(cluster, r)
		return false, nil, nil

	default:
		err = fmt.Errorf("unexpceted change leader task state: %s", t.state.String())
	}
	return
}

func (t *ChangeLeaderTask) isChanged(r *Range) bool {
	if t.toNodeID != 0 {
		return r.Leader.GetNodeId() == t.toNodeID
	}
	return r.Leader.GetNodeId() != t.fromNodeID
}

func (t *ChangeLeaderTask) selectChangeTo(r *Range) uint64 {
	if t.toNodeID != 0 {
		return t.toNodeID
	}

	var candidates []uint64
	peers := r.GetPeers()
	for _, p := range peers {
		if p.GetNodeId() != t.fromNodeID {
			candidates = append(candidates, p.GetNodeId())
		}
	}
	if len(candidates) == 0 {
		return 0
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i] < candidates[j] })
	changeTo := candidates[t.roundRobinCounter%len(candidates)]
	t.roundRobinCounter++
	return changeTo
}

func (t *ChangeLeaderTask) tryToChangeLeader(cluster *Cluster, r *Range) {
	t.retries++

	changeTo := t.selectChangeTo(r)
	if changeTo == 0 {
		log.Warn("%s invalid target node(0), current peers: %v", t.loggingID, r.GetPeers())
		return
	}

	node := cluster.FindNodeById(changeTo)
	if node == nil {
		log.Warn("%s could not find target node(%d)", t.loggingID, changeTo)
		return
	}

	log.Info("%s try to change leader to %d, current: %d", t.loggingID, changeTo, r.Leader.GetNodeId())

	//TODO:可能对堵塞时间比较长
	cluster.cli.TransferLeader(node.GetServerAddr(), r.GetId())
}
