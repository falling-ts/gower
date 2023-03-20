package translate

import (
	"errors"
	"gower/services"
	"regexp"
)

type Service struct {
	services.TransAll
}

var config services.Config

func Mount(all services.TransAll) *Service {
	s := new(Service)
	s.TransAll = all

	return s
}

func (s *Service) Init(args ...any) {
	config = args[0].(services.Config)
	s.initDBError()
}

// DBError 翻译数据库错误
func (s *Service) DBError(err error) error {
	if err == nil {
		return err
	}

	msg := err.Error()
	dbError, ok := s.TransAll["DBError"]
	if !ok {
		return err
	}
	dbErrorMap, ok := dbError.(services.TransMap)
	if !ok {
		return err
	}

	for k, v := range dbErrorMap {
		r := regexp.MustCompile(k)
		trans := r.ReplaceAllString(msg, v)
		if trans != msg && trans != "" {
			return errors.New(trans)
		}
	}

	return err
}

func (s *Service) initDBError() {
	dbErrorCateKey := config.Get("db.driver", "mysql").(string)
	dbError, ok := s.TransAll["DBError"]
	if !ok {
		return
	}
	dbErrorCate, ok := dbError.(services.TransCategory)
	if !ok {
		return
	}

	if dbErrorMap, ok := dbErrorCate[dbErrorCateKey]; ok {
		s.TransAll["DBError"] = dbErrorMap
	}
}
