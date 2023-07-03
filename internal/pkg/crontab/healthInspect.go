package crontab

import (
	"context"
	"fmt"
	"github.com/pro911/golang-demo/internal/pkg/biz/consts"
	"github.com/pro911/golang-demo/internal/pkg/dal"
	"github.com/pro911/golang-demo/internal/pkg/dal/mao"
	"go.mongodb.org/mongo-driver/bson"
)

func HealthInspect() {
	fmt.Println("开始执行定时任务")
	tb := dal.GetQuery().Project
	tbI, err := tb.WithContext(context.Background()).Where(tb.ProjectID.Eq(1)).First()
	if err != nil {
		fmt.Println("错了")
	} else {
		fmt.Println(tbI)
	}

	ctx := context.Background()
	//查询mongo
	var taskTestEventsData = mao.TaskTestEvents{}
	c := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectTaskTestEvents)
	err = c.FindOne(ctx, bson.D{{"task_id", "834635b8-971c-4e93-b01d-8e3892aea300"}}).Decode(&taskTestEventsData)
	fmt.Println(err, taskTestEventsData)
	fmt.Printf("结束执行定时任务\n\n\n")
}
