package mocktracer

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// MockSpanContext is an opentracing.SpanContext implementation.
//
// It is entirely unsuitable for production use, but appropriate for tests
// that want to verify tracing behavior in other frameworks/applications.
//
// By default all spans have Sampled=true flag, unless {"sampling.priority": 0}
// tag is set.
type MockSpanContext struct {
	TraceID int
	SpanID  int
	Sampled bool
	Baggage map[string]string
}

var mockIDSource = uint32(42)

func nextMockID() int {
	atomic.AddUint32(&mockIDSource, 1)
	return int(atomic.LoadUint32(&mockIDSource))
}

// ForeachBaggageItem belongs to the SpanContext interface
func (c MockSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {
	for k, v := range c.Baggage {
		if !handler(k, v) {
			break
		}
	}
}

// WithBaggageItem creates a new context with an extra baggage item.
func (c MockSpanContext) WithBaggageItem(key, value string) MockSpanContext {
	var newBaggage map[string]string
	if c.Baggage == nil {
		newBaggage = map[string]string{key: value}
	} else {
		newBaggage = make(map[string]string, len(c.Baggage)+1)
		for k, v := range c.Baggage {
			newBaggage[k] = v
		}
		newBaggage[key] = value
	}
	// Use positional parameters so the compiler will help catch new fields.
	return MockSpanContext{c.TraceID, c.SpanID, c.Sampled, newBaggage}
}

// MockSpan is an opentracing.Span implementation that exports its internal
// state for testing purposes.
type MockSpan struct {
	sync.RWMutex

	ParentID int

	OperationName string
	StartTime     time.Time
	FinishTime    time.Time

	// All of the below are protected by the embedded RWMutex.
	SpanContext MockSpanContext
	tags        map[string]interface{}
	logs        []opentracing.LogData
	tracer      *MockTracer
}

func newMockSpan(t *MockTracer, name string, opts opentracing.StartSpanOptions) *MockSpan {
	tags := opts.Tags
	if tags == nil {
		tags = map[string]interface{}{}
	}
	traceID := nextMockID()
	parentID := int(0)
	var baggage map[string]string
	sampled := true
	if len(opts.References) > 0 {
		traceID = opts.References[0].ReferencedContext.(MockSpanContext).TraceID
		parentID = opts.References[0].ReferencedContext.(MockSpanContext).SpanID
		sampled = opts.References[0].ReferencedContext.(MockSpanContext).Sampled
		baggage = opts.References[0].ReferencedContext.(MockSpanContext).Baggage
	}
	spanContext := MockSpanContext{traceID, nextMockID(), sampled, baggage}
	startTime := opts.StartTime
	if startTime.IsZero() {
		startTime = time.Now()
	}
	return &MockSpan{
		ParentID:      parentID,
		OperationName: name,
		StartTime:     startTime,
		tags:          tags,
		logs:          []opentracing.LogData{},
		SpanContext:   spanContext,

		tracer: t,
	}
}

// Tags returns a copy of tags accumulated by the span so far
func (s *MockSpan) Tags() map[string]interface{} {
	s.RLock()
	defer s.RUnlock()
	tags := make(map[string]interface{})
	for k, v := range s.tags {
		tags[k] = v
	}
	return tags
}

// Tag returns a single tag
func (s *MockSpan) Tag(k string) interface{} {
	s.RLock()
	defer s.RUnlock()
	return s.tags[k]
}

// Logs returns a copy of logs accumulated in the span so far
func (s *MockSpan) Logs() []opentracing.LogData {
	s.RLock()
	defer s.RUnlock()
	logs := make([]opentracing.LogData, len(s.logs))
	copy(logs, s.logs)
	return logs
}

// Context belongs to the Span interface
func (s *MockSpan) Context() opentracing.SpanContext {
	return s.SpanContext
}

// SetTag belongs to the Span interface
func (s *MockSpan) SetTag(key string, value interface{}) opentracing.Span {
	s.Lock()
	defer s.Unlock()
	if key == string(ext.SamplingPriority) {
		if v, ok := value.(uint16); ok {
			s.SpanContext.Sampled = v > 0
			return s
		}
		if v, ok := value.(int); ok {
			s.SpanContext.Sampled = v > 0
			return s
		}
	}
	s.tags[key] = value
	return s
}

// SetBaggageItem belongs to the Span interface
func (s *MockSpan) SetBaggageItem(key, val string) opentracing.Span {
	s.Lock()
	defer s.Unlock()
	s.SpanContext = s.SpanContext.WithBaggageItem(key, val)
	return s
}

// BaggageItem belongs to the Span interface
func (s *MockSpan) BaggageItem(key string) string {
	s.RLock()
	defer s.RUnlock()
	return s.SpanContext.Baggage[key]
}

// Finish belongs to the Span interface
func (s *MockSpan) Finish() {
	s.Lock()
	s.FinishTime = time.Now()
	s.Unlock()
	s.tracer.recordSpan(s)
}

// FinishWithOptions belongs to the Span interface
func (s *MockSpan) FinishWithOptions(opts opentracing.FinishOptions) {
	s.Lock()
	s.FinishTime = opts.FinishTime
	s.logs = append(s.logs, opts.BulkLogData...)
	s.Unlock()
	s.tracer.recordSpan(s)
}

// String allows printing span for debugging
func (s *MockSpan) String() string {
	return fmt.Sprintf(
		"traceId=%d, spanId=%d, parentId=%d, sampled=%t, name=%s",
		s.SpanContext.TraceID, s.SpanContext.SpanID, s.ParentID,
		s.SpanContext.Sampled, s.OperationName)
}

// LogEvent belongs to the Span interface
func (s *MockSpan) LogEvent(event string) {
	s.Log(opentracing.LogData{
		Event: event,
	})
}

// LogEventWithPayload belongs to the Span interface
func (s *MockSpan) LogEventWithPayload(event string, payload interface{}) {
	s.Log(opentracing.LogData{
		Event:   event,
		Payload: payload,
	})
}

// Log belongs to the Span interface
func (s *MockSpan) Log(data opentracing.LogData) {
	s.Lock()
	defer s.Unlock()
	s.logs = append(s.logs, data)
}

// SetOperationName belongs to the Span interface
func (s *MockSpan) SetOperationName(operationName string) opentracing.Span {
	s.Lock()
	defer s.Unlock()
	s.OperationName = operationName
	return s
}

// Tracer belongs to the Span interface
func (s *MockSpan) Tracer() opentracing.Tracer {
	return s.tracer
}
