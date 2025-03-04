package core_redis

//func InitRedisService() {
//	var articleList []text_service.ParticipleArticleModel
//	global.DB.Model(models.ArticleModel{}).Find(&articleList)
//	var textList []text_service.ParticipleTextModel
//	global.DB.Model(&models.TextModel{}).Find(&textList)
//	InitRedisIndex(articleList, textList)
//}

//func InitRedisIndex(articleList []text_service.ParticipleArticleModel, textList []text_service.ParticipleTextModel) {
//	logrus.Info("标签聚合加载中...")
//	initRedisTagAgg(articleList)
//	logrus.Info("标签聚合加载完成")
//	logrus.Info("文章搜索分词加载中...")
//	initRedisArticleSearch(articleList)
//	logrus.Info("文章搜索分词加载完成")
//	logrus.Info("文章搜索排序加载中...")
//	initRedisArticleSort(articleList...)
//	logrus.Info("文章搜索排序加载完成")
//	logrus.Info("全文搜索分词加载中...")
//	initRedisTextSearch(textList)
//	logrus.Info("全文搜索分词加载完成")
//}
