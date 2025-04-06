package main

type ZpoolDevice struct {
	Name  string `json:"name"`
	State string `json:"state"`
	Read  string `json:"read"`
	Write string `json:"write"`
	Cksum string `json:"cksum"`
}

type ZpoolStatus struct {
	PoolName  string        `json:"pool_name"`
	Status    string        `json:"status"`
	Scan      string        `json:"scan,omitempty"`
	Errors    string        `json:"errors"`
	Devices   []ZpoolDevice `json:"devices"`
	RawOutput string        `json:"raw_output"`
} 