package observer_comment

type ObserverCommentNotifier struct {
	AfterCommentDiggIncrList []AfterCommentDiggIncrListener
	AfterCommentDiggDecList  []AfterCommentDiggDecListener
	AfterCommentSubIncrList  []AfterCommentSubIncrListener
	AfterCommentSubDecList   []AfterCommentSubDecListener
}

func NewCommentNotifier() *ObserverCommentNotifier {
	return &ObserverCommentNotifier{
		AfterCommentDiggIncrList: make([]AfterCommentDiggIncrListener, 0),
		AfterCommentDiggDecList:  make([]AfterCommentDiggDecListener, 0),
		AfterCommentSubIncrList:  make([]AfterCommentSubIncrListener, 0),
		AfterCommentSubDecList:   make([]AfterCommentSubDecListener, 0),
	}
}

func (a *ObserverCommentNotifier) AddCommentDiggIncrListener(listeners ...AfterCommentDiggIncrListener) {
	for _, listener := range listeners {
		a.AfterCommentDiggIncrList = append(a.AfterCommentDiggIncrList, listener)
	}
}
func (a *ObserverCommentNotifier) AddCommentDiggDecListener(listeners ...AfterCommentDiggDecListener) {
	for _, listener := range listeners {
		a.AfterCommentDiggDecList = append(a.AfterCommentDiggDecList, listener)
	}
}

func (a *ObserverCommentNotifier) AddCommentSubIncrListener(listeners ...AfterCommentSubIncrListener) {
	for _, listener := range listeners {
		a.AfterCommentSubIncrList = append(a.AfterCommentSubIncrList, listener)
	}
}
func (a *ObserverCommentNotifier) AddCommentSubDecListener(listeners ...AfterCommentSubDecListener) {
	for _, listener := range listeners {
		a.AfterCommentSubDecList = append(a.AfterCommentSubDecList, listener)
	}
}

func (a *ObserverCommentNotifier) RemoveCommentDiggIncrListener(listeners ...AfterCommentDiggIncrListener) {
	for index, l := range a.AfterCommentDiggIncrList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterCommentDiggIncrList = append(a.AfterCommentDiggIncrList[:index], a.AfterCommentDiggIncrList[index+1:]...)
			}
		}
	}
}
func (a *ObserverCommentNotifier) RemoveCommentDiggDecListener(listeners ...AfterCommentDiggDecListener) {
	for index, l := range a.AfterCommentDiggDecList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterCommentDiggDecList = append(a.AfterCommentDiggDecList[:index], a.AfterCommentDiggDecList[index+1:]...)
			}
		}
	}
}

func (a *ObserverCommentNotifier) RemoveCommentSubIncrListener(listeners ...AfterCommentSubIncrListener) {
	for index, l := range a.AfterCommentSubIncrList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterCommentSubIncrList = append(a.AfterCommentSubIncrList[:index], a.AfterCommentSubIncrList[index+1:]...)
			}
		}
	}
}
func (a *ObserverCommentNotifier) RemoveCommentSubDecListener(listeners ...AfterCommentSubDecListener) {
	for index, l := range a.AfterCommentSubDecList {
		for _, listener := range listeners {
			if l == listener {
				a.AfterCommentSubDecList = append(a.AfterCommentSubDecList[:index], a.AfterCommentSubDecList[index+1:]...)
			}
		}
	}
}

func (a *ObserverCommentNotifier) AfterCommentDiggIncrNotify(commentID uint) {
	for _, listener := range a.AfterCommentDiggIncrList {
		listener.AfterCommentDiggIncr(commentID)
	}
}
func (a *ObserverCommentNotifier) AfterCommentDiggDecNotify(commentID uint) {
	for _, listener := range a.AfterCommentDiggDecList {
		listener.AfterCommentDiggDec(commentID)
	}
}

func (a *ObserverCommentNotifier) AfterCommentSubIncrNotify(commentID uint) {
	for _, listener := range a.AfterCommentSubIncrList {
		listener.AfterCommentSubIncr(commentID)
	}
}
func (a *ObserverCommentNotifier) AfterCommentSubDecNotify(commentID uint, n int) {
	for _, listener := range a.AfterCommentSubDecList {
		listener.AfterCommentSubDec(commentID, n)
	}
}
