package observer_article

import (
	"github.com/gin-gonic/gin"
)

type ObserverArticleNotifier struct {
	AfterArticleLookList        []AfterArticleLookListener
	AfterArticleDiggIncrList    []AfterArticleDiggIncrListener
	AfterArticleDiggDecList     []AfterArticleDiggDecListener
	AfterArticleExamineList     []AfterArticleExamineListener
	AfterArticleCollectIncrList []AfterArticleCollectIncrListener
	AfterArticleCollectDecList  []AfterArticleCollectDecListener
	AfterArticleCommentIncrList []AfterArticleCommentIncrListener
	AfterArticleCommentDecList  []AfterArticleCommentDecListener
}

func NewArticleNotifier() *ObserverArticleNotifier {
	return &ObserverArticleNotifier{
		AfterArticleLookList:        make([]AfterArticleLookListener, 0),
		AfterArticleDiggIncrList:    make([]AfterArticleDiggIncrListener, 0),
		AfterArticleDiggDecList:     make([]AfterArticleDiggDecListener, 0),
		AfterArticleExamineList:     make([]AfterArticleExamineListener, 0),
		AfterArticleCollectIncrList: make([]AfterArticleCollectIncrListener, 0),
		AfterArticleCollectDecList:  make([]AfterArticleCollectDecListener, 0),
		AfterArticleCommentIncrList: make([]AfterArticleCommentIncrListener, 0),
		AfterArticleCommentDecList:  make([]AfterArticleCommentDecListener, 0),
	}
}

func (a *ObserverArticleNotifier) AddArticleLookListener(listeners ...AfterArticleLookListener) {
	for _, listener := range listeners {
		a.AfterArticleLookList = append(a.AfterArticleLookList, listener)
	}
}
func (a *ObserverArticleNotifier) AddArticleDiggIncrListener(listeners ...AfterArticleDiggIncrListener) {
	for _, listener := range listeners {
		a.AfterArticleDiggIncrList = append(a.AfterArticleDiggIncrList, listener)
	}
}
func (a *ObserverArticleNotifier) AddArticleDiggDecListener(listeners ...AfterArticleDiggDecListener) {
	for _, listener := range listeners {
		a.AfterArticleDiggDecList = append(a.AfterArticleDiggDecList, listener)
	}
}
func (a *ObserverArticleNotifier) AddArticleExamineListener(listeners ...AfterArticleExamineListener) {
	for _, listener := range listeners {
		a.AfterArticleExamineList = append(a.AfterArticleExamineList, listener)
	}
}
func (a *ObserverArticleNotifier) AddArticleCollectIncrListener(listeners ...AfterArticleCollectIncrListener) {
	for _, listener := range listeners {
		a.AfterArticleCollectIncrList = append(a.AfterArticleCollectIncrList, listener)
	}
}
func (a *ObserverArticleNotifier) AddArticleCollectDecListener(listeners ...AfterArticleCollectDecListener) {
	for _, listener := range listeners {
		a.AfterArticleCollectDecList = append(a.AfterArticleCollectDecList, listener)
	}
}
func (a *ObserverArticleNotifier) AddArticleCommentIncrListener(listeners ...AfterArticleCommentIncrListener) {
	for _, listener := range listeners {
		a.AfterArticleCommentIncrList = append(a.AfterArticleCommentIncrList, listener)
	}
}
func (a *ObserverArticleNotifier) AddArticleCommentDecListener(listeners ...AfterArticleCommentDecListener) {
	for _, listener := range listeners {
		a.AfterArticleCommentDecList = append(a.AfterArticleCommentDecList, listener)
	}
}

func (a *ObserverArticleNotifier) RemoveArticleLookListener(listeners ...AfterArticleLookListener) {
	for index, l := range a.AfterArticleLookList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterArticleLookList = append(a.AfterArticleLookList[:index], a.AfterArticleLookList[index+1:]...)
			}
		}
	}
}
func (a *ObserverArticleNotifier) RemoveArticleDiggIncrListener(listeners ...AfterArticleDiggIncrListener) {
	for index, l := range a.AfterArticleDiggIncrList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterArticleDiggIncrList = append(a.AfterArticleDiggIncrList[:index], a.AfterArticleDiggIncrList[index+1:]...)
			}
		}
	}
}
func (a *ObserverArticleNotifier) RemoveArticleDiggDecListener(listeners ...AfterArticleDiggDecListener) {
	for index, l := range a.AfterArticleDiggDecList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterArticleDiggDecList = append(a.AfterArticleDiggDecList[:index], a.AfterArticleDiggDecList[index+1:]...)
			}
		}
	}
}
func (a *ObserverArticleNotifier) RemoveArticleExamineListener(listeners ...AfterArticleExamineListener) {
	for index, l := range a.AfterArticleExamineList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterArticleExamineList = append(a.AfterArticleExamineList[:index], a.AfterArticleExamineList[index+1:]...)
			}
		}
	}
}
func (a *ObserverArticleNotifier) RemoveArticleCollectIncrListener(listeners ...AfterArticleCollectIncrListener) {
	for index, l := range a.AfterArticleCollectIncrList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterArticleCollectIncrList = append(a.AfterArticleCollectIncrList[:index], a.AfterArticleCollectIncrList[index+1:]...)
			}
		}
	}
}
func (a *ObserverArticleNotifier) RemoveArticleCollectDecListener(listeners ...AfterArticleCollectDecListener) {
	for index, l := range a.AfterArticleCollectDecList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterArticleCollectDecList = append(a.AfterArticleCollectDecList[:index], a.AfterArticleCollectDecList[index+1:]...)
			}
		}
	}
}
func (a *ObserverArticleNotifier) RemoveArticleCommentIncrListener(listeners ...AfterArticleCommentIncrListener) {
	for index, l := range a.AfterArticleCommentIncrList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterArticleCommentIncrList = append(a.AfterArticleCommentIncrList[:index], a.AfterArticleCommentIncrList[index+1:]...)
			}
		}
	}
}
func (a *ObserverArticleNotifier) RemoveArticleCommentDecListener(listeners ...AfterArticleCommentDecListener) {
	for index, l := range a.AfterArticleCommentDecList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterArticleCommentDecList = append(a.AfterArticleCommentDecList[:index], a.AfterArticleCommentDecList[index+1:]...)
			}
		}
	}
}

func (a *ObserverArticleNotifier) AfterArticleLookNotify(c *gin.Context, articleID uint) {
	for _, listener := range a.AfterArticleLookList {
		listener.AfterArticleLook(c, articleID)
	}
}
func (a *ObserverArticleNotifier) AfterArticleDiggIncrNotify(articleID uint) {
	for _, listener := range a.AfterArticleDiggIncrList {
		listener.AfterArticleDiggIncr(articleID)
	}
}
func (a *ObserverArticleNotifier) AfterArticleDiggDecNotify(articleID uint) {
	for _, listener := range a.AfterArticleDiggDecList {
		listener.AfterArticleDiggDec(articleID)
	}
}
func (a *ObserverArticleNotifier) AfterArticleExamineNotify(articleID uint, status int8) {
	for _, listener := range a.AfterArticleExamineList {
		listener.AfterArticleExamine(articleID, status)
	}
}
func (a *ObserverArticleNotifier) AfterArticleCollectIncrNotify(articleID uint) {
	for _, listener := range a.AfterArticleCollectIncrList {
		listener.AfterArticleCollectIncr(articleID, 2)
	}
}
func (a *ObserverArticleNotifier) AfterArticleCollectDecNotify(articleID uint) {
	for _, listener := range a.AfterArticleCollectDecList {
		listener.AfterArticleCollectDec(articleID, 2)
	}
}
func (a *ObserverArticleNotifier) AfterArticleCommentIncrNotify(articleID uint) {
	for _, listener := range a.AfterArticleCommentIncrList {
		listener.AfterArticleCommentIncr(articleID)
	}
}
func (a *ObserverArticleNotifier) AfterArticleCommentDecNotify(articleID uint, n int) {
	for _, listener := range a.AfterArticleCommentDecList {
		listener.AfterArticleCommentDec(articleID, 2)
	}
}
