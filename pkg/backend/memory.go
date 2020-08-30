package backend

import (
	"context"

	"github.com/jeffrom/job-manager/pkg/job"
)

// Memory is an in-memory backend intended for testing.
type Memory struct {
	configs map[string]*job.Queue
	jobs    map[string]*job.Job
}

func NewMemory() *Memory {
	return &Memory{
		configs: make(map[string]*job.Queue),
		jobs:    make(map[string]*job.Job),
	}
}

func (m *Memory) GetQueue(ctx context.Context, name string) (*job.Queue, error) {
	if cfg, ok := m.configs[name]; ok {
		return cfg, nil
	}
	return nil, ErrNotFound
}

func (m *Memory) SaveQueue(ctx context.Context, queue *job.Queue) error {
	m.configs[queue.Name] = queue
	return nil
}

func (m *Memory) ListQueues(ctx context.Context, opts *job.ListOpts) (*job.Queues, error) {
	jobs := &job.Queues{}
	for _, job := range m.configs {
		jobs.Queues = append(jobs.Queues, job)
	}
	return jobs, nil
}

func (m *Memory) EnqueueJobs(ctx context.Context, jobArgs *job.Jobs) error {
	for _, jobArg := range jobArgs.Jobs {
		m.jobs[jobArg.Id] = jobArg
	}
	return nil
}

func (m *Memory) DequeueJobs(ctx context.Context, num int, opts *job.ListOpts) (*job.Jobs, error) {
	if opts == nil {
		opts = &job.ListOpts{}
	}
	if len(opts.Statuses) == 0 {
		opts.Statuses = []job.Status{job.Status_QUEUED}
	}
	jobs, err := m.ListJobs(ctx, opts)
	if err != nil {
		return nil, err
	}
	for _, jobData := range jobs.Jobs {
		jobData.Status = job.Status_RUNNING
	}
	return jobs, nil
}

func (m *Memory) AckJobs(ctx context.Context, results *job.Results) error {
	for _, res := range results.Results {
		job, ok := m.jobs[res.Id]
		if !ok {
			return ErrNotFound
		}
		job.Status = res.Status
	}
	return nil
}

func (m *Memory) GetJobByID(ctx context.Context, id string) (*job.Job, error) {
	return m.jobs[id], nil
}

func (m *Memory) ListJobs(ctx context.Context, opts *job.ListOpts) (*job.Jobs, error) {
	if opts == nil {
		opts = &job.ListOpts{}
	}
	res := &job.Jobs{}
	for _, jobData := range m.jobs {
		if statuses := opts.Statuses; len(statuses) > 0 && !job.HasStatus(jobData, statuses) {
			continue
		}
		res.Jobs = append(res.Jobs, jobData)
	}
	return res, nil
}
