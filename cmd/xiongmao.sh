#!/bin/bash

# 提示用户输入文件名
echo "请输入Go文件的名称（不包括扩展名，例如：main）："
read fileName

# 如果用户未输入文件名，默认使用test
if [ -z "$fileName" ]; then
    fileName="test"
fi

# 将文件名的首字母转为大写
capitalizedFileName="${fileName^}"

handler="../internal/handler"
model="../internal/model"
router="../internal/router"
service="../internal/service"

if [ ! -d "$handler" ]; then
    mkdir -p "$handler"
fi

if [ ! -d "$model" ]; then
    mkdir -p "$model"
fi

if [ ! -d "$router" ]; then
    mkdir -p "$router"
fi

if [ ! -d "$service" ]; then
    mkdir -p "$service"
fi

# handler 开始
handlerContent="package $(basename $handler)

import (
	\"github.com/gofiber/fiber/v2\"
	\"xiongmao/internal/model\"
	\"xiongmao/internal/service\"
	\"xiongmao/pkg/respond\"
)

// ${capitalizedFileName}Add 添加
func ${capitalizedFileName}Add(c *fiber.Ctx) error {

	var req model.${capitalizedFileName}AddParam

	if err := c.BodyParser(&req); err != nil {
		panic(respond.Error(\"参数错误\"))
	}

    service.${capitalizedFileName}Add(req)

	return respond.Ok(c)
}"
# handler 结束

# model 开始
modelContent="package $(basename $model)

import (
    \"xiongmao/pkg/custom\"
)

// $capitalizedFileName 模型
type $capitalizedFileName struct {
    custom.Id
    Name string \`gorm:\"type:varchar(20);unique;comment:名称\" json:\"name,omitempty\"\`
    custom.At
}

// ${capitalizedFileName}AddParam 添加请求参数
type ${capitalizedFileName}AddParam struct {
    Name string \`json:\"name\"\`
}"
# model 结束

# router 开始
routerContent="package $(basename $router)

import (
	\"github.com/gofiber/fiber/v2\"
    \"xiongmao/internal/handler\"
)

// ${fileName} 路由，记得把 ${fileName}(app) 添加把 router.go 中
func ${fileName}(app *fiber.App) {

	${fileName}Router := app.Group(\"${fileName}\")
	{
		${fileName}Router.Post(\"add\", handler.${capitalizedFileName}Add)
	}

}"
# router 结束

# service 结束
serviceContent="package $(basename $service)

import (
	\"gorm.io/gorm\"
    \"xiongmao/internal/model\"
    \"xiongmao/pkg/db\"
    \"xiongmao/pkg/respond\"
    \"xiongmao/pkg/utils\"
)

type ${fileName} struct {
	model.$capitalizedFileName
}

// ${capitalizedFileName}Add 添加
func ${capitalizedFileName}Add(req model.${capitalizedFileName}AddParam) {

    dbBegin := db.Gorm.Begin()

	name := req.Name

	var ${fileName} ${fileName}

	${fileName}.Name = name

	r := dbBegin.Create(&${fileName})

	if r.Error != nil {
		dbBegin.Rollback()
		panic(respond.Error(\"保存失败\"))
	}

	dbBegin.Commit()
}

// BeforeCreate 请不到删除 db *gorm.DB，要不然重写不了id值，用于添加逻辑
func (param *${fileName}) BeforeCreate(db *gorm.DB) (err error) {

	param.ID = utils.UUID()

	return
}"
# model 结束

# 在目录中创建一个main.go文件并写入内容
echo "$handlerContent" > "$handler/$fileName.go"
echo "$modelContent"   > "$model/$fileName.go"
echo "$routerContent"  > "$router/$fileName.go"
echo "$serviceContent" > "$service/$fileName.go"