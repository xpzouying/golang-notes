package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoLockQueue_EmptyQueue(t *testing.T) {
	q := NewTwoLockQueue()

	assert.Nil(t, q.Head.Next)
	assert.Nil(t, q.Tail.Next)
}

func TestTwoLockQueue_Enqueue(t *testing.T) {

	q := NewTwoLockQueue()

	q.Enqueue(1)

	v, err := q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 1, v)
}
