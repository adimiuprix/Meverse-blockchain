// Package keydb came from buntdb that implements a low-level in-memory key/value store in pure Go.
// It persists to disk, is ACID compliant, and uses locking for multiple
// readers and a single writer. Bunt is ideal for projects that need
// a dependable database, and favor speed over data size.
package keydb2

// // SyncPolicy represents how often data is synced to disk.
// type SyncPolicy int

// // Config represents database configuration options. These
// // options are used to change various behaviors of the database.
// type Config struct {
// 	// SyncPolicy adjusts how often the data is synced to disk.
// 	// This value can be Never, EverySecond, or Always.
// 	// The default is EverySecond.
// 	SyncPolicy SyncPolicy

// 	// AutoShrinkPercentage is used by the background process to trigger
// 	// a shrink of the aof file when the size of the file is larger than the
// 	// percentage of the result of the previous shrunk file.
// 	// For example, if this value is 100, and the last shrink process
// 	// resulted in a 100mb file, then the new aof file must be 200mb before
// 	// a shrink is triggered.
// 	AutoShrinkPercentage int

// 	// AutoShrinkMinSize defines the minimum size of the aof file before
// 	// an automatic shrink can occur.
// 	AutoShrinkMinSize int

// 	// AutoShrinkDisabled turns off automatic background shrinking
// 	AutoShrinkDisabled bool

// 	// OnExpired is used to custom handle the deletion option when a key
// 	// has been expired.
// 	OnExpired func(keys [][]byte)

// 	// OnExpiredSync will be called inside the same transaction that is
// 	// performing the deletion of expired items. If OnExpired is present then
// 	// this callback will not be called. If this callback is present, then the
// 	// deletion of the timeed-out item is the explicit responsibility of this
// 	// callback.
// 	OnExpiredSync func(key []byte, value []byte, tx *Tx) error
// }
