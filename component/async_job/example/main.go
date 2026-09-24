package main

import (
	"context"
	"log"
	"time"

	"food_delivery/component/async_job"
)

func main() {
	job_1 := async_job.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second)
		log.Println("job 1 started")
		return nil
		//return errors.New("some error at job 1")
	})

	//job_1.SetRetryDurations([]time.Duration{time.Second * 2})
	//
	//if err := job_1.Execute(context.Background()); err != nil {
	//	log.Println(job_1.State(), err)
	//
	//	for {
	//		if err := job_1.Retry(context.Background()); err != nil {
	//			log.Println(err)
	//		}
	//
	//		if job_1.State() != async_job.StateRetryFailed || job_1.State() != async_job.StateCompleted {
	//			break
	//		}
	//	}
	//}

	job_2 := async_job.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 2)
		log.Println("job 2 started")
		//return errors.New("some error at job 2")
		return nil
	})

	job_3 := async_job.NewJob(func(ctx context.Context) error {
		time.Sleep(time.Second * 3)
		log.Println("job 3 started")
		return nil
	})

	group := async_job.NewGroup(true, job_1, job_2, job_3)

	if err := group.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
