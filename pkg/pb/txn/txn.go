// Copyright 2022 MatrixOrigin.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package txn

import (
	"bytes"
	"encoding/hex"
	"fmt"
)

const (
	// SkipResponseFlag skip response.
	SkipResponseFlag uint32 = 1
)

// NewTxnRequest create TxnRequest by CNOpRequest
func NewTxnRequest(request *CNOpRequest) TxnRequest {
	return TxnRequest{CNRequest: request}
}

// GetCNOpResponse returns the CNOpResponse from TxnResponse
func GetCNOpResponse(response TxnResponse) CNOpResponse {
	return *response.CNOpResponse
}

// HasFlag returns true if has the spec flag
func (m TxnResponse) HasFlag(flag uint32) bool {
	return m.Flag&flag > 0
}

// DebugString returns debug string
func (m TxnRequest) DebugString() string {
	var buffer bytes.Buffer

	buffer.WriteString("txn-meta: <")
	buffer.WriteString(m.Txn.DebugString())
	buffer.WriteString(">, ")

	buffer.WriteString("method: ")
	buffer.WriteString(m.Method.String())
	buffer.WriteString(", ")

	buffer.WriteString("flag: ")
	buffer.WriteString(fmt.Sprintf("%d", m.Flag))
	buffer.WriteString(", ")

	if m.CNRequest != nil {
		buffer.WriteString("cn-request: <")
		buffer.WriteString(m.CNRequest.DebugString())
		buffer.WriteString(">, ")
	}

	return buffer.String()
}

// DebugString returns debug string
func (m TxnError) DebugString() string {
	// TODO: impl
	return ""
}

// DebugString returns debug string
func (m TxnResponse) DebugString() string {
	var buffer bytes.Buffer

	buffer.WriteString("txn-meta: <")
	buffer.WriteString(m.Txn.DebugString())
	buffer.WriteString(">, ")

	buffer.WriteString("method: ")
	buffer.WriteString(m.Method.String())
	buffer.WriteString(", ")

	buffer.WriteString("flag: ")
	buffer.WriteString(fmt.Sprintf("%d", m.Flag))
	buffer.WriteString(", ")

	if m.TxnError != nil {
		buffer.WriteString("error: <")
		buffer.WriteString(m.TxnError.DebugString())
		buffer.WriteString(">, ")
	}

	if m.CNOpResponse != nil {
		buffer.WriteString("cn-response: <")
		buffer.WriteString(m.CNOpResponse.DebugString())
		buffer.WriteString(">, ")
	}

	return buffer.String()
}

// DebugString returns debug string
func (m CNOpRequest) DebugString() string {
	var buffer bytes.Buffer

	buffer.WriteString("op: ")
	buffer.WriteString(fmt.Sprintf("%d", m.OpCode))
	buffer.WriteString(", ")

	buffer.WriteString("payload: ")
	buffer.WriteString(fmt.Sprintf("%d bytes", m.Payload))
	buffer.WriteString(", ")

	buffer.WriteString("dn: <")
	buffer.WriteString(m.Target.DebugString())
	buffer.WriteString(">")
	return buffer.String()
}

// CNOpResponse returns debug string
func (m CNOpResponse) DebugString() string {
	var buffer bytes.Buffer

	buffer.WriteString("payload: ")
	buffer.WriteString(fmt.Sprintf("%d bytes", m.Payload))
	return buffer.String()
}

// DebugString returns debug string
func (m TxnMeta) DebugString() string {
	var buffer bytes.Buffer

	buffer.WriteString("txn-id: ")
	buffer.WriteString(hex.EncodeToString(m.ID))
	buffer.WriteString(", ")

	buffer.WriteString("status: ")
	buffer.WriteString(m.Status.String())
	buffer.WriteString(", ")

	buffer.WriteString("snapshot-ts: ")
	buffer.WriteString(m.SnapshotTS.String())
	buffer.WriteString(", ")

	if m.PreparedTS != nil {
		buffer.WriteString("prepared-ts: ")
		buffer.WriteString(m.PreparedTS.String())
		buffer.WriteString(", ")
	}

	if m.CommitTS != nil {
		buffer.WriteString("commit-ts: ")
		buffer.WriteString(m.CommitTS.String())
		buffer.WriteString(", ")
	}
	return buffer.String()
}