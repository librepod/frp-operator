package controllers

import (
	"strings"
	"testing"
)

func TestResolveFRPCImage_PerCRWins(t *testing.T) {
	got, err := resolveFRPCImage("private.example/frpc:custom", "private.example/frpc:default")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "private.example/frpc:custom" {
		t.Errorf("expected per-CR image to win, got %q", got)
	}
}

func TestResolveFRPCImage_DefaultUsedWhenPerCREmpty(t *testing.T) {
	got, err := resolveFRPCImage("", "private.example/frpc:default")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if got != "private.example/frpc:default" {
		t.Errorf("expected default image, got %q", got)
	}
}

func TestResolveFRPCImage_BothEmptyReturnsError(t *testing.T) {
	_, err := resolveFRPCImage("", "")
	if err == nil {
		t.Fatal("expected error when neither image is set, got nil")
	}
	msg := err.Error()
	if !strings.Contains(msg, "spec.podTemplate.image") {
		t.Errorf("error should mention spec.podTemplate.image, got %q", msg)
	}
	if !strings.Contains(msg, "--frpc-default-image") {
		t.Errorf("error should mention --frpc-default-image, got %q", msg)
	}
}
