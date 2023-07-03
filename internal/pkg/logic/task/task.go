package task

import (
	"context"
	"fmt"
	"github.com/pro911/golang-demo/internal/pkg/biz/consts"
	"github.com/pro911/golang-demo/internal/pkg/dal"
	"github.com/pro911/golang-demo/internal/pkg/dal/mao"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func TaskSendData(ctx context.Context) ([]*mao.TaskOption, error) {
	unixTime := time.Now().Unix()
	//查询定时任务区
	taskModel := dal.GetQuery().Task
	tasks, err := taskModel.
		WithContext(context.Background()).
		Where(
			taskModel.Status.Eq(1),
			taskModel.TriggerEnable.Eq(1),
			taskModel.NextAt.Lte(unixTime),
		).
		Order(taskModel.NextAt.Desc(), taskModel.CreatedAt).
		Limit(10000).
		Find()
	if err != nil {
		return nil, err
	}

	// 统计接口id
	taskIDs := make([]string, 0, len(tasks))
	for _, tInfo := range tasks {
		if tInfo.TaskID != "" {
			taskIDs = append(taskIDs, tInfo.TaskID)
		}
	}

	// 获取可选参数
	collectionTaskOption := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectTaskOption)
	cursor, err := collectionTaskOption.Find(ctx, bson.D{{"task_id", bson.D{{"$in", taskIDs}}}})
	if err != nil {
		return nil, err
	}
	var taskOptions []*mao.TaskOption
	if err = cursor.All(ctx, &taskOptions); err != nil {
		return nil, err
	}

	for _, v := range taskOptions {
		fmt.Printf("%#v\n", v)
	}

	return taskOptions, err
	//for k, v := range taskOptions {
	//	fmt.Println(k, v)
	//}

	//查询时间
	//collectionTaskTestEvents := dal.GetMongo().Database(dal.MongoDB()).Collection(consts.CollectTaskTestEvents)
	//taskTestEvents, err := collectionTaskTestEvents.Find(ctx, bson.D{{"task_id", bson.D{{"$in", taskIDs}}}})
	//if err != nil {
	//	return nil, err
	//}

	//return nil, err
}
