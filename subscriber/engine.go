package subscriber

import (
	"context"
	"log"

	"food_delivery/common"

	"food_delivery/component/async_job"

	"food_delivery/component/appctx"

	"food_delivery/pubsub"
)

type consumerJob struct {
	Title string
	Hdl   func(ctx context.Context, message *pubsub.Message) error
}

type consumerEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appCtx appctx.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appCtx}
}

func (engine *consumerEngine) Start() error {
	engine.startSubTopic(common.TopicUserLikeRestaurant, true, IncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx))
	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isConcurrent bool, consumerJobs ...consumerJob) error {
	c, _ := engine.appCtx.GetPubSub().Subscribe(context.Background(), topic)

	for _, item := range consumerJobs {
		log.Println("start sub topic", item.Title)
	}

	getJobHandler := func(job *consumerJob, message *pubsub.Message) async_job.JobHandler {
		return func(ctx context.Context) error {
			log.Println("start sub topic", job.Title)
			return job.Hdl(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]async_job.Job, len(consumerJobs))

			for i := range consumerJobs {
				jobHdl := getJobHandler(&consumerJobs[i], msg)
				jobHdlArr[i] = async_job.NewJob(jobHdl)
			}

			group := async_job.NewGroup(isConcurrent, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
