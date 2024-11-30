package waitstatus

import (
	"context"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	testdesc "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc/testing"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/clock"
	testclock "k8s.io/utils/clock/testing"
)

func disk(id, status string, instances ...string) *ecs.Disk {
	disk := &ecs.Disk{
		DiskId: id,
		Status: status,
	}
	as := make([]ecs.Attachment, len(instances))
	for i, instance := range instances {
		as[i] = ecs.Attachment{
			InstanceId: instance,
		}
	}
	disk.Attachments.Attachment = as
	return disk
}

func isNextStatus(previousStatus string, instanceID string) StatusPredicate[*ecs.Disk] {
	return func(disk *ecs.Disk) bool {
		found := IsInstanceAttached(disk, instanceID)
		switch previousStatus {
		case "Attaching":
			if found {
				return true
			}
		case "Detaching":
			if !found {
				return true
			}
		}
		return disk.Status != previousStatus
	}
}

func testEachImpl(t *testing.T, clk clock.WithTicker, f func(*testing.T, *testdesc.FakeClient, StatusWaiter[ecs.Disk])) {
	t.Run("batched", func(t *testing.T) {
		client := &testdesc.FakeClient{}

		ctx := context.Background()
		ctx, cancel := context.WithCancel(ctx)
		t.Cleanup(cancel)
		batched := NewBatched(client, clk)
		go batched.Run(ctx)

		f(t, client, batched)
	})
	t.Run("simple", func(t *testing.T) {
		client := &testdesc.FakeClient{}
		f(t, client, NewSimple(client, clk))
	})
}

func TestWaitForDisks(t *testing.T) {
	clk := testclock.NewFakeClock(time.Now())
	testEachImpl(t, clk, func(t *testing.T, client *testdesc.FakeClient, waiter StatusWaiter[ecs.Disk]) {
		client.Disks.Store("d1", disk("d1", "Attaching"))

		finalDisk := disk("d1", "In_use", "i1")
		go func() {
			time.Sleep(100 * time.Millisecond) // for request to be picked up
			t.Logf("Step for first poll")
			clk.Step(100 * time.Millisecond)
			time.Sleep(100 * time.Millisecond) // for the first poll to finish

			client.Disks.Store("d1", finalDisk)
			t.Logf("Step for second poll")
			clk.Step(pollInterval)
		}()

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		disk, err := waiter.WaitFor(ctx, "d1", isNextStatus("Attaching", "i1"))
		assert.NoError(t, err)
		assert.Equal(t, finalDisk, disk)
	})
}

func TestWaitForDisksCancel(t *testing.T) {
	testEachImpl(t, clock.RealClock{}, func(t *testing.T, client *testdesc.FakeClient, waiter StatusWaiter[ecs.Disk]) {
		client.Disks.Store("d1", disk("d1", "Attaching"))
		ctx, cancel := context.WithCancel(context.Background())
		cancel()
		disk, err := waiter.WaitFor(ctx, "d1", isNextStatus("Attaching", "i1"))
		assert.Equal(t, ctx.Err(), err)
		assert.Nil(t, disk)
	})
}

func TestWaitForDisksUnfound(t *testing.T) {
	testEachImpl(t, clock.RealClock{}, func(t *testing.T, client *testdesc.FakeClient, waiter StatusWaiter[ecs.Disk]) {
		disk, err := waiter.WaitFor(context.Background(), "d1", isNextStatus("Attaching", "i1"))
		assert.NoError(t, err)
		assert.Nil(t, disk)
	})
}

func TestIsInstanceAttached(t *testing.T) {
	tests := []struct {
		name       string
		disk       *ecs.Disk
		instanceID string
		attached   bool
	}{
		{
			name:       "attaching",
			disk:       disk("d1", "Attaching"),
			instanceID: "i1",
			attached:   false,
		}, {
			name:       "detaching",
			disk:       disk("d1", "Detaching", "i1"),
			instanceID: "i1",
			attached:   true,
		}, {
			name:       "attaching - multi",
			disk:       disk("d1", "Attaching", "i2"),
			instanceID: "i1",
			attached:   false,
		}, {
			name:       "detaching - multi",
			disk:       disk("d1", "Detaching", "i1", "i2"),
			instanceID: "i1",
			attached:   true,
		}, {
			name:       "attaching - done",
			disk:       disk("d1", "In_use", "i1"),
			instanceID: "i1",
			attached:   true,
		}, {
			name:       "detaching - done",
			disk:       disk("d1", "Available"),
			instanceID: "i1",
			attached:   false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.attached, IsInstanceAttached(test.disk, test.instanceID))
		})
	}
}
