---
title: "Polyaxon Experiments Query Syntax"
sub_link: "polyaxon-query-syntax/experiments"
meta_title: "Polyaxon Experiments Query Syntax Specification - Polyaxon References"
meta_description: "Polyaxon Experiments Query Syntax Specification."
visibility: public
status: published
tags:
    - api
    - reference
    - polyaxon
    - query
    - syntax
    - search
    - experiment
    - tracking
    - insights
sidebar: "polyaxon-query-syntax"
---

# Searching Experiments

## Query

field                         | condition
------------------------------|------------------
`id`                          | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`name`                        | [string condition](/references/polyaxon-query-syntax/#query-with-string-condition)
`description`                 | [string condition](/references/polyaxon-query-syntax/#query-with-string-condition)
`created_at`                  | [datetime condition](/references/polyaxon-query-syntax/#query-with-datetime-condition)
`updated_at`                  | [datetime condition](/references/polyaxon-query-syntax/#query-with-datetime-condition)
`started_at`                  | [datetime condition](/references/polyaxon-query-syntax/#query-with-datetime-condition)
`finished_at`                 | [datetime condition](/references/polyaxon-query-syntax/#query-with-datetime-condition)
`user.id`                     | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`user.username`               | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`build.id`                    | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`project.*` (e.g. project.id) | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`group.*` (e.g. group.id)     | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`commit`                      | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`name`                        | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`status`                      | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`tags`                        | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`params.*`              | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`metrics.*`                    | [comparison condition](/references/polyaxon-query-syntax/#query-with-comparison-condition)
`backend`                     | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)
`framework`                   | [value condition](/references/polyaxon-query-syntax/#query-with-value-condition)


## Sort

field:

 * `created_at`
 * `updated_at`
 * `started_at`
 * `finished_at`
 * `metrics.*`
