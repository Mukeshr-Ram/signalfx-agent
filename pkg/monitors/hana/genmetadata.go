// Code generated by monitor-code-gen. DO NOT EDIT.

package hana

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "hana"

var groupSet = map[string]bool{}

const (
	sapHanaConnectionCount                       = "sap.hana.connection.count"
	sapHanaConnectionMessageReceivedCount        = "sap.hana.connection.message.received.count"
	sapHanaConnectionMessageReceivedSize         = "sap.hana.connection.message.received.size"
	sapHanaConnectionMessageSentCount            = "sap.hana.connection.message.sent.count"
	sapHanaConnectionMessageSentSize             = "sap.hana.connection.message.sent.size"
	sapHanaConnectionRecordAffected              = "sap.hana.connection.record.affected"
	sapHanaConnectionRecordFetched               = "sap.hana.connection.record.fetched"
	sapHanaDiskTotalSize                         = "sap.hana.disk.total_size"
	sapHanaDiskUsedSize                          = "sap.hana.disk.used_size"
	sapHanaHostCPUIdle                           = "sap.hana.host.cpu.idle"
	sapHanaHostCPUSystem                         = "sap.hana.host.cpu.system"
	sapHanaHostCPUUser                           = "sap.hana.host.cpu.user"
	sapHanaHostCPUWio                            = "sap.hana.host.cpu.wio"
	sapHanaHostFileOpen                          = "sap.hana.host.file.open"
	sapHanaHostMemoryAllocationLimit             = "sap.hana.host.memory.allocation_limit"
	sapHanaHostMemoryCode                        = "sap.hana.host.memory.code"
	sapHanaHostMemoryPhysicalFree                = "sap.hana.host.memory.physical.free"
	sapHanaHostMemoryPhysicalUsed                = "sap.hana.host.memory.physical.used"
	sapHanaHostMemoryShared                      = "sap.hana.host.memory.shared"
	sapHanaHostMemorySwapFree                    = "sap.hana.host.memory.swap.free"
	sapHanaHostMemorySwapUsed                    = "sap.hana.host.memory.swap.used"
	sapHanaHostMemoryTotalAllocated              = "sap.hana.host.memory.total_allocated"
	sapHanaHostMemoryTotalUsed                   = "sap.hana.host.memory.total_used"
	sapHanaIoAppendCount                         = "sap.hana.io.append.count"
	sapHanaIoReadAsyncCount                      = "sap.hana.io.read.async.count"
	sapHanaIoReadCount                           = "sap.hana.io.read.count"
	sapHanaIoReadFailed                          = "sap.hana.io.read.failed"
	sapHanaIoReadSize                            = "sap.hana.io.read.size"
	sapHanaIoReadTime                            = "sap.hana.io.read.time"
	sapHanaIoTotalTime                           = "sap.hana.io.total.time"
	sapHanaIoWriteAsyncCount                     = "sap.hana.io.write.async.count"
	sapHanaIoWriteCount                          = "sap.hana.io.write.count"
	sapHanaIoWriteFailed                         = "sap.hana.io.write.failed"
	sapHanaIoWriteSize                           = "sap.hana.io.write.size"
	sapHanaIoWriteTime                           = "sap.hana.io.write.time"
	sapHanaServiceComponentMemoryUsed            = "sap.hana.service.component.memory.used"
	sapHanaServiceCPUUtilization                 = "sap.hana.service.cpu.utilization"
	sapHanaServiceFileOpen                       = "sap.hana.service.file.open"
	sapHanaServiceMemoryAllocationLimit          = "sap.hana.service.memory.allocation_limit"
	sapHanaServiceMemoryAllocationLimitEffective = "sap.hana.service.memory.allocation_limit_effective"
	sapHanaServiceMemoryCode                     = "sap.hana.service.memory.code"
	sapHanaServiceMemoryHeapAllocated            = "sap.hana.service.memory.heap.allocated"
	sapHanaServiceMemoryHeapUsed                 = "sap.hana.service.memory.heap.used"
	sapHanaServiceMemoryLogical                  = "sap.hana.service.memory.logical"
	sapHanaServiceMemoryPhysical                 = "sap.hana.service.memory.physical"
	sapHanaServiceMemorySharedAllocated          = "sap.hana.service.memory.shared.allocated"
	sapHanaServiceMemorySharedUsed               = "sap.hana.service.memory.shared.used"
	sapHanaServiceMemoryStack                    = "sap.hana.service.memory.stack"
	sapHanaServiceMemoryTotalUsed                = "sap.hana.service.memory.total_used"
	sapHanaStatementActiveCount                  = "sap.hana.statement.active.count"
	sapHanaStatementActiveExecutionCount         = "sap.hana.statement.active.execution.count"
	sapHanaStatementActiveExecutionMemoryMax     = "sap.hana.statement.active.execution.memory.max"
	sapHanaStatementActiveExecutionMemoryMean    = "sap.hana.statement.active.execution.memory.mean"
	sapHanaStatementActiveExecutionMemorySum     = "sap.hana.statement.active.execution.memory.sum"
	sapHanaStatementActiveExecutionSum           = "sap.hana.statement.active.execution.sum"
	sapHanaStatementActiveExecutionTimeMax       = "sap.hana.statement.active.execution.time.max"
	sapHanaStatementActiveExecutionTimeMean      = "sap.hana.statement.active.execution.time.mean"
	sapHanaStatementActiveRecompileCount         = "sap.hana.statement.active.recompile.count"
	sapHanaStatementExpensiveCount               = "sap.hana.statement.expensive.count"
	sapHanaStatementExpensiveCPUTime             = "sap.hana.statement.expensive.cpu_time"
	sapHanaStatementExpensiveDuration            = "sap.hana.statement.expensive.duration"
	sapHanaStatementExpensiveErrors              = "sap.hana.statement.expensive.errors"
	sapHanaStatementExpensiveLockWaitDuration    = "sap.hana.statement.expensive.lock_wait_duration"
	sapHanaStatementExpensiveRecords             = "sap.hana.statement.expensive.records"
	sapHanaTableRecordCount                      = "sap.hana.table.record.count"
	sapHanaTableSize                             = "sap.hana.table.size"
)

var metricSet = map[string]monitors.MetricInfo{
	sapHanaConnectionCount:                       {Type: datapoint.Gauge},
	sapHanaConnectionMessageReceivedCount:        {Type: datapoint.Gauge},
	sapHanaConnectionMessageReceivedSize:         {Type: datapoint.Gauge},
	sapHanaConnectionMessageSentCount:            {Type: datapoint.Gauge},
	sapHanaConnectionMessageSentSize:             {Type: datapoint.Gauge},
	sapHanaConnectionRecordAffected:              {Type: datapoint.Gauge},
	sapHanaConnectionRecordFetched:               {Type: datapoint.Gauge},
	sapHanaDiskTotalSize:                         {Type: datapoint.Gauge},
	sapHanaDiskUsedSize:                          {Type: datapoint.Gauge},
	sapHanaHostCPUIdle:                           {Type: datapoint.Count},
	sapHanaHostCPUSystem:                         {Type: datapoint.Count},
	sapHanaHostCPUUser:                           {Type: datapoint.Count},
	sapHanaHostCPUWio:                            {Type: datapoint.Count},
	sapHanaHostFileOpen:                          {Type: datapoint.Gauge},
	sapHanaHostMemoryAllocationLimit:             {Type: datapoint.Gauge},
	sapHanaHostMemoryCode:                        {Type: datapoint.Gauge},
	sapHanaHostMemoryPhysicalFree:                {Type: datapoint.Gauge},
	sapHanaHostMemoryPhysicalUsed:                {Type: datapoint.Gauge},
	sapHanaHostMemoryShared:                      {Type: datapoint.Gauge},
	sapHanaHostMemorySwapFree:                    {Type: datapoint.Gauge},
	sapHanaHostMemorySwapUsed:                    {Type: datapoint.Gauge},
	sapHanaHostMemoryTotalAllocated:              {Type: datapoint.Gauge},
	sapHanaHostMemoryTotalUsed:                   {Type: datapoint.Gauge},
	sapHanaIoAppendCount:                         {Type: datapoint.Count},
	sapHanaIoReadAsyncCount:                      {Type: datapoint.Count},
	sapHanaIoReadCount:                           {Type: datapoint.Count},
	sapHanaIoReadFailed:                          {Type: datapoint.Count},
	sapHanaIoReadSize:                            {Type: datapoint.Count},
	sapHanaIoReadTime:                            {Type: datapoint.Count},
	sapHanaIoTotalTime:                           {Type: datapoint.Count},
	sapHanaIoWriteAsyncCount:                     {Type: datapoint.Count},
	sapHanaIoWriteCount:                          {Type: datapoint.Count},
	sapHanaIoWriteFailed:                         {Type: datapoint.Count},
	sapHanaIoWriteSize:                           {Type: datapoint.Count},
	sapHanaIoWriteTime:                           {Type: datapoint.Count},
	sapHanaServiceComponentMemoryUsed:            {Type: datapoint.Gauge},
	sapHanaServiceCPUUtilization:                 {Type: datapoint.Gauge},
	sapHanaServiceFileOpen:                       {Type: datapoint.Gauge},
	sapHanaServiceMemoryAllocationLimit:          {Type: datapoint.Gauge},
	sapHanaServiceMemoryAllocationLimitEffective: {Type: datapoint.Gauge},
	sapHanaServiceMemoryCode:                     {Type: datapoint.Gauge},
	sapHanaServiceMemoryHeapAllocated:            {Type: datapoint.Gauge},
	sapHanaServiceMemoryHeapUsed:                 {Type: datapoint.Gauge},
	sapHanaServiceMemoryLogical:                  {Type: datapoint.Gauge},
	sapHanaServiceMemoryPhysical:                 {Type: datapoint.Gauge},
	sapHanaServiceMemorySharedAllocated:          {Type: datapoint.Gauge},
	sapHanaServiceMemorySharedUsed:               {Type: datapoint.Gauge},
	sapHanaServiceMemoryStack:                    {Type: datapoint.Gauge},
	sapHanaServiceMemoryTotalUsed:                {Type: datapoint.Gauge},
	sapHanaStatementActiveCount:                  {Type: datapoint.Gauge},
	sapHanaStatementActiveExecutionCount:         {Type: datapoint.Gauge},
	sapHanaStatementActiveExecutionMemoryMax:     {Type: datapoint.Gauge},
	sapHanaStatementActiveExecutionMemoryMean:    {Type: datapoint.Gauge},
	sapHanaStatementActiveExecutionMemorySum:     {Type: datapoint.Gauge},
	sapHanaStatementActiveExecutionSum:           {Type: datapoint.Gauge},
	sapHanaStatementActiveExecutionTimeMax:       {Type: datapoint.Gauge},
	sapHanaStatementActiveExecutionTimeMean:      {Type: datapoint.Gauge},
	sapHanaStatementActiveRecompileCount:         {Type: datapoint.Gauge},
	sapHanaStatementExpensiveCount:               {Type: datapoint.Count},
	sapHanaStatementExpensiveCPUTime:             {Type: datapoint.Count},
	sapHanaStatementExpensiveDuration:            {Type: datapoint.Count},
	sapHanaStatementExpensiveErrors:              {Type: datapoint.Count},
	sapHanaStatementExpensiveLockWaitDuration:    {Type: datapoint.Count},
	sapHanaStatementExpensiveRecords:             {Type: datapoint.Count},
	sapHanaTableRecordCount:                      {Type: datapoint.Gauge},
	sapHanaTableSize:                             {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	sapHanaConnectionCount:                       true,
	sapHanaConnectionMessageReceivedCount:        true,
	sapHanaConnectionMessageReceivedSize:         true,
	sapHanaConnectionMessageSentCount:            true,
	sapHanaConnectionMessageSentSize:             true,
	sapHanaConnectionRecordAffected:              true,
	sapHanaConnectionRecordFetched:               true,
	sapHanaDiskTotalSize:                         true,
	sapHanaDiskUsedSize:                          true,
	sapHanaHostCPUIdle:                           true,
	sapHanaHostCPUSystem:                         true,
	sapHanaHostCPUUser:                           true,
	sapHanaHostCPUWio:                            true,
	sapHanaHostFileOpen:                          true,
	sapHanaHostMemoryAllocationLimit:             true,
	sapHanaHostMemoryCode:                        true,
	sapHanaHostMemoryPhysicalFree:                true,
	sapHanaHostMemoryPhysicalUsed:                true,
	sapHanaHostMemoryShared:                      true,
	sapHanaHostMemorySwapFree:                    true,
	sapHanaHostMemorySwapUsed:                    true,
	sapHanaHostMemoryTotalAllocated:              true,
	sapHanaHostMemoryTotalUsed:                   true,
	sapHanaIoAppendCount:                         true,
	sapHanaIoReadAsyncCount:                      true,
	sapHanaIoReadCount:                           true,
	sapHanaIoReadFailed:                          true,
	sapHanaIoReadSize:                            true,
	sapHanaIoReadTime:                            true,
	sapHanaIoTotalTime:                           true,
	sapHanaIoWriteAsyncCount:                     true,
	sapHanaIoWriteCount:                          true,
	sapHanaIoWriteFailed:                         true,
	sapHanaIoWriteSize:                           true,
	sapHanaIoWriteTime:                           true,
	sapHanaServiceComponentMemoryUsed:            true,
	sapHanaServiceCPUUtilization:                 true,
	sapHanaServiceFileOpen:                       true,
	sapHanaServiceMemoryAllocationLimit:          true,
	sapHanaServiceMemoryAllocationLimitEffective: true,
	sapHanaServiceMemoryCode:                     true,
	sapHanaServiceMemoryHeapAllocated:            true,
	sapHanaServiceMemoryHeapUsed:                 true,
	sapHanaServiceMemoryLogical:                  true,
	sapHanaServiceMemoryPhysical:                 true,
	sapHanaServiceMemorySharedAllocated:          true,
	sapHanaServiceMemorySharedUsed:               true,
	sapHanaServiceMemoryStack:                    true,
	sapHanaServiceMemoryTotalUsed:                true,
	sapHanaStatementActiveCount:                  true,
	sapHanaStatementActiveExecutionCount:         true,
	sapHanaStatementActiveExecutionMemoryMax:     true,
	sapHanaStatementActiveExecutionMemoryMean:    true,
	sapHanaStatementActiveExecutionMemorySum:     true,
	sapHanaStatementActiveExecutionSum:           true,
	sapHanaStatementActiveExecutionTimeMax:       true,
	sapHanaStatementActiveExecutionTimeMean:      true,
	sapHanaStatementActiveRecompileCount:         true,
	sapHanaStatementExpensiveCount:               true,
	sapHanaStatementExpensiveCPUTime:             true,
	sapHanaStatementExpensiveDuration:            true,
	sapHanaStatementExpensiveErrors:              true,
	sapHanaStatementExpensiveLockWaitDuration:    true,
	sapHanaStatementExpensiveRecords:             true,
	sapHanaTableRecordCount:                      true,
	sapHanaTableSize:                             true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "hana",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
