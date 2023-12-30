package cinnamon

import (
	"time"
)

// Total of 768 different priorities, initialized to middle
// 0-5 tiers where 0 > 5 and 0-127 cohorts where 0 > 127
var TIER_COHORT_THRESHOLD = 767

// Number of workers actively popping from priority queue
// and forwarding requests
var NUM_WORKERS = 2

// Expiration timeout for requests in priority queue
var MAX_AGE = 3 * time.Second

// IN corresponds to requests being added to priority queue
// OUT corresponds to requests leaving the priority queue to be
// forwarded to respective services
var IN float64 = 0.0
var OUT float64 = 0.0

// View last 1000 request priorities needed by PID controller
// to set the new TIER_COHORT_THRESHOLD
var MAX_HISTORY = 1000

// Inflight Request Limit to Forward Requests
var INFLIGHT_LIMIT float64 = 20000.0
var CURR_INFLIGHT float64 = 0.0

var CIRCULAR_QUEUE_LENGTH = 1000
