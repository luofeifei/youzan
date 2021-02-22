package validator

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translationsEn "github.com/go-playground/validator/v10/translations/en"
	translationsZh "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"sync"
)

// Validator 验证器
type Validator struct {
	once     sync.Once
	validate *validator.Validate
}

var (
	_                   binding.StructValidator = &Validator{}
	enUs                                        = en.New()
	zhCn                                        = zh.New()
	universalTranslator                         = ut.New(enUs, zhCn)
	zhTrans, _                                  = universalTranslator.GetTranslator(zhCn.Locale())
	enTrans, _                                  = universalTranslator.GetTranslator(enUs.Locale())
)

// ValidateStruct 验证结构体
func (v *Validator) ValidateStruct(obj interface{}) error {
	if kindOfData(obj) == reflect.Struct {
		v.lazyInit()
		if err := v.validate.Struct(obj); err != nil {
			return err
		}
	}
	return nil
}

// Engine 获取验证器
func (v *Validator) Engine() interface{} {
	v.lazyInit()
	return v.validate
}

// lazyInit 延迟初始化
func (v *Validator) lazyInit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")
		// 获取form tag
		v.validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := field.Tag.Get("form")
			if name != "" {
				return name
			}
			return field.Name
		})
		_ = translationsZh.RegisterDefaultTranslations(v.validate, zhTrans)
		_ = translationsEn.RegisterDefaultTranslations(v.validate, enTrans)
	})
}

// ValidErrors 验证之后的错误信息
type ValidErrors struct {
	ErrorsInfo map[string]string
	triggered  bool
}

func (validErrors *ValidErrors) add(key, value string) {
	validErrors.ErrorsInfo[key] = value
	validErrors.triggered = true
}

// IsValid 是否验证成功
func (validErrors *ValidErrors) IsValid() bool {
	return !validErrors.triggered
}

func newValidErrors() *ValidErrors {
	return &ValidErrors{
		triggered:  false,
		ErrorsInfo: make(map[string]string),
	}
}

// 获取验证码对象
// 	errs := validator.Get().Var(form.User, "required,email,gt=1,lt=10")
//	fmt.Println(validator.GetValidErrors(ctx, errs))
func Get() *validator.Validate {
	v, _ := binding.Validator.Engine().(*validator.Validate)
	return v
}

func GetVal(c *gin.Context, field interface{}, tag string) *ValidErrors {
	err := Get().Var(field, tag)
	if err != nil {
		return GetValidErrors(c, err)
	}
	return &ValidErrors{}
}

// 转换错误
func GetValidErrors(c *gin.Context, err error) *ValidErrors {
	lang := c.DefaultQuery("lang", "zh-cn")
	var trans ut.Translator
	if lang == "zh-cn" || lang == "zh" {
		trans, _ = universalTranslator.GetTranslator(zhCn.Locale())
	} else {
		trans, _ = universalTranslator.GetTranslator(enUs.Locale())
	}
	var validErrors = newValidErrors()
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if ok {
			for _, value := range errs {
				validErrors.add(value.Field(), value.Translate(trans))
			}
		} else {
			validErrors.add("err", "数据解析失败")
		}
	}
	return validErrors
}

// Bind 自定义错误信息, 如果没有匹配需要在 configs/validator-messages.yaml 中添加对应处理数据
func Bind(c *gin.Context, param interface{}) *ValidErrors {
	err := c.ShouldBindJSON(param)
	if err != nil {
		return GetValidErrors(c, err)
	}
	return &ValidErrors{}
}

func BindQuery(c *gin.Context, param interface{}) *ValidErrors {
	err := c.ShouldBindQuery(param)
	if err != nil {
		return GetValidErrors(c, err)
	}
	return &ValidErrors{}
}

func kindOfData(data interface{}) reflect.Kind {
	value := reflect.ValueOf(data)
	valueType := value.Kind()
	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}
