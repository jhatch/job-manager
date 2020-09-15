package handler

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"net/http"

	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"

	apiv1 "github.com/jeffrom/job-manager/pkg/api/v1"
	"github.com/jeffrom/job-manager/pkg/backend"
	"github.com/jeffrom/job-manager/pkg/internal"
	"github.com/jeffrom/job-manager/pkg/label"
	"github.com/jeffrom/job-manager/pkg/resource"
	jobv1 "github.com/jeffrom/job-manager/pkg/resource/job/v1"
	"github.com/jeffrom/job-manager/pkg/schema"
)

type EnqueueJobs struct {
	// schemaCache
}

func (h *EnqueueJobs) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Func(func(w http.ResponseWriter, r *http.Request) error {
		ctx := r.Context()
		be := backend.FromMiddleware(ctx)
		var params apiv1.EnqueueJobsRequest
		if err := UnmarshalBody(r, &params, true); err != nil {
			return err
		}

		resources := &resource.Jobs{Jobs: make([]*resource.Job, len(params.Jobs))}
		jobs := &jobv1.Jobs{Jobs: make([]*jobv1.Job, len(params.Jobs))}
		now := timestamppb.New(internal.GetTimeProvider(ctx).Now())
		for i, jobArg := range params.Jobs {
			queue, err := be.GetQueue(ctx, jobArg.Job)
			if err != nil {
				return handleBackendErrors(err, "queue", jobArg.Job)
			}

			// validate args if there is a schema
			scm, err := schema.Parse(queue.SchemaRaw)
			if err != nil {
				return err
			}

			if err := scm.ValidateArgs(ctx, jobArg.Args); err != nil {
				return handleSchemaErrors(err, "job", "", "invalid job arguments")
			}

			if queue.Unique {
				unique, err := checkArgUniqueness(ctx, be, scm, jobArg.Args)
				if err != nil {
					return err
				}
				if unique {
					// return conflict error
					return resource.NewUnprocessableEntityError("queue", queue.ID, "A job with matching arguments is executing")
				}
			}

			var claims label.Claims
			if jobArg.Data != nil {
				claims, err = label.ParseClaims(jobArg.Data.Claims)
				if err != nil {
					return err
				}
			}

			jb := &jobv1.Job{
				Name:       jobArg.Job,
				Args:       jobArg.Args,
				Data:       jobArg.Data,
				Status:     jobv1.StatusQueued,
				EnqueuedAt: now,
			}
			jobs.Jobs[i] = jb

			jobRes := jobv1.NewJobFromProto(jb, claims)
			resources.Jobs[i] = jobRes
			// fmt.Printf("JOB: %+v\n", jobs.Jobs[i])
		}

		res, err := be.EnqueueJobs(ctx, resources)
		if err != nil {
			return err
		}

		return MarshalResponse(w, r, &apiv1.EnqueueJobsResponse{
			Jobs: res.IDs(),
		})
	})(w, r)
}

func checkArgUniqueness(ctx context.Context, be backend.Interface, scm *schema.Schema, args []*structpb.Value) (bool, error) {
	iargs := make([]interface{}, len(args))
	for i, arg := range args {
		iargs[i] = arg.AsInterface()
	}
	key, err := uniquenessKeyFromArgs(iargs)
	if err != nil {
		return false, err
	}
	found, err := be.GetSetJobKeys(ctx, []string{key})
	if err != nil {
		return false, err
	}
	if found {
		return true, nil
	}
	return false, nil
}

func uniquenessKeyFromArgs(args []interface{}) (string, error) {
	b, err := json.Marshal(args)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(b)
	return string(sum[:]), nil
}
