
<!--- This file is automatically generated by make gen-cli-docs; changes should be made in the go CLI command code (under cmd/kops) -->

## kops validate

Validate a kops cluster.

### Synopsis


This commands validates the following components: 

  1. All k8s masters are running and have "Ready" status.  
  2. All k8s nodes are running and have "Ready" status.  
  3. Componentstatues returns healthly for all components.  
  4. All pods in the kube-system namespace are running and healthy.

### Examples

```
  # Validate a cluster.
  # This command uses the currently selected kops cluster as
  # set by the kubectl config.
  kops validate cluster
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
* [kops validate cluster](kops_validate_cluster.md)	 - Validate a kops cluster.
