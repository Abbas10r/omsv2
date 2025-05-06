package main

import "context"

type store struct {
	//add mongodb inst
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(ctx context.Context) error {
	return nil
}
