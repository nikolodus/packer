package main

import (
	"testing"
	"time"
)

func TestMinimalConfig(t *testing.T) {
	_, warns, errs := NewConfig(minimalConfig())

	testConfigOk(t, warns, errs)
}

func TestTimeout(t *testing.T) {
	raw := minimalConfig()
	raw["shutdown_timeout"] = "3m"
	conf, warns, err := NewConfig(raw)
	testConfigOk(t, warns, err)
	if conf.ShutdownConfig.Timeout != 3 * time.Minute {
		t.Fatalf("shutdown_timeout sould be equal 3 minutes, got %v", conf.ShutdownConfig.Timeout)
	}
}

func TestRAMReservation(t *testing.T) {
	raw := minimalConfig()
	raw["RAM_reservation"] = 1000
	raw["RAM_reserve_all"] = true
	_, warns, err := NewConfig(raw)
	testConfigErr(t, warns, err)
}

func minimalConfig() map[string]interface{} {
	return map[string]interface{}{
		"vcenter_server": "vcenter.domain.local",
		"username":     "root",
		"password":     "vmware",
		"template":     "ubuntu",
		"vm_name":      "vm1",
		"host":         "esxi1.domain.local",
		"ssh_username": "root",
		"ssh_password": "secret",
	}
}

func testConfigOk(t *testing.T, warns []string, err error) {
	if len(warns) > 0 {
		t.Fatalf("Should be no warnings: %#v", warns)
	}
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}
}

func testConfigErr(t *testing.T, warns []string, err error) {
	if len(warns) > 0 {
		t.Fatalf("Should be no warnings: %#v", warns)
	}
	if err == nil {
		t.Fatal("An error is not raised")
	}
}
