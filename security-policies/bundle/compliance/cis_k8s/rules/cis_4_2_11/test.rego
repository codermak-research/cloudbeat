package compliance.cis_k8s.rules.cis_4_2_11

import data.compliance.cis_k8s.data_adapter
import data.kubernetes_common.test_data
import data.lib.test

test_violation {
	eval_fail with input as rule_input("--rotate-certificates=false")
	eval_fail with input as rule_input_with_external("", create_process_config(false))
	eval_fail with input as rule_input_with_external("--rotate-certificates=false", create_process_config(true))
}

test_pass {
	eval_pass with input as rule_input("")
	eval_pass with input as rule_input("--rotate-certificates=true")
	eval_pass with input as rule_input_with_external("--rotate-certificates=true", create_process_config(false))
	eval_pass with input as rule_input_with_external("--rotate-certificates=true", create_process_config(true))
	eval_pass with input as rule_input_with_external("", create_process_config(true))
}

test_not_evaluated {
	not_eval with input as test_data.process_input("some_process", [])
}

rule_input(argument) = test_data.process_input("kubelet", [argument])

rule_input_with_external(argument, external_data) = test_data.process_input_with_external_data("kubelet", [argument], external_data)

create_process_config(rotateCertificates) = {"config": {"rotateCertificates": rotateCertificates}}

eval_fail {
	test.assert_fail(finding) with data.benchmark_data_adapter as data_adapter
}

eval_pass {
	test.assert_pass(finding) with data.benchmark_data_adapter as data_adapter
}

not_eval {
	not finding with data.benchmark_data_adapter as data_adapter
}