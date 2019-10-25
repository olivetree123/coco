package coco

import "testing"

func TestCoco(t *testing.T) {
	coco := NewCoco()
	go coco.Run("127.0.0.1", 8000)
}