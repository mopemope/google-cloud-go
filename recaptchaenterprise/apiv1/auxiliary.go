// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package recaptchaenterprise

import (
	recaptchaenterprisepb "cloud.google.com/go/recaptchaenterprise/v2/apiv1/recaptchaenterprisepb"
	"google.golang.org/api/iterator"
)

// FirewallPolicyIterator manages a stream of *recaptchaenterprisepb.FirewallPolicy.
type FirewallPolicyIterator struct {
	items    []*recaptchaenterprisepb.FirewallPolicy
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*recaptchaenterprisepb.FirewallPolicy, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *FirewallPolicyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *FirewallPolicyIterator) Next() (*recaptchaenterprisepb.FirewallPolicy, error) {
	var item *recaptchaenterprisepb.FirewallPolicy
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *FirewallPolicyIterator) bufLen() int {
	return len(it.items)
}

func (it *FirewallPolicyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// KeyIterator manages a stream of *recaptchaenterprisepb.Key.
type KeyIterator struct {
	items    []*recaptchaenterprisepb.Key
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*recaptchaenterprisepb.Key, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *KeyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *KeyIterator) Next() (*recaptchaenterprisepb.Key, error) {
	var item *recaptchaenterprisepb.Key
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *KeyIterator) bufLen() int {
	return len(it.items)
}

func (it *KeyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RelatedAccountGroupIterator manages a stream of *recaptchaenterprisepb.RelatedAccountGroup.
type RelatedAccountGroupIterator struct {
	items    []*recaptchaenterprisepb.RelatedAccountGroup
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*recaptchaenterprisepb.RelatedAccountGroup, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *RelatedAccountGroupIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RelatedAccountGroupIterator) Next() (*recaptchaenterprisepb.RelatedAccountGroup, error) {
	var item *recaptchaenterprisepb.RelatedAccountGroup
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RelatedAccountGroupIterator) bufLen() int {
	return len(it.items)
}

func (it *RelatedAccountGroupIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}

// RelatedAccountGroupMembershipIterator manages a stream of *recaptchaenterprisepb.RelatedAccountGroupMembership.
type RelatedAccountGroupMembershipIterator struct {
	items    []*recaptchaenterprisepb.RelatedAccountGroupMembership
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*recaptchaenterprisepb.RelatedAccountGroupMembership, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *RelatedAccountGroupMembershipIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *RelatedAccountGroupMembershipIterator) Next() (*recaptchaenterprisepb.RelatedAccountGroupMembership, error) {
	var item *recaptchaenterprisepb.RelatedAccountGroupMembership
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *RelatedAccountGroupMembershipIterator) bufLen() int {
	return len(it.items)
}

func (it *RelatedAccountGroupMembershipIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
