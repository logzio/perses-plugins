# PieChart Panel Model

```yaml
kind: "PieChart"
spec:
  legend:        <Legend specification> # Optional
  calculation:   <Calculation specification>
  format:        <Format specification> # Optional
  sort:          <enum = "asc" | "desc"> # Optional
  mode:          <enum = "value" | "percentage"> # Optional
  radius:        <number>
```

## Legend specification

```yaml
position: <enum = "bottom" | "right">
mode: <enum = "list" | "table"> # Optional
size: <enum = "small" | "medium"> # Optional
values:
  - <Calculation specification> # Optional
```

## Calculation specification

See [common plugin definitions](https://perses.dev/perses/docs/plugins/common/#calculation-specification).

## Format specification

See [common plugin definitions](https://perses.dev/perses/docs/plugins/common/#format-specification).
