package handler

import (
	"daily/reverse_tunnel/g"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"sync"
)

type progressSafeMap struct {
	sync.RWMutex
	M map[string]chan bool
}

var progressMap = progressSafeMap{M: map[string]chan bool{}}

func ProgressHandler(c *gin.Context) {
	req := g.Progress{}
	if err := c.BindJSON(&req); err != nil {
		log.Errorf(err.Error())
		c.JSON(400, g.Res{RetCode: 1, Desc: err.Error()})
		return
	}
	if req.Status == g.SUCCESS {
		progressMap.M[req.TaskID] <- true
	} else {
		progressMap.M[req.TaskID] <- false
	}
	c.JSON(200, g.Res{RetCode: 0, Desc: "OK"})
}

func (p *progressSafeMap) Get(key string) chan bool {
	p.RLock()
	defer p.RUnlock()
	return p.M[key]
}

func (p *progressSafeMap) Set(key string, value bool) {
	p.Lock()
	defer p.Unlock()
	if p.M[key] == nil {
		p.M[key] = make(chan bool)
	}
	p.M[key] <- value
}

func (p *progressSafeMap) Delete(key string) {
	p.Lock()
	defer p.Unlock()
	delete(p.M, key)
}

func (p *progressSafeMap) InitChan(key string) {
	p.Lock()
	defer p.Unlock()
	p.M[key] = make(chan bool)
}
