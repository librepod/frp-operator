package controllers

import "fmt"

// resolveFRPCImage returns the frpc image to run for a Client.
//
// Precedence:
//  1. perCR  — the Client's spec.podTemplate.image (highest priority)
//  2. defaultImage — the operator's --frpc-default-image flag
//
// It returns an error when neither is set, so the controller can fail the
// Client loudly rather than silently deploying nothing.
//
// Note: use fmt.Errorf, not errors.New — this package imports
// k8s.io/apimachinery/pkg/api/errors as "errors", which has no New().
func resolveFRPCImage(perCR, defaultImage string) (string, error) {
	if perCR != "" {
		return perCR, nil
	}
	if defaultImage != "" {
		return defaultImage, nil
	}
	return "", fmt.Errorf("no frpc image configured: set spec.podTemplate.image or the operator --frpc-default-image flag")
}
