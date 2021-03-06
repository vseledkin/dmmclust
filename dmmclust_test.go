package dmmclust

import (
	"math/rand"
	"testing"
)

func TestCluster_Words(t *testing.T) {
	c := Cluster{
		dist: distro{
			0: 1,
			1: 2,
		},
	}
	if len(c.Words()) != 2 {
		t.Errorf("Expected only 2 words in the cluster")
	}
}

func TestShitConfig(t *testing.T) {
	confs := []Config{
		{},
		{Sampler: NewGibbs(rand.New(rand.NewSource(1337)))},
		{Score: Algorithm3},
	}

	for _, conf := range confs {
		if _, err := FindClusters(nil, conf); err == nil {
			t.Errorf("Expected Config to fail")
		}
	}
}
