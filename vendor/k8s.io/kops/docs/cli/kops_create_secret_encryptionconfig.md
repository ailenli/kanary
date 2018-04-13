
<!--- This file is automatically generated by make gen-cli-docs; changes should be made in the go CLI command code (under cmd/kops) -->

## kops create secret encryptionconfig

Create an encryption config.

### Synopsis


Create a new encryption config, and store it in the state store. Used to configure encryption-at-rest by the kube-apiserver process on each of the master nodes. The config is not updated by this command.

```
kops create secret encryptionconfig
```

### Examples

```
  # Create a new encryption config.
  kops create secret encryptionconfig -f config.yaml \
  --name k8s-cluster.example.com --state s3://example.com
  # Replace an existing encryption config secret.
  kops create secret encryptionconfig -f config.yaml --force \
  --name k8s-cluster.example.com --state s3://example.com
```

### Options

```
  -f, -- string   Path to encryption config yaml file
      --force     Force replace the kops secret if it already exists
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
* [kops create secret](kops_create_secret.md)	 - Create a secret.
