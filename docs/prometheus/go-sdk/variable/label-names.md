# Label Names Variable Builder

## Constructor

```golang
import "github.com/perses/plugins/prometheus/sdk/go/label-names"

var options []labelnames.Option
labelnames.PrometheusLabelNames(options...)
```

Need a list of options.

## Default options

- None

## Available options

### Matchers

```golang
import "github.com/perses/plugins/prometheus/sdk/go/label-names"

var matchers []string
labelnames.Matchers(matchers...)
```

Define matchers filtering the result.

### AddMatcher

```golang
import "github.com/perses/plugins/prometheus/sdk/go/label-names"

labelnames.AddMatcher("my_super_matcher")
```

Define a matcher filtering the result.

### Datasource

```golang
import "github.com/perses/plugins/prometheus/sdk/go/label-names"

labelnames.Datasource("datasourceName")
```

Define the datasource where the expression will be executed.

### Filter

```golang
import "github.com/perses/perses/go-sdk/variable"

variable.Filter(variables...)
```

Mainly used by Mainly used by [variable group](https://perses.dev/perses/docs/dac/go/variable-group).. It will filter the current variable with the
provided variables.
The filter will be applied only if matchers don't have curly brackets.

## Example

```golang
package main

import (
	"github.com/perses/perses/go-sdk/dashboard"
	listvariable "github.com/perses/perses/go-sdk/variable/list-variable"
	labelnames "github.com/perses/plugins/prometheus/sdk/go/variable/label-names"
)

func main() {
	dashboard.New("Example Dashboard",
		dashboard.AddVariable("namespaceLabels", listvariable.List(
			labelnames.PrometheusLabelNames(
				labelnames.Matchers("kube_namespace_labels{stack=\"$stack\",prometheus=\"$prometheus\",prometheus_namespace=\"$prometheus_namespace\",namespace=\"$namespace\"}"),
				labelnames.Datasource("prometheusDemo"),
			),
		)),
	)
}
```
