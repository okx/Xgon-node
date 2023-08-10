package synchronizer

import (
	"fmt"
	"sync"

	"github.com/0xPolygonHermez/zkevm-node/log"
)

type SendOrdererResultsToSynchronizer struct {
	mutex                   sync.Mutex
	channel                 chan getRollupInfoByBlockRangeResult
	lastBlockOnSynchronizer uint64
	pendingResults          []getRollupInfoByBlockRangeResult
}

func (s *SendOrdererResultsToSynchronizer) toStringBrief() string {
	return fmt.Sprintf("lastBlockSenedToSync[%v] len(pending_results)[%d]",
		s.lastBlockOnSynchronizer, len(s.pendingResults))
}

func NewSendResultsToSynchronizer(ch chan getRollupInfoByBlockRangeResult, lastBlockOnSynchronizer uint64) *SendOrdererResultsToSynchronizer {
	return &SendOrdererResultsToSynchronizer{channel: ch, lastBlockOnSynchronizer: lastBlockOnSynchronizer}
}

func (s *SendOrdererResultsToSynchronizer) addResultAndSendToConsumer(result *getRollupInfoByBlockRangeResult) {
	if result == nil {
		log.Fatal("Nil results, make no sense!")
		return
	}

	log.Info("Received: ", result.toStringBrief())
	s.mutex.Lock()
	defer s.mutex.Unlock()
	if result.blockRange.fromBlock < s.lastBlockOnSynchronizer {
		log.Fatalf("It's not possible to receive a old block [%s] range that been send to synchronizer status:[%s]",
			result.blockRange.toString(), s.toStringBrief())
		return
	}

	if !s._matchNextBlock(result) {
		log.Debugf("The range %s is not the next block to be send, 	adding to pending results status:%s",
			result.blockRange.toString(), s.toStringBrief())
	}
	s._addPendingResult(result)
	s._sendResultIfPossible()
}

// _sendResultIfPossible returns true is have send any result
func (s *SendOrdererResultsToSynchronizer) _sendResultIfPossible() bool {
	brToRemove := []blockRange{}
	send := false
	for _, result := range s.pendingResults {
		if s._matchNextBlock(&result) {
			send = true
			log.Info("Sending results to synchronizer:", result.toStringBrief())
			s.channel <- result
			s._setLastBlockOnSynchronizerCorrespondingLatBlockRangeSend(result.blockRange)
			brToRemove = append(brToRemove, result.blockRange)
			break
		}
	}
	for _, br := range brToRemove {
		s._removeBlockRange(br)
	}
	if send {
		// Try to send more results
		s._sendResultIfPossible()
	}
	return send
}

func (s *SendOrdererResultsToSynchronizer) _removeBlockRange(toRemove blockRange) {
	for i, result := range s.pendingResults {
		if result.blockRange == toRemove {
			s.pendingResults = append(s.pendingResults[:i], s.pendingResults[i+1:]...)
			break
		}
	}
}

func (s *SendOrdererResultsToSynchronizer) _setLastBlockOnSynchronizerCorrespondingLatBlockRangeSend(lastBlock blockRange) {
	newVaule := lastBlock.toBlock
	log.Info("Moving lastBlockSend from ", s.lastBlockOnSynchronizer, " to ", newVaule)
	s.lastBlockOnSynchronizer = newVaule
}

func (s *SendOrdererResultsToSynchronizer) _matchNextBlock(results *getRollupInfoByBlockRangeResult) bool {
	return results.blockRange.fromBlock == s.lastBlockOnSynchronizer+1
}

func (s *SendOrdererResultsToSynchronizer) _addPendingResult(results *getRollupInfoByBlockRangeResult) {
	s.pendingResults = append(s.pendingResults, *results)
}
