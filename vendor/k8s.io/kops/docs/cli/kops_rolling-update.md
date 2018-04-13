
<!--- This file is automatically generated by make gen-cli-docs; changes should be made in the go CLI command code (under cmd/kops) -->

## kops rolling-update

Rolling update a cluster.

### Synopsis


This command updates a kubernetes cluster to match the cloud and kops specifications.

To perform a rolling update, you need to update the cloud resources first with the command
`kops update cluster`.

If rolling-update does not report that the cluster needs to be rolled, you can force the cluster to be
rolled with the force flag.  Rolling update drains and validates the cluster by default.  A cluster is
deemed validated when all required nodes are running and all pods in the kube-system namespace are operational.
When a node is deleted, rolling-update sleeps the interval for the node type, and then tries for the same period
of time for the cluster to be validated.  For instance, setting --master-interval=3m causes rolling-update
to wait for 3 minutes after a master is rolled, and another 3 minutes for the cluster to stabilize and pass
validation.

Note: terraform users will need to run all of the following commands from the same directory
`kops update cluster --target=terraform` then `terraform plan` then
`terraform apply` prior to running `kops rolling-update cluster`.

### Examples

```
  # Preview a rolling-update.
  kops rolling-update cluster
  
  # Roll the currently selected kops cluster with defaults.
  # Nodes will be drained and the cluster will be validated between node replacement.
  kops rolling-update cluster --yes
  
  # Roll the k8s-cluster.example.com kops cluster,
  # do not fail if the cluster does not validate,
  # wait 8 min to create new node, and wait at least
  # 8 min to validate the cluster.
  kops rolling-update cluster k8s-cluster.example.com --yes \
  --fail-on-validate-error="false" \
  --master-interval=8m \
  --node-interval=8m
  
  # Roll the k8s-cluster.example.com kops cluster,
  # do not validate the cluster because of the cloudonly flag.
  # Force the entire cluster to roll, even if rolling update
  # reports that the cluster does not need to be rolled.
  kops rolling-update cluster k8s-cluster.example.com --yes \
  --cloudonly \
  --force
  
  # Roll the k8s-cluster.example.com kops cluster,
  # only roll the node instancegroup,
  # use the new drain an validate functionality.
  kops rolling-update cluster k8s-cluster.example.com --yes \
  --fail-on-validate-error="false" \
  --node-interval 8m \
  --instance-group nodes
```

### Options inherited from parent commands

```
      --alsologtostderr                  log to standard error as well as files
      --config string                    config file (default is $HOME/.kops.yaml)
      --log_backtrace_at traceLocation   when logging hits line file:N, emit a stack trace (default :0)
      --log_dir string                   If non-empty, write log files in this directory
      --logtostderr                      log to standard error instead of files (default false)
      --name string                      Name of cluster. Overrides KOPS_CLUSTER_NAME environment variable
      --state string                     Location of state storage. Overrides KOPS_STATE_STORE environment variable
      --stderrthreshold severity         logs at or above this threshold go to stderr (default 2)
  -v, --v Level                          log level for V logs
      --vmodule moduleSpec               comma-separated list of pattern=N settings for file-filtered logging
```

### SEE ALSO
* [kops](kops.md)	 - kops is Kubernetes ops.
* [kops rolling-update cluster](kops_rolling-update_cluster.md)	 - Rolling update a cluster.
