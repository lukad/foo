package store

import (
	"github.com/lukad/helix/mail"
	"sync"
)

type Mail struct {
	mail.Mail
	Id int `json:"id"`
}

type Store interface {
	All() []Mail
	Get(id int) (Mail, bool)
	Insert(mail.Mail) int
	Del(id int) bool
	Count() int
}

type store struct {
	mails map[int]mail.Mail
	id    int
	mutex sync.RWMutex
}

func New() Store {
	return &store{
		mails: make(map[int]mail.Mail),
	}
}

func (s *store) All() (mails []Mail) {
	s.mutex.RLock()
	for id, m := range s.mails {
		mails = append(mails, Mail{m, id})
	}
	s.mutex.RUnlock()
	return
}

func (s *store) Get(id int) (m Mail, ok bool) {
	s.mutex.RLock()
	var ml mail.Mail
	if ml, ok = s.mails[id]; !ok {
		return
	}
	s.mutex.RUnlock()
	return Mail{ml, id}, true
}

func (s *store) Insert(m mail.Mail) int {
	s.mutex.Lock()
	currentId := s.id
	s.mails[currentId] = m
	s.id = s.id + 1
	s.mutex.Unlock()
	return currentId
}

func (s *store) Del(id int) (ok bool) {
	s.mutex.Lock()
	if _, ok := s.mails[id]; ok != true {
		return false
	}
	delete(s.mails, id)
	s.mutex.Unlock()
	return true
}

func (s *store) Count() int {
	return len(s.mails)
}
