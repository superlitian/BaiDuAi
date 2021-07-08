/* Copyright 2021 CoderLee Inc. All Rights Reserved. */
/* constants.go 常量 */
/* modification history
-----------------------
2021/7/1, by coderlee, 创建
*/
/*
DESCRIPTION
	包含BaiDu_Ai包下所有的常量
*/

package BaiDu_Ai

//Constants

// 接口地址

// auth
const (
	authurl = "https://aip.baidubce.com/oauth/2.0/token"
)

// face
const (
	// detectUrl 人脸检测url
	detectUrl = "https://aip.baidubce.com/rest/2.0/face/v3/detect"
	// matchUrl 人脸对比url
	matchUrl = "https://aip.baidubce.com/rest/2.0/face/v3/match"
	// searchUrl 人脸搜索url
	searchUrl = "https://aip.baidubce.com/rest/2.0/face/v3/search"
	// multiSearchUrl 人脸搜索 M:N url
	multiSearchUrl = "https://aip.baidubce.com/rest/2.0/face/v3/multi-search"
	// addUserUrl 人脸注册url
	addUserUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add"
	// updateUserUrl 人脸更新url
	updateUserUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/update"
	// faceDeleteUrl 人脸删除url
	faceDeleteUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/face/delete"
	// getUserUrl 用户信息查询url
	getUserUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/get"
	// faceGetlistUrl 获取用户人脸列表url
	faceGetlistUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/face/getlist"
	// getGroupUsersUrl 获取用户列表url
	getGroupUsersUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/getusers"
	// userCopyUrl 复制用户url
	userCopyUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/copy"
	// deleteUserUrl 删除用户url
	deleteUserUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/delete"
	// groupAddUrl 创建用户组url
	groupAddUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/add"
	// groupDeleteUrl 删除用户组url
	groupDeleteUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/delete"
	// getGroupListUrl 组列表查询url
	getGroupListUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/group/getlist"
	// personVerifyUrl 人脸实名认证url
	personVerifyUrl = "https://aip.baidubce.com/rest/2.0/face/v3/person/verify"
	// sessioncodeUrl 随机校验码接口url
	sessioncodeUrl = "https://aip.baidubce.com/rest/2.0/face/v1/faceliveness/sessioncode"
	// faceverifyUrl 在线图片活体检测url
	faceverifyUrl = "https://aip.baidubce.com/rest/2.0/face/v3/faceverify"
	// videoverifyUrl 视频活体检测url
	videoverifyUrl = "https://aip.baidubce.com/rest/2.0/face/v1/faceliveness/verify"
)

// image
const (
	//组合接口api
	combinationurl = "https://aip.baidubce.com/api/v1/solution/direct/imagerecognition/combination"

	//通用物体和场景识别高级版
	advancedgeneralurl = "https://aip.baidubce.com/rest/2.0/image-classify/v2/advanced_general"

	//图像单主体检测
	objectdetecturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/object_detect"

	//动物识别
	animalurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/animal"

	//植物识别
	planturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/plant"

	//logo识别
	logourl = "https://aip.baidubce.com/rest/2.0/image-classify/v2/logo"

	//果蔬识别
	ingredienturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/classify/ingredient"

	//自定义菜品识别-入库
	dishaddurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/dish/add"

	//自定义菜品-检索
	dishsearchurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/dish/search"

	//自定义菜品-删除
	dishdeleteurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/dish/delete"

	//菜品识别
	dishurl = "https://aip.baidubce.com/rest/2.0/image-classify/v2/dish"

	//红酒识别
	redwineurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/redwine"

	//货币识别
	currencyurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/currency"

	//地标识别
	landmarkurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/landmark"

	//图像多主体检测
	multiobjectdetecturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/multi_object_detect"

	//自定义红酒-入库
	redwineaddurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/redwine/add"

	//自定义红酒-检索
	redwinesearchurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/redwine/search"

	//自定义红酒-删除
	redwinedeleteurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/redwine/delete"

	//自定义红酒—更新
	redwineupdateurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/redwine/update"

	//黑白图像上色
	colourizeurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/colourize"

	//图像风格转换
	styletransurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/style_trans"

	//人像动漫化
	selfieanimeurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/selfie_anime"

	//天空分割
	skysegurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/sky_seg"

	//图像去雾
	dehazeurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/dehaze"

	//图像对比度增强
	contrastenhanceurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/contrast_enhance"

	//图像无损放大
	imagequalityenhanceurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/image_quality_enhance"

	//拉伸图像恢复
	stretchrestoreurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/stretch_restore"

	//图像修复
	inpaintingurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/inpainting"

	//图像清晰度增强
	imagedefinitionenhanceurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/image_definition_enhance"

	//图像色彩增强
	colorenhanceurl = "https://aip.baidubce.com/rest/2.0/image-process/v1/color_enhance"

	//相似图片搜索—入库
	similaraddurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/add"

	//相似图片搜索—检索
	similarsearchurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/search"

	//相似图片搜索—删除
	similardeleteurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/delete"

	//相似图片搜索—更新
	similarupdateurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/similar/update"

	//相同图片搜索—入库
	samehqaddurl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/add"

	//相同图片搜索—检索
	samehqsearchurl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/search"

	//相同图片搜索—删除
	samehqdeleteurl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/delete"

	//相同图片搜索—更新
	samehqupdateurl = "https://aip.baidubce.com/rest/2.0/realtime_search/same_hq/update"

	//商品图片搜索—入库
	productaddurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/add"

	//商品图片搜索—检索
	productsearch = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/search"

	//商品图片搜索—删除
	productdeleteurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/delete"

	//商品图片搜索—更新
	productupdateurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/realtime_search/product/update"

	//绘本图片搜索—入库
	picturebookaddurl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/add"

	//绘本图片搜索—检索
	picturebooksearchurl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/search"

	//绘本图片搜索—删除
	picturebookdeleteurl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/delete"

	//绘本图片搜索—更新
	picturebookupdateurl = "https://aip.baidubce.com/rest/2.0/imagesearch/v1/realtime_search/picturebook/update"

	//车型识别
	carurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/car"

	//车辆检测
	vehicledetecturl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_detect"

	//车辆外观损伤识别
	vehicledamageurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_damage"

	//车流统计
	trafficflowurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/traffic_flow"

	//车辆属性识别
	vehicleattrurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_attr"

	//车辆分割
	vehiclesegurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_seg"

	//车辆检测-高空版
	vehicledetecthighurl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/vehicle_detect_high"
)

// ocr
const (
	//通用文字识别，参数为本地图片
	basicGeneralurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/general_basic"

	//通用文字识别高精度版
	basicAccurateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"

	//通用文字识别（含位置信息版）
	generalurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/general"

	//调用通用文字识别（含位置高精度版）
	accurateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate"

	//网络图片文字识别
	webImageurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage"

	//身份证识别
	idcardurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/idcard"

	//身份证混贴识别
	multiidcardurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/multi_idcard"

	//银行卡识别
	bankcardurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/bankcard"

	//驾驶证识别
	drivingLicenseurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/driving_license"

	//行驶证识别
	vehicleLicenseurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_license"

	//车牌识别
	licensePlateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/license_plate"

	//营业执照识别
	businessLicenseurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/business_license"

	//名片识别
	businesscardurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/business_card"

	//通用票据识别
	receipturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/receipt"

	//表格文字识别同步接口
	formurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/form"

	//表格文字识别
	tableRecognitionAsyncurl = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/request"

	//表格识别结果
	getTableRecognitionResulturl = "https://aip.baidubce.com/rest/2.0/solution/v1/form_ocr/get_request_result"

	//试卷分析与识别
	docAnalysisurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/doc_analysis"

	//仪表盘读数识别
	meterurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/meter"

	//网络图片文字识别（含位置版）
	webimageLocurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/webimage_loc"

	//增值税发票识别
	vatInvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice"

	//增值税发票验真
	vatinvoiceverificationurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vat_invoice_verification"

	//定额发票识别
	quotainvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/quota_invoice"

	//出租车票识别
	taxiReceipturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/taxi_receipt"

	//VIN码识别
	vinCodeurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vin_code"

	//火车票识别
	trainTicketurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/train_ticket"

	//数字识别
	numberurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/numbers"

	//飞机行程单识别
	airTicketurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/air_ticket"

	//二维码识别
	qrcodeurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/qrcode"

	//手写文字识别
	handwritingurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/handwriting"

	//护照识别
	passporturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/passport"

	//通用机打发票识别
	invoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/invoice"

	//户口本识别
	householdRegisterurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/household_register"

	//机动车销售发票识别
	vehicleInvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_invoice"

	//车辆合格证识别
	vehicleCertificateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_certificate"

	//港澳通行证识别
	HKMacauexitentrypermiturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/HK_Macau_exitentrypermit"

	//台湾通行证识别
	taiwanexitentrypermiturl = "https://aip.baidubce.com/rest/2.0/ocr/v1/taiwan_exitentrypermit"

	//出生医学证明识别
	birthcertificateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/birth_certificate"

	//多卡证类别检测
	multicardclassifyurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/multi_card_classify"

	//汽车票识别
	busticketurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/bus_ticket"

	//过来过桥费发票识别
	tollinvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/toll_invoice"

	//船票识别
	ferryticketurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/ferry_ticket"

	//网约车行程单识别
	onlinetaxiitineraryurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/online_taxi_itinerary"

	//医疗发票识别
	medicalinvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/medical_invoice"

	//医疗费用结算单识别
	medicalstatementurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/medical_statement"

	//医疗费用明细识别
	medicaldetailurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/medical_detail"

	//病案首页识别
	medicalrecordurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/medical_record"

	//保单识别
	insurancedocumentsurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/insurance_documents"

	//机动车登记证书识别
	vehicleregistrationcertificateurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/vehicle_registration_certificate"

	//二手车销售发票识别
	usedvehicleinvoiceurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/used_vehicle_invoice"

	//公式识别
	formulaurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/formula"

	//印章识别
	sealurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/seal"

	//彩票识别
	lotteryurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/lottery"

	//门脸文字识别
	facadeurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/facade"

	//只能结构化识别
	intelligentocrurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/intelligent_ocr"

	//自定义模板通用版
	recogniseurl = "https://aip.baidubce.com/rest/2.0/solution/v1/iocr/recognise"

	//自定义模板财会版
	financeurl = "https://aip.baidubce.com/rest/2.0/solution/v1/iocr/recognise/finance"

	//办公文档识别
	docanalysisofficeurl = "https://aip.baidubce.com/rest/2.0/ocr/v1/doc_analysis_office"
)

// contentcensor
const (
	// imgCensorUrl 图像审核url
	imgCensorUrl = "https://aip.baidubce.com/rest/2.0/solution/v1/img_censor/v2/user_defined"

	// textCensorUrl 文本审核url
	textCensorUrl = "https://aip.baidubce.com/rest/2.0/solution/v1/text_censor/v2/user_defined"

	// videoCensorUrl 短视频审核url
	videoCensorUrl = "https://aip.baidubce.com/rest/2.0/solution/v1/video_censor/v2/user_defined"

	// voiceCensorUrl 语音审核url
	voiceCensorUrl = "https://aip.baidubce.com/rest/2.0/solution/v1/voice_censor/v2/user_defined"
)

// nlp
const (
	// lexerUrl 词法分析通用版url
	lexerUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/lexer"
	// lexerCustomUrl 词法分析定制版url
	lexerCustomUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/lexer_custom"
	// wordEmbVecUrl 词向量表示url
	wordEmbVecUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v2/word_emb_vec"
	// wordEmbSimUrl 词义相似度url
	wordEmbSimUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v2/word_emb_sim"
	// dnnlmCnUrl DNN语言模型url
	dnnlmCnUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v2/dnnlm_cn"
	// depparserUrl 依存句法分析url
	depparserUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v2/depparser"
	// simnetUrl 短文本相似度url
	simnetUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v2/simnet"
	// keywordUrl 文章标签url
	keywordUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword"
	// topicUrl 文章分类url
	topicUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/topic"
	// ecnetUrl 文本纠错url
	ecnetUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/ecnet"
	// newsSummaryUrl 新闻摘要url
	newsSummaryUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/news_summary"
	// commentTagUrl 评论观点抽取通用版url
	commentTagUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v2/comment_tag"
	// commentTagCustomUrl 评论观点抽取定制版url
	commentTagCustomUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v2/comment_tag_custom"
	// sentimentClassifyUrl 情感倾向分析通用版url
	sentimentClassifyUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/sentiment_classify"
	// sentimentClassifyCustomUrl 情感倾向分析定制版url
	sentimentClassifyCustomUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/sentiment_classify_custom"
	// emotionUrl 对话情绪识别url
	emotionUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/emotion"
	// entityLevelSentimentUrl 多实体情感倾向分析url
	entityLevelSentimentUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment"
	// entityLevelSentimentAdd 多实体情感倾向分析实体库新增url
	entityLevelSentimentAddUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/add"
	// entityLevelSentimentList 多实体情感倾向分析实体库查询url
	entityLevelSentimentListUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/list"
	// entityLevelSentimentDeleteRepo 多实体情感倾向分析实体库删除url
	entityLevelSentimentDeleteRepoUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/delete_repo"
	// entityLevelSentimentQuery 多实体情感倾向分析实体名单查询url
	entityLevelSentimentQueryUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/query"
	// entityLevelSentimentDelete 多实体情感倾向分析实体名单删除url
	entityLevelSentimentDeleteUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/entity_level_sentiment/delete"
	// addressUrl 地址识别url
	addressUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/address"
	// ecommentUrl 消费者评论分析业务url
	ecommentUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/ecomment"
	// ecommentMiningUrl 消费者评论分析-评论挖掘url
	ecommentMiningUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/ecomment/mining"
	// ecommentQueryMiningUrl 消费者评论分析-查询评论挖掘url
	ecommentQueryMiningUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/ecomment/query_mining"
	// ecommentTrainUrl 费者评论分析-评论分析训练url
	ecommentTrainUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/ecomment/train"
	// ecommentQueryTrainUrl 消费者评论分析-查询评论分析训练url
	ecommentQueryTrainUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/ecomment/query_train"
	// ecommentDeployUrl 消费者评论分析-模型服务部署url
	ecommentDeployUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/ecomment/deploy"
	// ecommentQueryDeployUrl 消费者评论分析-查询模型服务状态url
	ecommentQueryDeployUrl = "https://aip.baidubce.com/rpc/2.0/nlp/v1/ecomment/query_deploy"
)

// bodyAnalysis
const (
	// bodyAnalysisUrl 人体关键点识别url
	bodyAnalysisUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_analysis"

	// bodyAttrUrl 人体检测和属性识别url
	bodyAttrUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_attr"

	// bodyNumUrl 人流量统计url
	bodyNumUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_num"

	// gestureUrl 手势识别url
	gestureUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/gesture"

	// bodySegUrl 人像分割url
	bodySegUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_seg"

	// driverBehaviorUrl 驾驶行为分析url
	driverBehaviorUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/driver_behavior"

	// bodyTrackingUrl 人流量统计（动态版）url
	bodyTrackingUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/body_tracking"

	// handAnalysisUrl 手部关键点识别url
	handAnalysisUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/hand_analysis"

	// bodyDangerUrl 危险行为识别url
	bodyDangerUrl = "https://aip.baidubce.com/rest/2.0/video-classify/v1/body_danger"

	// fingertipUrl 指尖检测url
	fingertipUrl = "https://aip.baidubce.com/rest/2.0/image-classify/v1/fingertip"
)

// speech
const (
	// 短语音识别标准版url
	asrUrl = "http://vop.baidu.com/server_api"

	// 语音合成url
	synthesisUrl = "https://tsn.baidu.com/text2audio"
)
