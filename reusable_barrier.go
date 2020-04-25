package lilsem

// ReusableBarrier is a barrier that is reusable as in a loop
//
// note: personally unsure how the pattern described in the book
// adds any substantial value to the patterns already tackled.
// This one can easily be implemented with a combination of the other
// patterns. Whether there are one or two critical section/s interleaved
// between two or three rendezvous depends on specific usecase and so
// I will be skipping the implementation for this one.
//
// 	 (threads)	|  |  |  |     <--.
//				rendezvous         \
// 				    |               |  `threads` loop over the whole pattern
// 			critical section*       |  so even the second rendezvous can
// 					|               |  be ommitted
// 				rendezvous*        /
// 				|  |  |  |     --/
//
// type ReusableBarrier struct {
// 	sem *semaphore.Weighted
// 	mtx *semaphore.Weighted
// }
