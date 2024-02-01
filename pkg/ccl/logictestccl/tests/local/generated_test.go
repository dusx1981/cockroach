// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

// Code generated by generate-logictest, DO NOT EDIT.

package testlocal

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/cockroach/pkg/base"
	"github.com/cockroachdb/cockroach/pkg/build/bazel"
	"github.com/cockroachdb/cockroach/pkg/ccl"
	"github.com/cockroachdb/cockroach/pkg/security/securityassets"
	"github.com/cockroachdb/cockroach/pkg/security/securitytest"
	"github.com/cockroachdb/cockroach/pkg/server"
	"github.com/cockroachdb/cockroach/pkg/sql/logictest"
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/testutils/skip"
	"github.com/cockroachdb/cockroach/pkg/testutils/testcluster"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/randutil"
)

const configIdx = 0

var cclLogicTestDir string

func init() {
	if bazel.BuiltWithBazel() {
		var err error
		cclLogicTestDir, err = bazel.Runfile("pkg/ccl/logictestccl/testdata/logic_test")
		if err != nil {
			panic(err)
		}
	} else {
		cclLogicTestDir = "../../../../ccl/logictestccl/testdata/logic_test"
	}
}

func TestMain(m *testing.M) {
	defer ccl.TestingEnableEnterprise()()
	securityassets.SetLoader(securitytest.EmbeddedAssets)
	randutil.SeedForTests()
	serverutils.InitTestServerFactory(server.TestServerFactory)
	serverutils.InitTestClusterFactory(testcluster.TestClusterFactory)

	defer serverutils.TestingSetDefaultTenantSelectionOverride(
		base.TestIsForStuffThatShouldWorkWithSecondaryTenantsButDoesntYet(76378),
	)()

	os.Exit(m.Run())
}

func runCCLLogicTest(t *testing.T, file string) {
	skip.UnderDeadlock(t, "times out and/or hangs")
	logictest.RunLogicTest(t, logictest.TestServerArgs{}, configIdx, filepath.Join(cclLogicTestDir, file))
}

// TestLogic_tmp runs any tests that are prefixed with "_", in which a dedicated
// test is not generated for. This allows developers to create and run temporary
// test files that are not checked into the repository, without repeatedly
// regenerating and reverting changes to this file, generated_test.go.
//
// TODO(mgartner): Add file filtering so that individual files can be run,
// instead of all files with the "_" prefix.
func TestLogic_tmp(t *testing.T) {
	defer leaktest.AfterTest(t)()
	var glob string
	glob = filepath.Join(cclLogicTestDir, "_*")
	logictest.RunLogicTests(t, logictest.TestServerArgs{}, configIdx, glob)
}

func TestCCLLogic_as_of(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "as_of")
}

func TestCCLLogic_auto_rehoming(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "auto_rehoming")
}

func TestCCLLogic_case_sensitive_names(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "case_sensitive_names")
}

func TestCCLLogic_changefeed(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "changefeed")
}

func TestCCLLogic_crdb_internal(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "crdb_internal")
}

func TestCCLLogic_explain_call_plpgsql(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "explain_call_plpgsql")
}

func TestCCLLogic_explain_redact(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "explain_redact")
}

func TestCCLLogic_fips_ready(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "fips_ready")
}

func TestCCLLogic_new_schema_changer(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "new_schema_changer")
}

func TestCCLLogic_partitioning(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "partitioning")
}

func TestCCLLogic_partitioning_all_by_nothing(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "partitioning_all_by_nothing")
}

func TestCCLLogic_partitioning_constrained_scans(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "partitioning_constrained_scans")
}

func TestCCLLogic_partitioning_enum(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "partitioning_enum")
}

func TestCCLLogic_partitioning_implicit(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "partitioning_implicit")
}

func TestCCLLogic_partitioning_index(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "partitioning_index")
}

func TestCCLLogic_pgcrypto_builtins(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "pgcrypto_builtins")
}

func TestCCLLogic_plpgsql_block(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "plpgsql_block")
}

func TestCCLLogic_plpgsql_cursor(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "plpgsql_cursor")
}

func TestCCLLogic_plpgsql_record(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "plpgsql_record")
}

func TestCCLLogic_plpgsql_unsupported(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "plpgsql_unsupported")
}

func TestCCLLogic_procedure_plpgsql(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "procedure_plpgsql")
}

func TestCCLLogic_read_committed(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "read_committed")
}

func TestCCLLogic_redact_descriptor(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "redact_descriptor")
}

func TestCCLLogic_refcursor(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "refcursor")
}

func TestCCLLogic_restore(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "restore")
}

func TestCCLLogic_schema_change_in_txn(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "schema_change_in_txn")
}

func TestCCLLogic_show_create(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "show_create")
}

func TestCCLLogic_tenant_capability(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "tenant_capability")
}

func TestCCLLogic_tenant_usage(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "tenant_usage")
}

func TestCCLLogic_udf_params(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "udf_params")
}

func TestCCLLogic_udf_plpgsql(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "udf_plpgsql")
}

func TestCCLLogic_udf_rewrite(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "udf_rewrite")
}

func TestCCLLogic_udf_volatility_check(
	t *testing.T,
) {
	defer leaktest.AfterTest(t)()
	runCCLLogicTest(t, "udf_volatility_check")
}
